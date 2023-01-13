package router

import (
	"bharvest-vo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	r.GET("/products", getProducts)
	r.GET("/accounts/bharvest", getAccountsBharvest)
	r.GET("/balances", getBalancePerAddress)
	r.GET("/txs", getAllTxCryptoFlow)
	r.GET("/validators", getValidators)
	r.GET("/votes", getVotes)
	r.GET("/proposals", getProposals)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":  "pong",
	// 		"rankType": "랭크",
	// 	})
	// })
}

func getProducts(c *gin.Context) {
	products := models.GetProducts()

	if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}
