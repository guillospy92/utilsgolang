package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"guihub.com/guillospy92/servicemongo/v1/mongo"
	"guihub.com/guillospy92/servicemongo/v1/routers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routers.RoutesMap(r)

	err := initMongo()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = http.ListenAndServe(":9090", r)
	if err != nil {
		fmt.Println(err)
	}
}

func initMongo() error {
	db := mongo.ParamConnectionDocumentDB{
		Cluster:        "localhost",
		UserName:       "root",
		Password:       "root",
		DBName:         "prime",
		Port:           mongo.DefaultPort,
		ReadPreference: mongo.DefaultReadPreference,
		ConnectTimeOut: mongo.DefaultTimeOut,
	}

	client, err := mongo.NewCreateConnectionDocumentDB(&db)

	if err != nil {
		return err
	}

	mongo.ConnectApiMongoGlobal = client

	return nil
}
