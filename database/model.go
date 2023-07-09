package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type ModelClient struct {
	Col *mongo.Collection
}
type Model interface {
	Create(ctx context.Context, doc interface{}) error
	FindById(ctx context.Context, id string, schema interface{}) (*DocumentResult, error)
	FindOne(ctx context.Context, filter interface{}) (*DocumentResult, error)
	Find(ctx context.Context, filter interface{}, schema interface{}) error
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) (*DocumentResult, error)
	FindDocById(ctx context.Context, id string) (*DocumentResult, error)
}

func GetId(doc interface{}) primitive.ObjectID {
	val := reflect.ValueOf(doc)
	field := val.FieldByName("ID")
	if reflect.Value.IsNil(field) {
		return primitive.NewObjectID()
	}
	return field.Interface().(primitive.ObjectID)
}

func setDefaults(doc interface{}) interface{} {
	id := reflect.ValueOf(doc).Elem().FieldByName("ID")
	id.Set(reflect.ValueOf(primitive.NewObjectID()))
	createdAt := reflect.ValueOf(doc).Elem().FieldByName("CreatedAt")
	createdAt.Set(reflect.ValueOf(time.Now().UTC()))
	updatedAt := reflect.ValueOf(doc).Elem().FieldByName("UpdatedAt")
	updatedAt.Set(reflect.ValueOf(time.Now().UTC()))
	return doc
}

func (m *ModelClient) Create(ctx context.Context, doc interface{}) error {
	d := setDefaults(doc)
	_, err := m.Col.InsertOne(ctx, d)
	return err
}

func (m *ModelClient) FindById(ctx context.Context, id string) *DocumentResult {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic("Invalid object id.")
	}
	s := m.Col.FindOne(ctx, bson.M{
		"_id": _id,
	})
	return &DocumentResult{
		Res: s,
	}
}

func (m *ModelClient) FindOne(ctx context.Context, filter bson.M) *DocumentResult {

	doc := m.Col.FindOne(ctx, filter)
	return &DocumentResult{
		Res: doc,
	}
}

func (m *ModelClient) Find(ctx context.Context, filter interface{}, schema interface{}) error {

	cur, err := m.Col.Find(ctx, filter)
	if err != nil {
		return err
	}
	err = cur.All(ctx, schema)

	return err
}

func (m *ModelClient) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) (*mongo.SingleResult, error) {
	res := m.Col.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return nil, res.Err()
	}
	return res, nil
}

func (m *ModelClient) FindDocById(ctx context.Context, id string) (*DocumentResult, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic("Invalid object id.")
	}
	s := m.Col.FindOne(ctx, bson.M{
		"_id": _id,
	})
	return &DocumentResult{
		Res: s,
	}, nil
}
