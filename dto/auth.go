package dto

type RegisterRequest struct {
	Account  string `gorm:"uniqueIndex;size:32" json:"account"`
	Password string `gorm:"size:128" json:"-"`
	Name     string `gorm:"size:50" json:"name"`
	Email    string `gorm:"uniqueIndex;size:100" json:"email"`
	Role     string `gorm:"default:'user'" json:"role"`
}

type LoginRequest struct {
	Account  string `gorm:"uniqueIndex;size:32" json:"account"`
	Password string `gorm:"size:128" json:"-"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
