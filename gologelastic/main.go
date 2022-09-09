package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github/guillospy92/utilsgolang/gologelastic/loggers"
	"net/http"

	_ "gopkg.in/natefinch/lumberjack.v2"
)

// https://programmer.help/blogs/zap-log-base-practice.html practice full good
// https://pmihaylov.com/go-service-with-elk/ configure with container golang
// https://www.thoutam.com/2020/05/09/elasticsearch-filebeat-custom-index/ configure index filebeat

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/init_logger", func(response http.ResponseWriter, request *http.Request) {
		loggers.AddLoggerLogStash()
	}).Methods(http.MethodGet)
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		fmt.Println(err)
	}
}
