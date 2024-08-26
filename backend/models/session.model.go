package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SessionToken string             `bson:"session_token" json:"session_token" validate:"required"`
	Device       []string           `bson:"device" json:"device" validate:"required"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Expires      time.Time          `bson:"expires" json:"expires" validate:"required"`
}
