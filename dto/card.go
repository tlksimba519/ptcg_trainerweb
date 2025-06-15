package dto

type CardResponse struct {
	ID       uint   `json:"id"`
	CardID   string `json:"card_id"`
	Name     string `json:"name"`
	CardType string `json:"card_type"`
	Series   string `json:"series"`
	Rarity   string `json:"rarity"`
	ImageURL string `json:"image_url"`
}
