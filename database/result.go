package database

import "go.mongodb.org/mongo-driver/mongo"

type DocumentResult struct {
	Res *mongo.SingleResult
}
type Result interface {
	Decode(document interface{}) error
}

func (r *DocumentResult) Decode(document interface{}) error {
	err := r.Res.Decode(document)
	return err
}
