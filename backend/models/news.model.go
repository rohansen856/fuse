package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type ArticleStatus string

const (
	Published ArticleStatus = "published"
	Archived  ArticleStatus = "archived"
)

type News struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title         string             `bson:"title" json:"title"`
	Content       string             `bson:"content" json:"content"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
	Comments      []Comment          `bson:"comments,omitempty" json:"comments,omitempty"`
	UpvoteNumber  int                `bson:"upvote_number" json:"upvote_number" validate:"required,min=0"`
	DownvoteNumber int               `bson:"downvote_number" json:"downvote_number" validate:"required,min=0"`
	PublisherID   primitive.ObjectID `bson:"publisher_id" json:"publisher_id" validate:"required"`
	Category      string             `bson:"category" json:"category" validate:"required,min=2,max=100"`
	Status        ArticleStatus      `bson:"status" json:"status" validate:"required,oneof='published' 'archived'"`
}
