package controllers

import (
	"File-Management-System/database"
	"File-Management-System/models"
	"context"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var fileTransactionRecords = database.OpenCollection(database.Client, "files")

func getFileFormat(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the first 512 bytes to determine the file type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Determine the file format using DetectContentType
	fileType := http.DetectContentType(buffer)

	return fileType, nil
}

func getFileById(fileID string) (*models.FileRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var record models.FileRecord

	// Retrieve file information from the database using the file ID
	err := fileTransactionRecords.FindOne(ctx, bson.M{"file_id": fileID}).Decode(&record)

	if err != nil {
		return nil, err
	}
	return &record, nil
}

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
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

		fileId := primitive.NewObjectID()

		// Save the file to the local system
		filePath := "user_file_uploads/" + fileId.Hex()
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileFormat, err := getFileFormat(filePath)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		record := models.FileRecord{
			ID:        fileId,
			FileId:    fileId.Hex(),
			UserID:    userId.(string),
			FileName:  file.Filename,
			Timestamp: time.Now(),
			FileSize:  file.Size,
			Format:    fileFormat,
		}

		_, insertErr := fileTransactionRecords.InsertOne(ctx, record)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		c.JSON(http.StatusOK, record)
	}

}

func DownloadFile() gin.HandlerFunc {
	return func(c *gin.Context) {

		fileID := c.Param("file_id")

		record, err := getFileById(fileID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Construct the file path
		filePath := "user_file_uploads/" + record.FileId

		// Check if the file exists
		_, err = os.Stat(filePath)
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		// Set the Content-Disposition header to suggest a filename for the browser
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", record.FileName))

		// Serve the file
		c.File(filePath)
	}
}

func DeleteFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get file ID from request parameters
		fileID := c.Param("file_id")

		// Retrieve file information from the database using the file ID
		_, err := getFileById(fileID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		// Delete the file from local storage
		filePath := "user_file_uploads/" + fileID
		err = os.Remove(filePath)
		if err != nil {
			fmt.Println("here")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Delete the file details from MongoDB
		_, err = fileTransactionRecords.DeleteOne(ctx, bson.M{"fileid": fileID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file details from database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
	}
}

func RetrieveFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		userId, exists := c.Get("uid")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User id not found"})
			return
		}

		var records []models.FileRecord
		cursor, err := fileTransactionRecords.Find(ctx, bson.M{"user_id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		err = cursor.All(ctx, &records)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve files"})
			return
		}

		// Sort files by size in ascending order
		sort.Slice(records, func(i, j int) bool {
			return records[i].FileSize < records[j].FileSize
		})

		c.JSON(http.StatusOK, records)
	}
}
