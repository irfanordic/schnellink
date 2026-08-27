package main
import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func main(){
	r :=  gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/shorten", func(c *gin.Context){
		
        var req shortenRequestURL

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request payload",
			})
			return
		}

		shortCodeLength := 6
		shortCode := GenerateShortCode(shortCodeLength)

		c.JSON(http.StatusOK, gin.H{
			"short_url": shortCode,
			"original_url": req.URL,
		})
	})

	r.Run(":8081")
}