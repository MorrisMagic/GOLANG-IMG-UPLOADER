package controllers

import (
	"img-uploader/initalizers"
	"img-uploader/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 6<<20) //limit how data server accepet , 6<<20 6mb

	file, err := c.FormFile("image") // file name , file : information about the file , err
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Image is required",
		})
		return
	}

	if file.Size > 5<<20 { //check the size
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Image must be smaller than 5MB",
		})
		return
	}

	openedFile, err := file.Open() // gives us access to the actual file contents
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot open image",
		})
		return
	}
	defer openedFile.Close() //defer mean exucate when the function finishis

	imageBytes, err := io.ReadAll(openedFile) //convert img into bytes
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot read image",
		})
		return
	}

	contentType := http.DetectContentType(imageBytes) // img type png jpg webp

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}

	if !allowedTypes[contentType] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Only JPEG, PNG and WebP images are allowed",
		})
		return
	}

	image := models.IMG{ //add the img
		Name: file.Filename,
		Type: contentType,
		Data: imageBytes,
	}

	result := initalizers.DB.Create(&image) //add img

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Image uploaded successfully",
		"id":      image.ID,
	})
}

func GetImage(c *gin.Context) {
	id := c.Param("id")

	var image models.IMG

	result := initalizers.DB.First(&image, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Image not found",
		})
		return
	}

	c.Data(http.StatusOK, image.Type, image.Data)
}
