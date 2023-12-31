package routes

import (
	"File-Management-System/controllers"
	"File-Management-System/middleware"

	"github.com/gin-gonic/gin"
)

func FileRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/upload", middleware.Authenticate(), controllers.UploadFile())
	incomingRoutes.GET("/download/:file_id", middleware.Authenticate(), controllers.DownloadFile())
	incomingRoutes.DELETE("/upload/:file_id", middleware.Authenticate(), controllers.DeleteFile())
	incomingRoutes.GET("/files", middleware.Authenticate(), controllers.RetrieveFiles())
	incomingRoutes.GET("/transactions", middleware.Authenticate(), controllers.AllTransactions())

}
