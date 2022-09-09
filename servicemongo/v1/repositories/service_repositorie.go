package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"guihub.com/guillospy92/servicemongo/v1/mongo"
	"log"
	"time"
)

// ServiceRepositories struct
type ServiceRepositories struct{}

// User struct
type User struct {
	ID             int    `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	PhonePrefix    string `json:"phone_prefix" bson:"phone_prefix"`
	Cellphone      string `json:"cellphone" bson:"cellphone"`
	Address        string `json:"address" bson:"address"`
	PostalCode     string `json:"postal_code" bson:"postal_code"`
	IdentityNumber string `json:"identity_number" bson:"identity_number"`
	IdentityType   string `json:"identity_type" bson:"identity_type"`
}

// Plan struct
type Plan struct {
	Id               primitive.ObjectID `bson:"_id"`
	Name             string             `bson:"name"`
	Price            float64            `bson:"price"`
	PromotionalPrice float64            `bson:"promotional_price"`
	Benefits         []string           `bson:"benefits"`
	Period           int                `bson:"period"`
}

// Payment struct
type Payment struct {
	CreditCard CreditCard `json:"credit_card" bson:"credit_card"`
	History    []History  `json:"history,omitempty" bson:"history"`
}

// CreditCard attributes of the credit card.
type CreditCard struct {
	ID           int    `json:"id" bson:"id"`
	Franchise    string `json:"franchise" bson:"franchise"`
	LastFour     int    `json:"last_four" bson:"last_four"`
	Installments int    `json:"installments" bson:"installments"`
	Type         string `json:"type" bson:"type"`
	CVV          *int   `json:"cvv,omitempty"`
}

// History attributes of the transaction history.
type History struct {
	TransactionID     string    `json:"transaction_id" bson:"transaction_id"`
	ReferenceCode     string    `json:"reference_code" bson:"reference_code"`
	TransactionStatus string    `json:"transaction_status" bson:"transaction_status"`
	TransactionDate   time.Time `json:"transaction_date" bson:"transaction_date"`
}

// Service struct
type Service struct {
	Name          string   `bson:"name"`
	Id            string   `bson:"id"`
	Type          string   `bson:"type"`
	PaymentMethod []string `bson:"payment_method"`
	CountryCode   string   `bson:"country_code"`
	Plan          []Plan   `bson:"plan"`
}

// ServicePurchased struct deprecated
type ServicePurchased struct {
	// ID             string             `bson:"_id"`
	ServiceID      primitive.ObjectID `bson:"service_id"`
	CountryCode    string             `bson:"country_code"`
	CityName       string             `bson:"city_name"`
	TimeZone       string             `bson:"time_zone"`
	ActivationDate primitive.DateTime `bson:"activation_date"`
	ExpirationDate primitive.DateTime `bson:"expiration_date"`
	CancelDate     primitive.DateTime `bson:"cancel_date"`
	IsRenovate     bool               `bson:"is_renovate"`
	Status         string             `bson:"status"`
	User           User               `bson:"user"`
	Plan           Plan               `bson:"plan"`
	Payment        Payment            `bson:"payment"`
}

// ServicePurchasedLogs struct repository in the collection logs in documentDB
type ServicePurchasedLogs struct {
	UserID            int       `bson:"user_id"`
	SubscriptionID    string    `bson:"subscription_id"`
	ServiceID         string    `bson:"service_id"`
	PlanID            string    `bson:"plan_id"`
	PlanName          string    `bson:"plan_name"`
	PlanPrice         float64   `bson:"plan_price"`
	TransactionType   string    `bson:"transaction_type"`
	TransactionDate   time.Time `bson:"transaction_date"`
	TransactionStatus string    `bson:"transaction_status"`
}

// NewService add new service in mongo
// https://ichi.pro/es/golang-y-mongodb-con-polimorfismo-y-bson-unmarshall-145649298531523
func NewService(s Service) (Service, error) {
	result, err := mongo.ConnectApiMongoGlobal.InsertOne(context.Background(), "services", s)

	if err != nil {
		return Service{}, err
	}

	fmt.Println(result)

	return Service{}, nil
}

// NewServicePurchased add new service in mongo
func NewServicePurchased(s ServicePurchased) (Service, error) {
	result, err := mongo.ConnectApiMongoGlobal.InsertOne(context.Background(), "service_purchased", s)

	if err != nil {
		return Service{}, err
	}

	fmt.Println(result)

	return Service{}, nil
}

// SearchServicePurchased search services
func SearchServicePurchased() (ServicePurchased, error) {
	var service map[interface{}]interface{}

	filter := bson.D{
		{"plan.benefits", "free_delivery"},
		{"country_code", "co"},
		{"user.id", 19776},
		{"status", "active"},
	}

	err := mongo.ConnectApiMongoGlobal.FindOne(context.Background(), "service_purchased", filter).Decode(&service)

	if err != nil {
		return ServicePurchased{}, err
	}

	return ServicePurchased{}, err
}

func GetServiceOriginal(subscriptionID string) (ServicePurchasedOriginal, error) {
	filter := bson.M{"subscription_id": subscriptionID}
	var service ServicePurchasedOriginal

	err := mongo.ConnectApiMongoGlobal.FindOne(context.Background(), "services_purchased", filter).Decode(&service)

	if err != nil {
		log.Println(err)
		return ServicePurchasedOriginal{}, err
	}

	return service, err
}

// SaveLogsServicePurchased saved log all transaction in documentDB
func SaveLogsServicePurchased(service ServicePurchasedOriginal, transaction string) error {
	// preparing data logs
	logs := ServicePurchasedLogs{
		UserID:            service.User.ID,
		SubscriptionID:    service.SubscriptionID,
		ServiceID:         service.ServiceID,
		PlanID:            service.Plan.ID,
		PlanName:          service.Plan.Name,
		PlanPrice:         service.Plan.Price,
		TransactionType:   transaction,
		TransactionDate:   time.Now(),
		TransactionStatus: service.Status,
	}
	_, err := mongo.ConnectApiMongoGlobal.InsertOne(context.Background(), "logs", logs)

	if err != nil {
		log.Println(err)
	}

	return err
}
