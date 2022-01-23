package repositories

import "time"

// ServicePurchasedOriginal attributes of the service purchased for saved in documentDb.
type ServicePurchasedOriginal struct {
	SubscriptionID string    `json:"subscription_id" bson:"subscription_id"`
	ServiceID      string    `json:"service_id" bson:"service_id"`
	ServiceType    string    `json:"service_type" bson:"service_type"`
	CountryCode    string    `json:"country_code" bson:"country_code"`
	CityCode       string    `json:"city_code" bson:"city_code"`
	Timezone       string    `json:"timezone" bson:"timezone"`
	ActivationDate time.Time `json:"activation_date" bson:"activation_date"`
	ExpirationDate time.Time `json:"expiration_date "bson:"expiration_date"`
	IsRenovated    bool      `json:"is_renovated" bson:"is_renovated"`
	Status         string    `json:"status" bson:"status"`
	User           struct {
		ID        int    `json:"id" bson:"id"`
		Name      string `json:"name" bson:"name"`
		Email     string `json:"email" bson:"email"`
		Cellphone string `json:"cellphone" bson:"cellphone"`
		Address   string `json:"address" bson:"address"`
	} `bson:"user"`
	Plan struct {
		ID               string   `json:"id" bson:"id"`
		Name             string   `json:"name" bson:"name"`
		Price            float64  `json:"price" bson:"price"`
		PromotionalPrice float64  `json:"promotional_price" bson:"promotional_price"`
		Benefits         []string `json:"benefits" bson:"benefits"`
		Period           int      `json:"period", bson:"period"`
	} `bson:"plan"`
}

// ServicePurchasedChargeOriginal attributes of the service purchased for saved in documentDb
type ServicePurchasedChargeOriginal struct {
	SubscriptionID string `json:"subscription_id" bson:"subscription_id"`
	IsRenovated    bool   `json:"is_renovated" bson:"is_renovated"`
	Payment        struct {
		CardID       int    `json:"card_id" bson:"card_id"`
		Installments int    `json:"installments" bson:"installments"`
		CardType     string `json:"card_type" bson:"card_type"`
		LastFour     int32  `json:"last_four" bson:"last_four"`
	} `bson:"payment"`
}

// ServicePurchasedCancelOriginal attributes of the service cancel for saved in documentDb
type ServicePurchasedCancelOriginal struct {
	SubscriptionID string    `json:"subscription_id" bson:"subscription_id"`
	IsRenovated    bool      `json:"is_renovated" bson:"is_renovated"`
	CancelDate     time.Time `json:"cancel_date" bson:"cancel_date"`
	Status         string    `json:"status" bson:"status"`
}
