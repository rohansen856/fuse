package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type News struct {
	Category  string `bson:"category" json:"category"`     
	Datetime  time.Time  `bson:"datetime" json:"datetime"`     
	Headline  string `bson:"headline" json:"headline"`     
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                
	Image     string `bson:"image" json:"image"`           
	Related   string `bson:"related" json:"related"`       
	Source    string `bson:"source" json:"source"`        
	Summary   string `bson:"summary" json:"summary"`      
	URL       string `bson:"url" json:"url"`              
}

