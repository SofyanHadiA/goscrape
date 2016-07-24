package api

import (
	"net/http"

	"github.com/SofyanHadiA/goscrape/services/sessions"
	"github.com/SofyanHadiA/linqcore/database"
)

type loginAPI struct {
	db database.IDB
}

func NewLoginAPI(db database.IDB) loginAPI {
	return loginAPI{db: db}
}

// LoginHandler handle login operations
func (ctrl loginAPI) LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		sessions.SetSession(name, response)
		redirectTarget = "/internal"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// LogoutHandler handle logout handler
func (ctrl loginAPI) LogoutHandler(response http.ResponseWriter, request *http.Request) {
	sessions.ClearSession(response)
	http.Redirect(response, request, "/", 302)
}
