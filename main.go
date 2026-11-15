package main

import (
	"img-uploader/controllers"
	"img-uploader/initalizers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnv()
	initalizers.ConnectDB()
	initalizers.SyncData()
}

func main() {
	app := gin.Default()
	app.POST("/upload", controllers.UploadImage)
	app.GET("/images/:id", controllers.GetImage)
	app.Run(os.Getenv("PORT"))
}
