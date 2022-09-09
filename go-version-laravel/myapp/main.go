package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/guillospy92/utilsgolang/go-version-laravel/celeritas"
	"github.com/guillospy92/utilsgolang/go-version-laravel/myapp/handlers"
	"log"
	"net/http"
	"os"
)

// application struct
type application struct {
	App      *celeritas.Accelerator
	Handlers *handlers.Handlers
}

func main() {
	a := initApplication()
	a.App.StartServer()
}

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cel := celeritas.NewAccelerator(path)
	app := application{}
	app.App = cel

	app.Handlers = &handlers.Handlers{
		App: cel,
	}

	cel.Routes = app.LoadRoutes()

	if err != nil {
		log.Fatal(err)
	}
	return &app
}

// LoadRoutes we should improve it
func (a *application) LoadRoutes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/jet-page", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.HomeWithGo)
	a.App.Routes.Get("/sessions", a.Handlers.SessionPage)

	a.App.Routes.Get("/test-database", func(writer http.ResponseWriter, request *http.Request) {
		query := `select id, name from users where id = 1::text`
		row := a.App.DB.Pool.QueryRowContext(request.Context(), query)

		var id string
		var name string

		err := row.Scan(&id, &name)

		if err != nil {
			a.App.ErrorLog.Println("err query", err)
		}

		_, _ = fmt.Fprintf(writer, "%s, %s", id, name)
	})

	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
