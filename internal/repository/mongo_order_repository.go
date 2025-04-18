package repository

import (
	"context"
	"github.com/facelessEmptiness/order_service/internal/domain"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoOrderRepo struct {
	coll *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) OrderRepository {
	return &mongoOrderRepo{coll: db.Collection("orders")}
}

func (r *mongoOrderRepo) Create(o *domain.Order) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"user_id":        o.UserID,
		"items":          o.Items,
		"payment_method": o.PaymentMethod,
		"status":         o.Status,
	}
	res, err := r.coll.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}
	oid := res.InsertedID.(primitive.ObjectID).Hex()
	return oid, nil
}

func (r *mongoOrderRepo) GetByID(id string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, _ := primitive.ObjectIDFromHex(id)
	var o domain.Order
	if err := r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&o); err != nil {
		return nil, err
	}
	return &o, nil
}
