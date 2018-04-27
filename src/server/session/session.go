package session

import (
	"github.com/astaxie/beego/session"
)

var GlobalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"iceseesion",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory",sessionConfig)
	go GlobalSessions.GC()
}
