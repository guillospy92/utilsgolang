package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(w, "%q: %q", key, str)
		} else {
			fmt.Fprintf(w, "%q: %v", key, value)
		}
	}

	fmt.Fprintf(w, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/vars", metricsHandler)
	err := http.ListenAndServe("localhost:6060", mux)

	if err != nil {
		log.Println("error serve listen", err)
	}
}
