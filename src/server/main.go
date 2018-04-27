package main

import (
	"net/http"
	"log"
	"os"
	"./auth"
	"runtime"
	"./config"
	"./api"
	"./fileuntil"
	"fmt"
	"github.com/gin-gonic/gin/json"
)

func init() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

//
// upload file to web server
//
func upload (w http.ResponseWriter,req *http.Request) {
	defer req.Body.Close()
	//auth
	if !auth.ADAuth() {
		w.Write([]byte("Auth Error"))
		return
	}

	// upload file
	api.Upload(w,req)
}

//
// check file exist
//
// client post json format example :
//{
//	"fn" : "server.key"
//}
func isExist (w http.ResponseWriter,req *http.Request) {
	defer req.Body.Close()
	dec := json.NewDecoder(req.Body)
	p := new(fileuntil.Path)
	err := dec.Decode(p)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}

	if isIn := fileuntil.CheckFile(fmt.Sprintf("%s/%s",config.Cfg.UploadFolder,p.Fn));!isIn {
		w.Write([]byte("false"))
	} else {
		w.Write([]byte("true"))
	}


}

func main() {
	http.HandleFunc("/upload",upload)
	http.HandleFunc("/isExist",isExist)
	if err := http.ListenAndServe(config.Cfg.ListenPort,nil);err != nil {
		log.Fatal(err)
	}
}
