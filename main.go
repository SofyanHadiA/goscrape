package main

import (
	"net/http"
	"strconv"

	"github.com/SofyanHadiA/linqcore"
	"github.com/SofyanHadiA/linqcore/database"
	"github.com/SofyanHadiA/linqcore/utils"
	"github.com/gorilla/mux"
)

// server main method

var router = mux.NewRouter()

func main() {
	envarConfigPrefix := "LINQ_"
	configs := linqcore.NewConfig(envarConfigPrefix)

	utils.SetLogLevel(configs.GetIntConfig("app.logLevel"))
	server := configs.GetStrConfig("app.server") + ":" + strconv.Itoa(configs.GetIntConfig("app.port"))

	var db = database.MySqlDB(
		configs.GetStrConfig("db.host"),
		configs.GetStrConfig("db.username"),
		configs.GetStrConfig("db.password"),
		configs.GetStrConfig("db.database"),
		configs.GetIntConfig("db.port"),
	)

	view := linqcore.NewView("./views/", configs)

	router := linqcore.NewRouter(GetRoutes(view, db))

	http.Handle("/", router)
	utils.Log.Info("Listen and serve " + server)
	err := http.ListenAndServe(server, nil)
	utils.HandleFatal(err)
}
