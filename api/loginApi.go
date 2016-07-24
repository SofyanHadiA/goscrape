package api

import (
	"fmt"
	"net/http"

	"github.com/SofyanHadiA/goscrape/repositories"
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
		repo := repositories.NewUserRepo(ctrl.db)

		result, err := repo.Login(name, pass)

		if err != nil {
			fmt.Println(err)
			http.Redirect(response, request, redirectTarget, 302)
		}
		sessions.SetSession(*result, response)

		fmt.Println((*result).Name)
		fmt.Println((*result).UserRole)

		if result.RoleID == 1 {
			redirectTarget = "/user"
		} else if result.RoleID == 2 {
			redirectTarget = "/admin"
		} else if result.RoleID == 3 {
			redirectTarget = "/superadmin"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// LogoutHandler handle logout handler
func (ctrl loginAPI) LogoutHandler(response http.ResponseWriter, request *http.Request) {
	sessions.ClearSession(response)
	http.Redirect(response, request, "/", 302)
}
