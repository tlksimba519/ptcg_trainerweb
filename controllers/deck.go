package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tlksimba519/ptcg_trainerweb/database"
	"github.com/tlksimba519/ptcg_trainerweb/dto"
	"github.com/tlksimba519/ptcg_trainerweb/models"
	"github.com/tlksimba519/ptcg_trainerweb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDecks(c *gin.Context) {
	userID := c.GetUint("userID")

	// 建立 filter 條件
	filter := bson.M{"user_id": userID}

	// 查詢該使用者所有牌組
	cursor, err := database.DeckCollection.Find(context.TODO(), filter)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "無法取得牌組資料")
		return
	}
	defer cursor.Close(context.TODO())

	// 將結果 decode 成 models.Deck 切片
	var decks []models.Deck
	if err := cursor.All(context.TODO(), &decks); err != nil {
		utils.Error(c, http.StatusInternalServerError, "解析牌組資料失敗")
		return
	}

	// 整理回傳格式
	var res []dto.DeckResponse
	for _, d := range decks {
		res = append(res, dto.DeckResponse{
			ID:          d.ID.Hex(), // ← 這邊是 string 型別
			Title:       d.Title,
			Description: d.Description,
			Cards:       d.Cards,
		})
	}

	utils.Success(c, res)
}

func GetDeckByID(c *gin.Context) {
	userID := c.GetUint("userID")
	var deck models.Deck
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&deck).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Deck not found")
		return
	}

	res := dto.DeckResponse{ID: deck.ID.Hex(), Title: deck.Title, Description: deck.Description}
	utils.Success(c, res)
}

func CreateDeck(c *gin.Context) {
	userID := c.GetUint("userID")

	var input dto.DeckCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	deck := models.Deck{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Cards:       input.Cards,
		Tags:        []string{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := database.DeckCollection.InsertOne(c.Request.Context(), deck)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to insert deck")
		return
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	utils.Success(c, gin.H{"id": insertedID})
}

func UpdateDeck(c *gin.Context) {
	userID := c.GetUint("userID")
	var deck models.Deck
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&deck).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Deck not found")
		return
	}

	var req dto.DeckUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	deck.Title = req.Title
	deck.Description = req.Description
	database.DB.Save(&deck)

	utils.Success(c, "Deck updated")
}

func DeleteDeck(c *gin.Context) {
	userID := c.GetUint("userID")
	var deck models.Deck
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&deck).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Deck not found")
		return
	}

	database.DB.Delete(&deck)
	utils.Success(c, "Deck deleted")
}
