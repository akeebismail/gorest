package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `json:"email" json:"email"`
	Password  string             `json:"password" json:"password"`
	CreatedAt time.Time          `json:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" json:"updatedAt"`
}
