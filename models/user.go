package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Account   string `gorm:"uniqueIndex;size:32" json:"account"` // 登入帳號
	Password  string `gorm:"size:128" json:"-"`                  // 加密密碼，不回傳
	Name      string `gorm:"size:50" json:"name"`                // 暱稱 or 顯示名稱
	Email     string `gorm:"uniqueIndex;size:100" json:"email"`  // 信箱
	Role      string `gorm:"default:'user'" json:"role"`         // 權限：user / admin 等
	CreatedAt time.Time
	UpdatedAt time.Time
}
