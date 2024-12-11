package controller_anime

import (
	"fmt"
	"gin-gorm/database"
	"gin-gorm/models"
	"gin-gorm/requests"
	"gin-gorm/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAnime(ctx *gin.Context) {
	animes := new([]models.Anime)
	err := database.DB.Table("animes").Find(&animes).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Server Internal Error!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success Get All Anime",
		"data": animes,
	})
}

func GetAnimeById(ctx *gin.Context){
	id := ctx.Param("id")

	anime := new(responses.AnimeResponse)

	errDb := database.DB.Table("animes").Where("id = ?", id).Find(&anime).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if anime.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	// Muat data genre
	var genre models.Genre
	if err := database.DB.First(&genre, anime.GenreId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Failed to load genre",
	})
		return
	}

	// Mapping hasil query ke AnimeResponse
    animeResponse := responses.AnimeResponse{
        ID:       anime.ID,
        Title:    anime.Title,
        GenreId:  anime.GenreId,
        Episodes: anime.Episodes,
		Sinopsis: anime.Sinopsis,
        Genre: responses.GenreResponse{
            ID:   genre.ID,
            Name: genre.Name,
        },
    }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data transmitted succesfully",
		"data": animeResponse,
	})
}

func CreateAnime (ctx *gin.Context) {
	animeReq := new(requests.AnimeRequest)

	if errReq := ctx.ShouldBind(&animeReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	animeTitleExist := new(models.Anime)
	erranimeTitleExist := database.DB.Table("animes").Where("title = ?", animeReq.Title).Find(&animeTitleExist).Error

	if erranimeTitleExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Internal Error",
		})
		return
	}

	if animeTitleExist.Title != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Anime already exist",
		})
		return
	}

	anime := new(models.Anime)
	anime.Title = &animeReq.Title
	anime.GenreId = &animeReq.GenreId
	anime.Episodes = &animeReq.Episodes
	anime.Sinopsis = &animeReq.Sinopsis
	anime.CreatedAt = &animeReq.CreatedAt
	anime.UpdatedAt = &animeReq.UpdatedAt

	errDb := database.DB.Table("animes").Create(&anime).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// Query dengan Preload untuk memuat relasi Genre
	if err := database.DB.Preload("Genre").Where("id = ?", anime.GenreId).Find(&anime).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data saved succesfully",
		"data": anime,
	})
	
}

func UpdateAnimeById(ctx *gin.Context){
	id := ctx.Param("id")
	anime := new(models.Anime)

	animeReq := new(requests.AnimeRequest)
	animeTitleExist := new(models.Anime)

	if errReq := ctx.ShouldBind(&animeReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDb := database.DB.Table("animes").Where("id = ?", id).Find(&anime).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if anime.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	// email exist
	erranimeTitleExist := database.DB.Table("animes").Where("title = ?", animeReq.Title).Find(&animeTitleExist).Error
	if erranimeTitleExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Internal Error",
		})
		return
	}

	if animeTitleExist.Title != nil && *anime.ID != *animeTitleExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Anime already exist",
		})
		return
	}

	anime.Title = &animeReq.Title
	anime.GenreId = &animeReq.GenreId
	anime.Episodes = &animeReq.Episodes
	anime.Sinopsis = &animeReq.Sinopsis
	anime.CreatedAt = &animeReq.CreatedAt
	anime.UpdatedAt = &animeReq.UpdatedAt

	errUpdate := database.DB.Table("animes").Where("id = ?", id).Updates(&anime).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't update data",
		})
		return
	}

	
	animeResponses := responses.AnimeResponse{
		ID: anime.ID,
		Title: anime.Title,
		GenreId: anime.GenreId,
		Sinopsis : anime.Sinopsis,
		Episodes: anime.Episodes,
		
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data updated succesfully",
		"data": animeResponses,
	})
}

func DeleteAnimeById(ctx *gin.Context){
	id := ctx.Param("id")
	anime := new(models.Anime)

	database.DB.Table("animes").Where("id = ?", id).Find(&anime)
	errDb := database.DB.Table("animes").Unscoped().Where("id = ?", id).Delete(&models.Anime{}).Error

	if anime.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Anime Not Found",
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
		"message": "Anime deleted succesfully",
	})
}

func GetAllUserPaginate(ctx *gin.Context){
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}

	perPage := ctx.Query("perPage")
	if perPage == ""{
		perPage = "10"
	}

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 1 {
		pageInt = 1
	}

	var totalRecords int64
	database.DB.Table("anime").Count(&totalRecords)
	totalPages := int((totalRecords + int64(perPageInt) - 1) / int64(perPageInt))

	if pageInt > totalPages {
		pageInt = totalPages
	}

	anime := new([]models.User)
	err := database.DB.Table("anime").Offset((pageInt-1)*perPageInt).Limit(perPageInt).Find(&anime).Error
	if(err != nil) {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Server Internal Error!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data" : anime,
		"page" : pageInt,
		"per_page" : perPageInt,
	})
}

func GetAllAnimePaginate(ctx *gin.Context){
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}

	perPage := ctx.Query("perPage")
	if perPage == ""{
		perPage = "10"
	}

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 1 {
		pageInt = 1
	}

	var totalRecords int64
	database.DB.Table("animes").Count(&totalRecords)
	totalPages := int((totalRecords + int64(perPageInt) - 1) / int64(perPageInt))

	if pageInt > totalPages {
		pageInt = totalPages
	}

	anime := new([]models.Anime)
	err := database.DB.Table("animes").Offset((pageInt-1)*perPageInt).Limit(perPageInt).Find(&anime).Error
	if(err != nil) {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Server Internal Error!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data" : anime,
		"page" : pageInt,
		"per_page" : perPageInt,
	})
}

func UploadTumbnailAnime(ctx *gin.Context) {
	
	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "File not found",
		})
		return
	}

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