package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/youngjae-lim/gosnel"
)

type application struct {
	App        *gosnel.Gosnel
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
	wg         sync.WaitGroup
}

func main() {
	g := initApplication()
	go g.listenForShutdown()
	err := g.App.ListenAndServe()
	if err != nil {
		g.App.ErrorLog.Println(err)
	}
}

func (a *application) shutDown() {
	// put any clean up tasks(such as sending an email) that have to finish or should take place
	// before the application completely shuts down.

	// block until the waitgroup is empty
	a.wg.Wait()
}

func (a *application) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	a.App.InfoLog.Println("Received signal", s.String())
	a.shutDown()

	os.Exit(0)
}
