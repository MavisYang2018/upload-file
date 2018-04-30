package api

import (
	"net/http"
	"../session"
	"../web"
)

//
// for login
//
type user struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//
// login result
//
//type result struct {
//	R bool `json:"r"`
//}


//
// login auth
//
// client example :
// POST "user_name","password"
//
func Login (w http.ResponseWriter,req *http.Request) {
	defer req.Body.Close()
	req.ParseForm()
	user := req.FormValue("user_name")
	pwd := req.FormValue("password")

	sess, _ := session.GlobalSessions.SessionStart(w, req)
	defer sess.SessionRelease(w)

	if checkAD(user,pwd) {
		sess.Set("ice",user)
		web.ServeIndexPage(w,req)
	} else {
		web.ServeLoginPage(w,req)
	}
}

//
//connect to ad
//
func checkAD (user,password string) bool {
	if user == "joe_ke" && password == "123" {
		return true
	}
	return false
}
