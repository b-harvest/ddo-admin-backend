package router

import (
	"bharvest-vo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TxProposalsRes struct {
	Type            string                `json:"type"`
	UpdateTimestamp string                `json:"updateTimestamp"`
	RawData         []models.TmProposalsT `json:"rawData"`
}

func getProposals(c *gin.Context) {
	items, err := models.GetTmProposals()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		results := []TxProposalsRes{}
		tempResult := TxProposalsRes{Type: "B-Harvest", RawData: items}
		if len(items) > 0 {
			fmt.Println(len(items))
			fmt.Println(items[0].Time)
			tempResult.UpdateTimestamp = items[0].Time
		}
		results = append(results, tempResult)
		tempResultForCrescent := TxProposalsRes{Type: "Crescent"}
		results = append(results, tempResultForCrescent)
		c.JSON(http.StatusOK, gin.H{
			"data": results,
		})
	}
}
