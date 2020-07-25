package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/gin-gonic/gin"

	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

type UserController struct{}

var userModel = new(models.UserModel)

// Get User ID from auth token
func getUserID(context *gin.Context) (userID uint) {
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

func getUser(ctx context.Context, app *firebase.App, idToken string) *auth.UserRecord {
	// [START get_user_golang]
	// Get an auth client from the firebase.App
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	user, err := client.GetUser(ctx, idToken)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", idToken, err)
	}
	log.Printf("Successfully fetched user data: %v\n", user)
	// [END get_user_golang]
	return user
}

func getUserByEmail(ctx context.Context, client *auth.Client) *auth.UserRecord {
	email := "some@email.com"
	// [START get_user_by_email_golang]
	user, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		log.Fatalf("error getting user by email %s: %v\n", email, err)
	}
	log.Printf("Successfully fetched user data: %v\n", user)
	// [END get_user_by_email_golang]
	return user
}

func createUser(ctx context.Context, client *auth.Client) *auth.UserRecord {
	// [START create_user_golang]
	params := (&auth.UserToCreate{}).
		Email("user@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550100").
		Password("secretPassword").
		DisplayName("John Doe").
		Disabled(false)
	user, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", user)
	// [END create_user_golang]
	return user
}

func updateUser(ctx context.Context, client *auth.Client) {
	uid := "d"
	// [START update_user_golang]
	params := (&auth.UserToUpdate{}).
		Email("user@example.com").
		EmailVerified(true).
		PhoneNumber("+15555550100").
		Password("newPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(true)
	u, err := client.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
	// [END update_user_golang]
}

func deleteUser(ctx context.Context, client *auth.Client) {
	uid := "d"
	// [START delete_user_golang]
	err := client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	log.Printf("Successfully deleted user: %s\n", uid)
	// [END delete_user_golang]
}

// Login with email and password
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

// Register with name, email and password
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
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not register this user"})
	}

}

// Logout
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

// Get user details for logged in user
func (ctrl UserController) One(context *gin.Context) {

	_, err := authModel.ExtractTokenMetadata(context.Request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		context.Abort()
		return
	}
	if userID := getUserID(context); userID != 0 {
		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"user": user})
	}
}

// Get user details by user ID
func (ctrl UserController) GetByUserID(context *gin.Context) {

	_, err := authModel.ExtractTokenMetadata(context.Request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		context.Abort()
		return
	}
	var userID uint
	userID64, err := strconv.ParseUint(context.Query("userid"), 10, 64)
	if err == nil {
		userID = uint(userID64)
	} else {
		userID = getUserID(context)
	}

	if userID != 0 {
		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"user": user})
	}
}
