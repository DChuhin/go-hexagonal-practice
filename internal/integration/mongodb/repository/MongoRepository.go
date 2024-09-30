package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepository[T any] interface {
	InsertOne(ctx context.Context, document *T) (string, error)

	Find(ctx context.Context, filter bson.D) ([]*T, error)
	FindById(ctx context.Context, id string) (*T, error)
	FindOne(ctx context.Context, filter bson.D) (*T, error)

	// UpdateById returns: true if there was matched and replaced document, false if there was no matching document, even if it was upserted
	UpdateById(ctx context.Context, id string, update bson.D, createIfNotExist bool) (bool, error)
	UpdateOne(ctx context.Context, filter bson.D, update bson.D, createIfNotExist bool) (bool, error)

	// ReplaceById returns: true if there was matched and replaced document, false if there was no matching document, even if it was upserted
	ReplaceById(ctx context.Context, id string, document *T, createIfNotExist bool) (bool, error)
	ReplaceOne(ctx context.Context, filter bson.D, document *T, createIfNotExist bool) (bool, error)

	// DeleteById returns: true if document was deleted
	DeleteById(ctx context.Context, id string) (bool, error)
	DeleteOne(ctx context.Context, filter bson.D) (bool, error)
}
