package router

import (
	"bharvest-vo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Temp struct {
	RankType        string                     `json:"rankType"`
	UpdateTimestamp int64                      `json:"updateTimestamp"`
	RankData        []models.AccountsBharvestT `json:"rankData"`
}

func getAccountsBharvest(c *gin.Context) {
	accounts, err := models.GetAccountsBharvestT()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// c.IndentedJSON()
		result := []Temp{}
		temp := Temp{RankType: "B-Harvest", RankData: accounts, UpdateTimestamp: 3}
		result = append(result, temp)
		temp.RankType = "Crescent"
		result = append(result, temp)
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
