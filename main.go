package main

import (
	"file-loader/config"
	"file-loader/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	r := gin.Default()

	fileHandler := handlers.NewFileHandler()

	r.GET("/files", fileHandler.List)
	r.GET("/files/:name", fileHandler.Read)
	r.POST("/files", fileHandler.Write)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}

	r.Run(":" + port)
}
