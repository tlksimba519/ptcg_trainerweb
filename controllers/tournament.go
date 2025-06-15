package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tlksimba519/ptcg_trainerweb/utils"
)

type Tournament struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Level    string    `json:"level"`
	Notes    string    `json:"notes"`
	Date     time.Time `json:"date"`
}

func RegisterTournament(c *gin.Context) {

	event := Tournament{
		Name:     "夏季盃全國大賽",
		Location: "台北市大安運動中心",
		Level:    "高階",
		Notes:    "請攜帶身分證與選手證，限滿16歲參加",
		Date:     time.Date(2025, 7, 13, 9, 0, 0, 0, time.Local),
	}

	var res []Tournament
	res = append(res, event)
	utils.Success(c, res)
}
