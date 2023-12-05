package filecontrollers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		userId, exists := c.Get("uid")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User id not found"})
			return

		}
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the file to the local system
		filePath := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			fmt.Println("Here2")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file_path": filePath})
	}

}
