package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tlksimba519/ptcg_trainerweb/database"
	"github.com/tlksimba519/ptcg_trainerweb/dto"
	"github.com/tlksimba519/ptcg_trainerweb/models"
	"github.com/tlksimba519/ptcg_trainerweb/utils"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "DB error")
		return
	}

	var res []dto.UserResponse
	for _, user := range users {
		res = append(res, dto.UserResponse{
			ID:      user.ID,
			Account: user.Account,
			Name:    user.Name,
			Email:   user.Email,
		})
	}

	utils.Success(c, res)
}

func GetUserByID(c *gin.Context) {

	var users []models.User
	if err := database.DB.Where("ID = ?", c.GetUint("ID")).First(&users).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "DB error")
		return
	}

	var res []dto.UserResponse
	for _, user := range users {
		res = append(res, dto.UserResponse{
			ID:      user.ID,
			Account: user.Account,
			Name:    user.Name,
			Email:   user.Email,
		})
	}

	utils.Success(c, res)
}
