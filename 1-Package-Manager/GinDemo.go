package __Package_Manager

import "github.com/gin-gonic/gin"

func CreateGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
			"success": true,
		})
	})
	r.Run(":8080")
}
