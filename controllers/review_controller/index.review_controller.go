package review_controller

func GetReviewById(ctx *gin.Context){
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


	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data transmitted succesfully",
		"data": animeResponse,
	})
}