package router

import (
	"bharvest-vo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BalancePerAddressRes struct {
	Type            string                      `json:"type"`
	UpdateTimestamp string                      `json:"updateTimestamp"`
	RawData         []models.BalancePerAddressT `json:"rawData"`
}

func getBalancePerAddress(c *gin.Context) {
	balances, err := models.GetBalancePerAddress()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// c.IndentedJSON()
		result := []BalancePerAddressRes{}
		temp := BalancePerAddressRes{Type: "B-Harvest", RawData: balances}
		if len(balances) > 0 {
			temp.UpdateTimestamp = balances[0].Time
		}
		result = append(result, temp)
		tempForCrescent := BalancePerAddressRes{Type: "Crescent"}
		result = append(result, tempForCrescent)
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
