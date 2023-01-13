package router

import (
	"bharvest-vo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TxCryptoFlowRes struct {
	Type            string                 `json:"type"`
	UpdateTimestamp string                 `json:"updateTimestamp"`
	RawData         []models.TxCryptoFlowT `json:"rawData"`
}

func getAllTxCryptoFlow(c *gin.Context) {
	txs, err := models.GetAllTxCryptoFlow()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// c.IndentedJSON()
		results := []TxCryptoFlowRes{}
		tempResult := TxCryptoFlowRes{Type: "B-Harvest", RawData: txs}
		if len(txs) > 0 {
			fmt.Println(len(txs))
			fmt.Println(txs[0].Timestamp)
			tempResult.UpdateTimestamp = txs[0].Timestamp
		}
		results = append(results, tempResult)
		tempResultForCrescent := TxCryptoFlowRes{Type: "Crescent"}
		results = append(results, tempResultForCrescent)
		c.JSON(http.StatusOK, gin.H{
			"data": results,
		})
	}
}
