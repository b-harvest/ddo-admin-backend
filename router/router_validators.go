package router

import (
	"bharvest-vo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TxRewardPerValidatorRes struct {
	Type            string                            `json:"type"`
	UpdateTimestamp string                            `json:"updateTimestamp"`
	RawData         []models.TxRewardPerValidatorResT `json:"rawData"`
}

func getValidators(c *gin.Context) {
	items, err := models.GetValidators()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		results := []TxRewardPerValidatorRes{}
		tempResult := TxRewardPerValidatorRes{Type: "B-Harvest", RawData: items}
		if len(items) > 0 {
			fmt.Println(len(items))
			fmt.Println(items[0].Time)
			tempResult.UpdateTimestamp = items[0].Time
		}
		results = append(results, tempResult)
		tempResultForCrescent := TxRewardPerValidatorRes{Type: "Crescent"}
		results = append(results, tempResultForCrescent)
		c.JSON(http.StatusOK, gin.H{
			"data": results,
		})
	}
}
