package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        			primitive.ObjectID 					`bson:"_id,omitempty" json:"id,omitempty"`
	Title     			string             					`bson:"title" json:"title"`
	Content   			string            					`bson:"content" json:"content"`
	CreatedAt 			time.Time          					`bson:"created_at" json:"created_at"`
	UpdatedAt 			time.Time          					`bson:"updated_at" json:"updated_at"`
	Comments  			[]Comment          					`bson:"comments,omitempty" json:"comments,omitempty"`
	
}
