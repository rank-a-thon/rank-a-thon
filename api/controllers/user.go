package controllers

import (
	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.UserModel)

func getUserID(context *gin.Context) (userID int64) {

	tokenAuth, err := authModel.ExtractTokenMetadata(context.Request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first."})
		return 0
	}
	userID, err = authModel.FetchAuth(tokenAuth)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first."})
		return 0
	}

	return userID
}

func (ctrl UserController) Login(context *gin.Context) {
	var loginForm forms.LoginForm

	if context.ShouldBindJSON(&loginForm) != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		context.Abort()
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"message": "User signed in", "user": user, "token": token})
	} else {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details", "error": err.Error()})
	}

}

func (ctrl UserController) Register(context *gin.Context) {
	var registerForm forms.RegisterForm

	if context.ShouldBindJSON(&registerForm) != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		context.Abort()
		return
	}

	user, err := userModel.Register(registerForm)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	if user.ID > 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
	} else {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not register this user", "error": err.Error()})
	}

}

func (ctrl UserController) Logout(context *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(context.Request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}
	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
