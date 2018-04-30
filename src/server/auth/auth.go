package auth

import (
	"../session"
	"net/http"
)

func ADAuth () bool {

	return true
}

func Login (w http.ResponseWriter, r *http.Request) {
	sess, _ := session.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	//username := sess.Get("username")
	if r.Method == "POST" {
		sess.Set("username", r.Form["username"])
	}
}
