package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

type AuthController struct{}

var authModel = new(models.AuthModel)

//
func (ctl AuthController) TokenValid(context *gin.Context) {
	err := authModel.TokenValid(context.Request)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
		context.Abort()
		return
	}
}

func verifyIDToken(ctx context.Context, app *firebase.App, idToken string) *auth.Token {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	log.Printf("Verified ID token: %v\n", token)
	return token
}

func (ctl AuthController) Refresh(context *gin.Context) {
	var tokenForm forms.Token

	if context.ShouldBindJSON(&tokenForm) != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "form": tokenForm})
		context.Abort()
		return
	}

	// Verify the token
	token, err := jwt.Parse(tokenForm.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	// If there is an error, the token must have expired
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
		return
	}
	// Is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
		return
	}
	// Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) // The token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string) // Convert the interface to string
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
			return
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
			return
		}
		// Delete the previous Refresh Token
		deleted, delErr := authModel.DeleteAuth(refreshUUID)
		if delErr != nil || deleted == 0 { // if any goes wrong
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
			return
		}

		// Create new pairs of refresh and access tokens
		ts, createErr := authModel.CreateToken(uint(userID))
		if createErr != nil {
			context.JSON(http.StatusForbidden, gin.H{"message": "Invalid authorization, please login again"})
			return
		}
		// Save the tokens metadata to redis
		saveErr := authModel.CreateAuth(uint(userID), ts)
		if saveErr != nil {
			context.JSON(http.StatusForbidden, gin.H{"message": "Invalid authorization, please login again"})
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		context.JSON(http.StatusOK, tokens)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
	}
}
