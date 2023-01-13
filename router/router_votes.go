package router

import (
	"bharvest-vo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TxVotesRes struct {
	Type            string            `json:"type"`
	UpdateTimestamp string            `json:"updateTimestamp"`
	RawData         []models.TxVotesT `json:"rawData"`
}

func getVotes(c *gin.Context) {
	items, err := models.GetVotes()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		results := []TxVotesRes{}
		tempResult := TxVotesRes{Type: "B-Harvest", RawData: items}
		if len(items) > 0 {
			fmt.Println(len(items))
			fmt.Println(items[0].Timestamp)
			tempResult.UpdateTimestamp = items[0].Timestamp
		}
		results = append(results, tempResult)
		tempResultForCrescent := TxVotesRes{Type: "Crescent"}
		results = append(results, tempResultForCrescent)
		c.JSON(http.StatusOK, gin.H{
			"data": results,
		})
	}
}
