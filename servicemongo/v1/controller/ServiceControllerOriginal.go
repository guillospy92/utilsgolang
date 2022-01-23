package controller

import (
	"encoding/json"
	"fmt"
	"guihub.com/guillospy92/servicemongo/v1/repositories"
	"net/http"
)

func AddServicePurchased(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Body)
	var serviceCreated repositories.ServicePurchasedOriginal

	err := json.Unmarshal([]byte(jsonCreated), &serviceCreated)

	fmt.Println(serviceCreated.SubscriptionID)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	service := repositories.ServicePurchasedDocumentRepository{}

	err = service.SaveServicePurchased(serviceCreated)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte("service created ok"))

}

func ServicePurchasedCharge(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Body)
	var serviceCreated repositories.ServicePurchasedChargeOriginal

	err := json.Unmarshal([]byte(jsonCharge), &serviceCreated)

	if err != nil {
		fmt.Println("paso por aqiui 22222")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	service := repositories.ServicePurchasedDocumentRepository{}

	err = service.SaveServiceCharge(serviceCreated)

	if err != nil {
		fmt.Println("paso por aqiui")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte("service created ok"))
}

func ServicePurchasedFind(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Body)
	var serviceCreated repositories.ServicePurchasedChargeOriginal

	err := json.Unmarshal([]byte(jsonCharge), &serviceCreated)

	if err != nil {
		fmt.Println("paso por aqiui 22222")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	service := repositories.ServicePurchasedDocumentRepository{}

	err = service.FindServicePurchased(serviceCreated.SubscriptionID)

	if err != nil {
		fmt.Println("paso por aqiui")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte("service created ok"))
}

func ServicePurchasedCancel(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Body)
	var serviceCreated repositories.ServicePurchasedCancelOriginal

	err := json.Unmarshal([]byte(jsonCancel), &serviceCreated)

	if err != nil {
		fmt.Println("paso por aqiui 22222")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	service := repositories.ServicePurchasedDocumentRepository{}

	err = service.SaveServiceCancel(serviceCreated)

	if err != nil {
		fmt.Println("paso por aqiui")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte("service created ok"))
}
