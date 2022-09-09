package controller

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"guihub.com/guillospy92/servicemongo/v1/repositories"
	"net/http"
	"time"
)

func AddService(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Body)

	serviceSave := repositories.Service{
		Name:          "service first prime",
		Id:            "co_1_co",
		Type:          "prime",
		PaymentMethod: []string{"tc", "cash"},
		CountryCode:   "co",
		Plan: []repositories.Plan{
			{
				Id:               primitive.NewObjectID(),
				Name:             "plan year",
				Price:            30_000,
				PromotionalPrice: 0,
				Benefits:         []string{"free_delivery, netflix"},
				Period:           12,
			},
			{
				Id:               primitive.NewObjectID(),
				Name:             "plan quarterly",
				Price:            30_000,
				PromotionalPrice: 0,
				Benefits:         []string{"free_delivery", "netflix"},
				Period:           12,
			},
		},
	}

	_, err := repositories.NewService(serviceSave)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte("service created ok"))
}

func NewServicePurchased(response http.ResponseWriter, request *http.Request) {

	serviceId, err := primitive.ObjectIDFromHex("61dbb57bbe58d186c793526d")

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	now := time.Now()

	servicePurchased := repositories.ServicePurchased{
		ServiceID:      serviceId,
		CountryCode:    "co",
		CityName:       "Bogota",
		TimeZone:       "America/Bogota",
		ActivationDate: primitive.NewDateTimeFromTime(now),
		IsRenovate:     true,
		Status:         "active",
		User: repositories.User{
			ID:        234,
			Name:      "Guillermo",
			Email:     "guillospy@gmail.com",
			Address:   "call 23",
			Cellphone: "3003001234",
		},
		Plan: repositories.Plan{
			Id:               primitive.NewObjectID(),
			Name:             "plan quarterly",
			Price:            30_000,
			PromotionalPrice: 0,
			Benefits:         []string{"free_delivery", "netflix"},
			Period:           12,
		},
	}

	_, err = repositories.NewServicePurchased(servicePurchased)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte("service created ok"))
}

func FindServicePurchased(response http.ResponseWriter, request *http.Request) {
	fmt.Println(1111)
	service, err := repositories.SearchServicePurchased()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(service)
}
