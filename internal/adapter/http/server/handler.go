package server

import "github.com/gin-gonic/gin"

func (ct Controller) payments(c *gin.Context) {
	c.JSON(200, gin.H{"message" : "ok"})
}

func (ct Controller) paymentsSummary(c* gin.Context) {
	c.JSON(200, gin.H{"message" : "ok"})

}