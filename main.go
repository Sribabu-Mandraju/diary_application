package main

import (
	"backend/database"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	client := database.DBconnection()
	defer client.Disconnect(context.Background())

		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "server created successfully",
			})
		})

	router.Run(":3000")

}
