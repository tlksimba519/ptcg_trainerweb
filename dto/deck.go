package dto

import "github.com/tlksimba519/ptcg_trainerweb/models"

type DeckResponse struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	UserID      uint                  `json:"user_id"`
	Description string                `json:"description,omitempty"`
	Cards       []models.DeckCardItem `json:"cards,omitempty"`
}

type DeckCreateRequest struct {
	Title       string                `bson:"title" json:"title"`
	Description string                `bson:"description" json:"description"`
	Cards       []models.DeckCardItem `bson:"cards" json:"cards"`
}

type DeckUpdateRequest struct {
	UserID      uint                  `bson:"user_id" json:"user_id"`
	Title       string                `bson:"title" json:"title"`
	Description string                `bson:"description" json:"description"`
	Cards       []models.DeckCardItem `bson:"cards" json:"cards"`
	Tags        []string              `bson:"tags" json:"tags"`
}
