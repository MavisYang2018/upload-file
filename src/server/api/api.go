package api

import (
	"net/http"
	"../auth"
	"fmt"
	"../config"
	"../web"
	"log"
)

func ApiRun (w http.ResponseWriter,req *http.Request) {
	//max files size 30MB
	if req.ContentLength > config.Cfg.UploadMaxSize {
		http.Error(w,"Over 30MB",415)
		return
	}

	// auth request and return request
	r,err := auth.AuthRun(w,req)
	if err != nil {
		Login(w,r.Req)
		return
	}
	defer r.Req.Body.Close()

	switch r.Path {
	case "/" :
		web.ServeIndexPage(w,r.Req)
	case "/api/upload" :
		err = Upload(w,r.Req)
	case "/api/isExist":
		err = FileStat(w,r.Req)
	}

	if err != nil {
		log.Println(err)
		fmt.Fprintln(w,err)
		return
	}
}

