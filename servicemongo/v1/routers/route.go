package routers

import (
	"github.com/gorilla/mux"
	"guihub.com/guillospy92/servicemongo/v1/controller"
	"net/http"
)

func RoutesMap(route *mux.Router) {
	route.HandleFunc("/service/add", controller.AddService).Methods(http.MethodGet)
	route.HandleFunc("/service/purchased/add", controller.NewServicePurchased).Methods(http.MethodGet)
	route.HandleFunc("/service/purchased/search", controller.FindServicePurchased).Methods(http.MethodGet)

	// routes Original
	route.HandleFunc("/service/purchased/original", controller.AddServicePurchased).Methods(http.MethodGet)
	route.HandleFunc("/service/purchased/charge", controller.ServicePurchasedCharge).Methods(http.MethodGet)
	route.HandleFunc("/service/purchased/find", controller.ServicePurchasedFind).Methods(http.MethodGet)
	route.HandleFunc("/service/purchased/cancel", controller.ServicePurchasedCancel).Methods(http.MethodGet)
}
