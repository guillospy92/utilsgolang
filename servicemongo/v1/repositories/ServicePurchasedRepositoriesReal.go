package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"guihub.com/guillospy92/servicemongo/v1/mongo"
)

const nameCollection = "service_purchased_original"

// ServicePurchasedDocumentRepository that implements interface ServicePurchasedRepositoryInterface
type ServicePurchasedDocumentRepository struct{}

// FindServicePurchased find collection in services_purchased by subscriptionID
// if this service not found a services_purchased return error
func (s *ServicePurchasedDocumentRepository) FindServicePurchased(subscriptionID string) error {
	filter := bson.M{"subscription_id": subscriptionID}
	return mongo.ConnectApiMongoGlobal.FindOne(context.Background(), nameCollection, filter).Err()
}

// SaveServicePurchased save new collection in services_purchased
func (s *ServicePurchasedDocumentRepository) SaveServicePurchased(service ServicePurchasedOriginal) error {
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
