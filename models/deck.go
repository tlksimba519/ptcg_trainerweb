package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deck struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      uint               `bson:"user_id" json:"user_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Cards       []DeckCardItem     `bson:"cards" json:"cards"`
	Tags        []string           `bson:"tags" json:"tags"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type DeckCardItem struct {
	CardID   string `bson:"card_id" json:"card_id"`
	Quantity int    `bson:"quantity" json:"quantity"`
	Note     string `bson:"note,omitempty" json:"note"`
}
