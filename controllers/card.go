package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tlksimba519/ptcg_trainerweb/database"
	"github.com/tlksimba519/ptcg_trainerweb/dto"
	"github.com/tlksimba519/ptcg_trainerweb/models"
	"github.com/tlksimba519/ptcg_trainerweb/utils"
)

func GetAllCards(c *gin.Context) {
	var cards []models.Card
	if err := database.DB.Find(&cards).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "DB error")
		return
	}

	var res []dto.CardResponse
	for _, card := range cards {
		res = append(res, dto.CardResponse{
			ID:       card.ID,
			CardID:   card.CardID,
			Name:     card.Name,
			CardType: card.CardType,
			Series:   card.Series,
			Rarity:   card.Rarity,
			ImageURL: card.ImageURL,
		})
	}

	utils.Success(c, res)
}

func GetCardByID(c *gin.Context) {

	var cards []models.Card
	if err := database.DB.Where("CardID = ?", c.GetUint("cardID")).First(&cards).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "DB error")
		return
	}

	var res []dto.CardResponse
	for _, card := range cards {
		res = append(res, dto.CardResponse{
			ID:       card.ID,
			CardID:   card.CardID,
			Name:     card.Name,
			CardType: card.CardType,
			Series:   card.Series,
			Rarity:   card.Rarity,
			ImageURL: card.ImageURL,
		})
	}

	utils.Success(c, res)
}
