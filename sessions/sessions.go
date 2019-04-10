package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

//Temporary secret key for cookie authentication
var Store = sessions.NewCookieStore([]byte("SECRETPASS"))

//IsLoggedIn will check if the user has an active session and return True
func IsLoggedIn(r *http.Request) bool {
	session, _ := Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		return true
	}
	return false
}
