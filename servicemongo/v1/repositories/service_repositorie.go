package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"guihub.com/guillospy92/servicemongo/v1/mongo"
)

// ServiceRepositories struct
type ServiceRepositories struct{}

// User struct
type User struct {
	Id        int64  `bson:"id"`
	Name      string `bson:"name"`
	Email     string `bson:"email"`
	Address   string `bson:"address"`
	CellPhone int64  `bson:"cellPhone"`
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
	CardID       int64 `bson:"card_id"`
	Installments int   `bson:"installments"`
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

// ServicePurchased struct
type ServicePurchased struct {
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
	var service ServicePurchased

	filter := bson.D{
		// {"plan.benefits", bson.D{{"$eq", "free_delivery"}}},
		{"plan.benefits", "free_delivery, netflix"},
		{"country_code", "co"},
		{"user.id", 234},
		{"status", "active"},
	}

	err := mongo.ConnectApiMongoGlobal.FindOne(context.Background(), "service_purchased", filter).Decode(&service)

	if err != nil {
		return ServicePurchased{}, err
	}

	return service, err
}
