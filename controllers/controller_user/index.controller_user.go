package controller_user

import (
	"gin-gorm/database"
	"gin-gorm/models"
	"gin-gorm/requests"
	"gin-gorm/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	users := new([]models.User)
	err := database.DB.Table("users").Find(&users).Error
	if(err != nil) {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Server Internal Error!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data" : users,
	})
}

func GetUserById(ctx *gin.Context){
	id := ctx.Param("id")

	user := new(responses.UserResponse)

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data transmitted succesfully",
		"data": user,
	})
}

func StoreUser (ctx *gin.Context) {
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	userEmailExist := new(models.User)
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error

	if errUserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Internal Error",
		})
		return
	}

	if userEmailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist",
		})
		return
	}

	user := new(models.User)
	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	user.BornDate = &userReq.BornDate

	errDb := database.DB.Table("users").Create(&user).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data saved succesfully",
		"data": user,
	})
}

func UpdateUserById(ctx *gin.Context){
	id := ctx.Param("id")
	user := new(models.User)

	userReq := new(requests.UserRequest)
	userEmailExist := new(models.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	// email exist
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Internal Error",
		})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist",
		})
		return
	}

	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't update data",
		})
		return
	}

	userResponses := responses.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data updated succesfully",
		"data": userResponses,
	})
}

func DeleteUserById(ctx *gin.Context){
	id := ctx.Param("id")
	user := new(models.User)

	database.DB.Table("users").Where("id = ?", id).Find(&user)
	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.User{}).Error

	if user.ID == nil {
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
	database.DB.Table("users").Count(&totalRecords)
	totalPages := int((totalRecords + int64(perPageInt) - 1) / int64(perPageInt))

	if pageInt > totalPages {
		pageInt = totalPages
	}

	users := new([]models.User)
	err := database.DB.Table("users").Offset((pageInt-1)*perPageInt).Limit(perPageInt).Find(&users).Error
	if(err != nil) {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "Server Internal Error!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data" : users,
		"page" : pageInt,
		"per_page" : perPageInt,
	})
}

