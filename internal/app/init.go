package app

import "net/http"

type App struct {
	srv  *http.Server
	port int
}

func (app *App) Start() {

}

func (app *App) Shutdown() {

}

func Init() {
	//app := &App{8080}
	//app.Start()
	//app.Shutdown()
}
