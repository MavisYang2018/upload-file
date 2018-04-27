package api

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"io"
	"../config"
	"encoding/json"
)

type jsonBody struct {
	FileName string `json:"file_name"`
}

//
// upload file
//
func Upload (w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//get upload file
	// Form attribut "uploadfile"
	file, handler, err := req.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	//create or overrite file in web server
	f, err := os.OpenFile(fmt.Sprintf("%s/%s",config.Cfg.UploadFolder,handler.Filename), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	defer f.Close()

	//write file to local
	_,err = io.Copy(f, file)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	jb := new(jsonBody)
	jb.FileName = fmt.Sprintf("upload %s success!",handler.Filename)

	err = json.NewEncoder(w).Encode(jb)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
}
