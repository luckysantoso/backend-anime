package controller_reviews

import (
	"gin-gorm/database"
	"gin-gorm/models"
	"gin-gorm/requests"
	"gin-gorm/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReviewsByAnimeId(ctx *gin.Context) {
	id := ctx.Param("anime_id")

	var reviews []responses.ReviewsResponse

	errDb := database.DB.Table("reviews").Where("anime_id = ?", id).Find(&reviews).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// if len(reviews) == 0 {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"message": "Data Not Found",
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data transmitted successfully",
		"data":    reviews,
	})
}

func CreateReview(ctx *gin.Context) {
	reviewReq := new(requests.ReviewRequest)

	if errReq := ctx.ShouldBind(&reviewReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	review := new(models.Review)
	review.UserId = &reviewReq.UserId
	review.AnimeId = &reviewReq.AnimeId
	review.Message = &reviewReq.Message

	errDb := database.DB.Table("reviews").Create(&review).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data saved succesfully",
		"data":    review,
	})

}
