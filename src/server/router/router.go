package router

import (
	"net/http"
	"../api"
	"github.com/gorilla/mux"
	"../config"
	"log"
	"../web"
)

func Registry () *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/",api.ApiRun)
	r.HandleFunc("/loginPage",web.ServeLoginPage)
	r.HandleFunc("/api/upload",api.ApiRun)
	r.HandleFunc("/api/isExist",api.ApiRun).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(config.Cfg.WorkDir))))
	return r
}

func Run (r *mux.Router) {
	log.Fatal(http.ListenAndServe(config.Cfg.ListenPort, r))
}
