package controllers

import (
	"net/http"

	"github.com/tlksimba519/ptcg_trainerweb/database"
	"github.com/tlksimba519/ptcg_trainerweb/dto"
	"github.com/tlksimba519/ptcg_trainerweb/models"
	"github.com/tlksimba519/ptcg_trainerweb/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request")
		return
	}

	var user models.User
	if err := database.DB.Where("account = ?", req.Account).First(&user).Error; err != nil {
		utils.Error(c, http.StatusUnauthorized, "User not found")
		return
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		utils.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Name)

	res := dto.LoginResponse{
		Token: token,
	}

	utils.Success(c, res)
}

func Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request")
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Hash error")
		return
	}

	user := models.User{
		Account:  req.Account,
		Password: hashedPassword,
		Name:     req.Name,
		Email:    req.Email,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Create failed")
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Name)

	res := dto.LoginResponse{
		Token: token,
	}

	utils.Success(c, res)
}
