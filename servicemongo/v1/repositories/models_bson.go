package repositories

import "time"

// ServicePurchasedOriginal attributes of the service purchased for saved in documentDb.
type ServicePurchasedOriginal struct {
	ID             string               `json:"id" bson:"_id,omitempty"`
	SubscriptionID string               `json:"subscription_id" bson:"subscription_id"`
	ServiceID      string               `json:"service_id" bson:"service_id"`
	ServiceType    string               `json:"service_type,omitempty"`
	CountryCode    string               `json:"country_code" bson:"country_code"`
	City           City                 `json:"city" bson:"city"`
	Timezone       string               `json:"timezone" bson:"timezone"`
	ActivationDate *time.Time           `json:"activation_date" bson:"activation_date"`
	ExpirationDate *time.Time           `json:"expiration_date" bson:"expiration_date"`
	CancelDate     *time.Time           `json:"cancel_date" bson:"cancel_date"`
	IsRenovated    bool                 `json:"is_renovated" bson:"is_renovated"`
	Status         string               `json:"status" bson:"status"` // active, pending_payment, cancel
	Source         Source               `json:"source" bson:"source"`
	User           User                 `json:"user" bson:"user"`
	Plan           ServicePurchasedPlan `json:"plan" bson:"plan"`
	Payment        Payment              `json:"payment" bson:"payment"`
}

// ServicePurchasedPlan attributes of the plan associated with the service.
type ServicePurchasedPlan struct {
	ID            string   `json:"id" bson:"id"`
	Name          string   `json:"name" bson:"name"`
	Price         float64  `json:"price" bson:"price"`
	OriginalPrice float64  `json:"original_price" bson:"original_price"`
	Benefits      []string `json:"benefits" bson:"benefits"`
	Period        int      `json:"period" bson:"period"`
}

// City attributes of the user city.
type City struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}

// Source attributes of the source.
type Source struct {
	DeviceID string `json:"device_id" bson:"device_id"`
}

// ServicePurchasedChargeOriginal attributes of the service purchased for saved in documentDb
type ServicePurchasedChargeOriginal struct {
	SubscriptionID string `json:"subscription_id" bson:"subscription_id"`
	IsRenovated    bool   `json:"is_renovated" bson:"is_renovated"`
	Payment        Payment
}

// ServicePurchasedCancelOriginal attributes of the service cancel for saved in documentDb
type ServicePurchasedCancelOriginal struct {
	SubscriptionID string    `json:"subscription_id" bson:"subscription_id"`
	IsRenovated    bool      `json:"is_renovated" bson:"is_renovated"`
	CancelDate     time.Time `json:"cancel_date" bson:"cancel_date"`
	Status         string    `json:"status" bson:"status"`
}
