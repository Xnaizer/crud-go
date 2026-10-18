package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	// inisiasi router utama (js: const app = express())
	r := gin.Default()

	// declare endpoint
	r.GET("/health", func(c *gin.Context) {
		// response
		c.JSON(http.StatusOK, gin.H {
			"message": "Server is running",
		})
	})

	// run the server
	r.Run(":8080")

}
