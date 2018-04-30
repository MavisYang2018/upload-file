package web

import (
	"net/http"
	"../config"
)

func ServeLoginPage (w http.ResponseWriter,req *http.Request) {
	//handle
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w,req,config.Cfg.WorkDir + "/login.html")
}

func ServeIndexPage (w http.ResponseWriter,req *http.Request) {
	//handle
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w,req,config.Cfg.WorkDir + "/index.html")
}
