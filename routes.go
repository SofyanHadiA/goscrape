package main

import (
	"github.com/SofyanHadiA/linqcore"
	"github.com/SofyanHadiA/linqcore/database"

	"github.com/SofyanHadiA/goscrape/api"
	"github.com/SofyanHadiA/goscrape/controllers"
)

// GetRoutes resolve and return mapped Routes
func GetRoutes(view linqcore.View, db database.IDB) linqcore.Routes {
	pageCtrl := controllers.NewPageCtrl(view)

	loginAPI := api.NewLoginAPI(db)
	scraperAPI := api.NewScraperAPI(db)

	return linqcore.Routes{
		linqcore.Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: pageCtrl.LoginPageHandler},
		linqcore.Route{Name: "Login", Method: "GET", Pattern: "/login", HandlerFunc: pageCtrl.LoginPageHandler},
		linqcore.Route{Name: "UserPage", Method: "GET", Pattern: "/user", HandlerFunc: pageCtrl.UserPageHandler},
		linqcore.Route{Name: "AdminPage", Method: "GET", Pattern: "/admin", HandlerFunc: pageCtrl.AdminPageHandler},
		linqcore.Route{Name: "SuperAdminPage", Method: "GET", Pattern: "/superadmin", HandlerFunc: pageCtrl.SuperAdminPageHandler},
		linqcore.Route{Name: "ScrapePage", Method: "GET", Pattern: "/scrape", HandlerFunc: pageCtrl.ScraperPageHandler},

		// API
		linqcore.Route{Name: "LoginAction", Method: "POST", Pattern: "/login", HandlerFunc: loginAPI.LoginHandler},
		linqcore.Route{Name: "LogoutAction", Method: "POST", Pattern: "/logout", HandlerFunc: loginAPI.LogoutHandler},
		linqcore.Route{Name: "ScrapeAction", Method: "POST", Pattern: "/scrape", HandlerFunc: scraperAPI.ScraperHandler},
	}
}
