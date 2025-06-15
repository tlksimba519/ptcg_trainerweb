package dto

type UserResponse struct {
	ID      uint   `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}
