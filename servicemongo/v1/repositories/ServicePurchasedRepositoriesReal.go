package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"guihub.com/guillospy92/servicemongo/v1/mongo"
)

const nameCollection = "services_purchased"

// ServicePurchasedDocumentRepository that implements interface ServicePurchasedRepositoryInterface
type ServicePurchasedDocumentRepository struct{}

// FindServicePurchased find collection in services_purchased by subscriptionID
// if this service not found a services_purchased return error
func (s *ServicePurchasedDocumentRepository) FindServicePurchased(subscriptionID string) (map[string]interface{}, error) {
	var a map[string]interface{}
	
	filter := bson.M{"subscription_id": subscriptionID}
	return a, mongo.ConnectApiMongoGlobal.FindOne(context.Background(), nameCollection, filter).Decode(&a)
}

// SaveServicePurchased save new collection in services_purchased
func (s *ServicePurchasedDocumentRepository) SaveServicePurchased(service ServicePurchasedMer) error {
	_, err := mongo.ConnectApiMongoGlobal.InsertOne(context.Background(), nameCollection, service)
	return err
}

// SaveServiceCharge save new collection in services_purchased
func (s *ServicePurchasedDocumentRepository) SaveServiceCharge(service ServicePurchasedChargeOriginal) error {

	filter := bson.M{"subscription_id": service.SubscriptionID}
	update := bson.M{"$set": service}
	_, err := mongo.ConnectApiMongoGlobal.UpdateOne(context.Background(), nameCollection, filter, update)

	return err
}

// SaveServiceCancel save new collection in services_purchased
func (s *ServicePurchasedDocumentRepository) SaveServiceCancel(service ServicePurchasedCancelOriginal) error {

	fmt.Println(service)

	filter := bson.M{"subscription_id": service.SubscriptionID}
	update := bson.M{"$set": service}
	_, err := mongo.ConnectApiMongoGlobal.UpdateOne(context.Background(), nameCollection, filter, update)

	return err
}
