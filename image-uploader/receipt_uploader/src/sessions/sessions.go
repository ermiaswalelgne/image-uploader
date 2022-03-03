package sessions

import (
	gsessions "github.com/gorilla/sessions"
	"net/http"
)

// A gorilla session wrapper to track session
// store stroes the session of login users
var store = gsessions.NewCookieStore([]byte("usersession"))

func Get(req *http.Request) (*gsessions.Session, error) {
	return store.Get(req, "default-session-name")
}

func GetNamed(req *http.Request, name string) (*gsessions.Session, error) {
	return store.Get(req, name)
}
