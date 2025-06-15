package models

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	// 通用
	ID          uint   `gorm:"primaryKey" json:"id"`
	CardID      string `gorm:"uniqueIndex" json:"card_id"`
	Name        string `json:"name"`
	CardType    string `json:"card_type"`
	Series      string `json:"series"`
	Rarity      string `json:"rarity"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`

	// 寶可夢
	HP            int    `json:"hp,omitempty"`
	Type          string `json:"type,omitempty"`
	Weakness      string `json:"weakness,omitempty"`
	Resisitance   string `json:"resistance,omitempty"`
	RetreatCost   int    `json:"retreat_cost,omitempty"`
	Attack1Name   string `json:"attack1_name,omitempty"`
	Attack1Cost   int    `json:"attack1_cost,omitempty"`
	Attack1Damage int    `json:"attack1_damage,omitempty"`
	Attack1Text   string `json:"attack1_text,omitempty"`
	Attack2Name   string `json:"attack2_name,omitempty"`
	Attack2Cost   int    `json:"attack2_cost,omitempty"`
	Attack2Damage int    `json:"attack2_damage,omitempty"`
	Attack2Text   string `json:"attack2_text,omitempty"`

	// 訓練家
	TrainerSubtype string `json:"trainer_subtype,omitempty"`

	// 能量
	EnergyType string `json:"energy_type,omitempty"`
	IsSpecial  bool   `json:"is_special,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
