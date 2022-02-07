package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/youngjae-lim/gosnel"
)

func initApplication() *application {
	// get a path of the working directory
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init gosnel
	gos := &gosnel.Gosnel{}
	err = gos.New(path)
	if err != nil {
		log.Fatal(err)
	}

	gos.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: gos,
	}

	myHandlers := &handlers.Handlers{
		App: gos,
	}

	app := &application{
		App:        gos,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
