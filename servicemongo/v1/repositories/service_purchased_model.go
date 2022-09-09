package repositories

import "time"

const (
	// ActiveStatus status when the service subscription is active for the user.
	ActiveStatus string = "active"
	// PendingPaymentStatus status when the service subscription is pending payment.
	PendingPaymentStatus string = "pending_payment"
	// CancelStatus status when the service subscription was canceled.
	CancelStatus string = "cancel"
)

// ServicePurchasedCollection collection name.
const ServicePurchasedCollection string = "services_purchased"

// ServicePurchasedMer attributes of the service purchased by the user.
type ServicePurchasedMer struct {
	ID             string     `json:"id" bson:"_id,omitempty"`
	SubscriptionID string     `json:"subscription_id" bson:"subscription_id"`
	ServiceID      string     `json:"service_id" bson:"service_id"`
	ServiceType    string     `json:"service_type,omitempty" bson:"service_type"`
	CountryCode    string     `json:"country_code" bson:"country_code"`
	CityCode       string     `json:"city_code,omitempty" bson:"city_code"`
	Timezone       string     `json:"timezone" bson:"timezone"`
	ActivationDate *time.Time `json:"activation_date" bson:"activation_date"`
	ExpirationDate *time.Time `json:"expiration_date" bson:"expiration_date"`
	CancelDate     *time.Time `json:"cancel_date" bson:"cancel_date"`
	IsRenovated    bool       `json:"is_renovated" bson:"is_renovated"`
	Status         string     `json:"status" bson:"status"` // active, pending_payment, cancel
	User           UserMer    `json:"user" bson:"user"`
	Plan           PlanMer    `json:"plan" bson:"plan"`
	Payment        PaymentMer `json:"payment" bson:"payment"`
}

// UserMer attributes of the user.
type UserMer struct {
	ID             int    `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	Cellphone      string `json:"cellphone" bson:"cellphone"`
	Address        string `json:"address" bson:"address"`
	PostalCode     string `json:"postal_code,omitempty" bson:"postal_code"`
	IdentityNumber string `json:"identity_number,omitempty" bson:"identity_number"`
	IdentityType   string `json:"identity_type,omitempty" bson:"identity_type"`
}

// PaymentMer attributes of the payment used to buy service.
type PaymentMer struct {
	CardID       int    `json:"card_id" bson:"card_id"`
	Installments int    `json:"installments" bson:"installments"`
	CardType     string `json:"card_type,omitempty" bson:"card_type"`
	LastFour     string `json:"last_four,omitempty" bson:"last_four"`
}

type ServiceMer struct {
	ID             string    `json:"id" bson:"_id"`
	Name           string    `json:"name" bson:"name"`
	IsActive       bool      `json:"is_active" bson:"is_active"`
	Type           string    `json:"type" bson:"type"` // merqueo_prime
	CountryCode    string    `json:"country_code" bson:"country_code"`
	PaymentMethods []string  `json:"payment_methods" bson:"payment_methods"`
	Plans          []PlanMer `json:"plans" bson:"plans"`
}

type PlanMer struct {
	ID               string   `json:"id" bson:"_id"`
	Name             string   `json:"name,omitempty" bson:"name"`
	Price            float64  `json:"price" bson:"price"`
	PromotionalPrice *float64 `json:"promotional_price" bson:"promotional_price"`
	Benefits         []string `json:"benefits" bson:"benefits"`
	Period           int      `json:"period" bson:"period"`
}
