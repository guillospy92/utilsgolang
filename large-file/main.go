package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/guillospy92/large-file/controllers"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	initServer()
}

func initServer() {
	r := mux.NewRouter()
	RoutesMap(r)
	r.Use(Logging)

	err := http.ListenAndServe(":9090", r)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("run port 9090")
}

func Logging(handler http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, handler)
}

func RoutesMap(route *mux.Router) {
	route.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := template.ParseFiles("./upload.html")
		if err != nil {
			log.Println("error template", err)
			return
		}
		err = tmpl.Execute(writer, "data goes here")
		if err != nil {
			log.Println("error load template", err)
			return
		}

	}).Methods(http.MethodGet)

	route.HandleFunc("/upload/large", controllers.UploadFileLarge).Methods(http.MethodGet)
	route.HandleFunc("/upload/normal", controllers.UploadNormal).Methods(http.MethodPost)
}
