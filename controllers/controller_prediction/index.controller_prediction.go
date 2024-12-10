package controller_prediction

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"

	"gin-gorm/database"
	"gin-gorm/models"
)

func UploadAndPredictAI(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	// Create the /public/files directory if it doesn't exist
	publicDir := "./public/files"
	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		os.MkdirAll(publicDir, os.ModePerm)
	}

	// Save the uploaded file to the /public/files directory
	filePath := filepath.Join(publicDir, file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	// Open the saved file
	imgFile, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
		return
	}
	defer imgFile.Close()

	// Decode the image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode the image"})
		return
	}

	// Resize the image to the required size (e.g., 224x224)
	resizedImg := resize.Resize(224, 224, img, resize.Lanczos3)

	// Convert the image to a buffer
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, resizedImg, nil); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to encode the image"})
		return
	}

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create form file"})
		return
	}
	part.Write(buf.Bytes())
	writer.Close()

	// Send the image to the Python service for prediction
	req, err := http.NewRequest("POST", "http://localhost:5000/predict", body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to call prediction service"})
		return
	}
	defer resp.Body.Close()

	// Read the response from the Python service
	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode prediction response"})
		return
	}

	// Extract the label with the highest score
	var bestLabel string
	var bestScore float64
	for _, prediction := range result {
		label, ok := prediction["label"].(string)
		if !ok {
			continue
		}
		score, ok := prediction["score"].(float64)
		if !ok {
			continue
		}
		if score > bestScore {
			bestLabel = label
			bestScore = score
		}
	}

	if bestLabel == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No valid prediction found"})
		return
	}

	// Save the prediction result to the database
	predict := &models.Predictions{
		ImagePath: &filePath,
		Label:     &bestLabel,
		Score:     &bestScore,
	}
	errDb := database.DB.Table("predictions").Create(predict).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"makeBy": bestLabel, "score": bestScore})
}
	