package controller_genre

import (
	"gin-gorm/database"
	"gin-gorm/models"
	"gin-gorm/requests"
	"gin-gorm/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGenre(ctx *gin.Context) {
	genres := new([]models.Genre)
	err := database.DB.Table("genres").Find(&genres).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Server Internal Error!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": genres,
	})
}

func GetGenreById(ctx *gin.Context){
	id := ctx.Param("id")

	genre := new(responses.GenreResponse)

	errDb := database.DB.Table("genres").Where("id = ?", id).Find(&genre).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if genre.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data transmitted succesfully",
		"data": genre,
	})
}

func CreateGenre (ctx *gin.Context) {
	genreReq := new(requests.GenreRequest)

	if errReq := ctx.ShouldBind(&genreReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	genreNameExist := new(models.Genre)
	errgenreNameExist := database.DB.Table("genres").Where("name = ?", genreReq.Name).Find(&genreNameExist).Error

	if errgenreNameExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Internal Error",
		})
		return
	}

	if genreNameExist.Name != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Genre already exist",
		})
		return
	}

	genre := new(models.Genre)
	genre.Name = &genreReq.Name

	errDb := database.DB.Table("genres").Create(&genre).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data saved succesfully",
		"data": genre,
	})
}

func UpdateGenreById(ctx *gin.Context){
	id := ctx.Param("id")
	genre := new(models.Genre)

	genreReq := new(requests.GenreRequest)
	genreNameExist := new(models.Genre)

	if errReq := ctx.ShouldBind(&genreReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDb := database.DB.Table("genres").Where("id = ?", id).Find(&genre).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if genre.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	// email exist
	errgenreNameExist := database.DB.Table("genres").Where("name = ?", genreReq.Name).Find(&genreNameExist).Error
	if errgenreNameExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Internal Error",
		})
		return
	}

	if genreNameExist.Name != nil && *genre.ID != *genreNameExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist",
		})
		return
	}

	genre.Name = &genreReq.Name

	errUpdate := database.DB.Table("genres").Where("id = ?", id).Updates(&genre).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't update data",
		})
		return
	}

	genreResponses := responses.GenreResponse{
		ID: genre.ID,
		Name: genre.Name,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data updated succesfully",
		"data": genreResponses,
	})
}

func DeleteGenreById(ctx *gin.Context){
	id := ctx.Param("id")
	genre := new(models.Genre)

	database.DB.Table("genres").Where("id = ?", id).Find(&genre)
	errDb := database.DB.Table("genres").Unscoped().Where("id = ?", id).Delete(&models.Genre{}).Error

	if genre.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error" : errDb.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data deleted succesfully",
	})
}