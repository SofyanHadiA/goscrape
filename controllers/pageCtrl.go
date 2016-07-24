package controllers

import (
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
	ctrl.view.ParseHTML("login.html", w, r)
}

func (ctrl pageCtrl) UserPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := sessions.GetUserName(r)
	if userName != "" {
		ctrl.view.ParseHTML("user.html", w, r)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func (ctrl pageCtrl) AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := sessions.GetUserName(r)
	if userName != "" {
		ctrl.view.ParseHTML("adminPage.html", w, r)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func (ctrl pageCtrl) SuperAdminPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := sessions.GetUserName(r)
	if userName != "" {
		ctrl.view.ParseHTML("superAdminPage.html", w, r)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func (ctrl pageCtrl) ScrapperPageHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.view.ParseHTML("scraper.html", w, r)
}
