package file_controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleUploadFile(ctx *gin.Context) {
	
	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "File not found",
		})
		return
	}

	// file, errFile := fileHeader.Open()

	// if errFile != nil {
	// 	panic(errFile)
	// }

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", fileHeader.Filename))
	if errUpload != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed to upload file",
		})
		return
	}


	ctx.JSON(200, gin.H{
		"message": "File uploaded successfully",
	})
}