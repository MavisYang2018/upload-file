package auth

import (
	"net/http"
	"errors"
	"bytes"
	"io"
	"../session"
)

var (
	AuthError error
)

func init() {
	AuthError = errors.New("Auth Error")
}

type Body struct {
	bytes.Buffer
}

func (Body) Close() error {
	return nil
}

//
//
//
type ReqInfo struct {
	Path string
	Req  *http.Request
	ReqNewError error
}

//
// auth request
//
func AuthRun (w http.ResponseWriter,req *http.Request) (*ReqInfo,error) {
	reqInfo := ReqInfo{Path:req.URL.Path}

	//copy body
	buf := new(bytes.Buffer)
	tmpr := io.TeeReader(req.Body,buf)
	body := new(Body)
	body.ReadFrom(tmpr)

	//auth
	if !ADAuth(w,req) {
		//http.Error(w,AuthError.Error(),415)
		reqInfo.Req,_ = http.NewRequest(req.Method,req.URL.String(),body)
		reqInfo.Req.Header = req.Header
		return &reqInfo,AuthError
	}
	defer req.Body.Close()

	//handle
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqInfo.Req,reqInfo.ReqNewError = http.NewRequest(req.Method,req.URL.String(),body)
	if reqInfo.ReqNewError != nil {
		return nil,reqInfo.ReqNewError
	}
	reqInfo.Req.Header = req.Header

	return &reqInfo,nil
}

func ADAuth (w http.ResponseWriter,req *http.Request) bool {
	sess, _ := session.GlobalSessions.SessionStart(w, req)
	defer sess.SessionRelease(w)

	if sess.Get("ice") != nil {
		return true
	} else {
		return false
	}
}
