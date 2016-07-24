package controllers

import (
	"fmt"
	"net/http"

	"github.com/SofyanHadiA/goscrape/services/sessions"
	"github.com/SofyanHadiA/linqcore"
)

type pageCtrl struct {
	view linqcore.View
}

func NewPageCtrl(view linqcore.View) pageCtrl {
	return pageCtrl{view: view}
}

func (ctrl pageCtrl) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.view.ParseHTML("login.html", w, r, nil)
}

func (ctrl pageCtrl) UserPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, userRole := sessions.GetUser(r)

	if userRole == "user" {
		data := make(map[string]interface{})
		data["UserName"] = userName
		data["UserRole"] = userRole

		fmt.Println(data)

		ctrl.view.ParseHTML("user.html", w, r, data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func (ctrl pageCtrl) AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, userRole := sessions.GetUser(r)
	if userRole == "admin" {

		data := make(map[string]interface{})
		data["UserName"] = userName
		data["UserRole"] = userRole

		fmt.Println(data)

		ctrl.view.ParseHTML("admin.html", w, r, data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func (ctrl pageCtrl) SuperAdminPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, userRole := sessions.GetUser(r)
	if userRole == "super admin" {

		data := make(map[string]interface{})
		data["UserName"] = userName
		data["UserRole"] = userRole

		fmt.Println(data)

		ctrl.view.ParseHTML("superAdmin.html", w, r, data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func (ctrl pageCtrl) ScraperPageHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.view.ParseHTML("scraper.html", w, r, nil)
}
