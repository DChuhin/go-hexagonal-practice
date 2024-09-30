package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseMongoRepository[T any] struct {
	collection *mongo.Collection
}

func New[T any](collection *mongo.Collection) *BaseMongoRepository[T] {
	return &BaseMongoRepository[T]{
		collection: collection,
	}
}

func (r *BaseMongoRepository[T]) InsertOne(ctx context.Context, document *T) (string, error) {
	result, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return result.InsertedID.(string), nil
	}

	return insertedID.Hex(), nil
}

func (r *BaseMongoRepository[T]) Find(ctx context.Context, filter bson.D) ([]*T, error) {
	var results []*T
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *BaseMongoRepository[T]) FindById(ctx context.Context, id string) (*T, error) {
	return r.FindOne(ctx, idFilter(id))
}

func (r *BaseMongoRepository[T]) FindOne(ctx context.Context, filter bson.D) (*T, error) {
	var result T

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BaseMongoRepository[T]) UpdateById(ctx context.Context, id string, update bson.D, createIfNotExist bool) (bool, error) {
	return r.UpdateOne(ctx, idFilter(id), update, createIfNotExist)
}

func (r *BaseMongoRepository[T]) UpdateOne(ctx context.Context, filter bson.D, update bson.D, createIfNotExist bool) (bool, error) {
	opts := options.Update().SetUpsert(createIfNotExist)
	result, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}

func (r *BaseMongoRepository[T]) ReplaceById(ctx context.Context, id string, document *T, createIfNotExist bool) (bool, error) {
	return r.ReplaceOne(ctx, idFilter(id), document, createIfNotExist)
}

func (r *BaseMongoRepository[T]) ReplaceOne(ctx context.Context, filter bson.D, document *T, createIfNotExist bool) (bool, error) {
	opts := options.Replace().SetUpsert(createIfNotExist)
	result, err := r.collection.ReplaceOne(ctx, filter, document, opts)
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}

func (r *BaseMongoRepository[T]) DeleteById(ctx context.Context, id string) (bool, error) {
	return r.DeleteOne(ctx, idFilter(id))
}

func (r *BaseMongoRepository[T]) DeleteOne(ctx context.Context, filter bson.D) (bool, error) {
	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return res.DeletedCount > 0, nil
}

func idFilter(id string) bson.D {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return bson.D{{"_id", id}}
	}
	return bson.D{{"_id", objectId}}
}
