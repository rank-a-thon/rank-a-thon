package controllers

import (
	"strconv"

	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

var articleModel = new(models.ArticleModel)

func (ctrl ArticleController) Create(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		var articleForm forms.ArticleForm

		if context.ShouldBindJSON(&articleForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}

		articleID, err := articleModel.Create(userID, articleForm)

		if articleID == 0 && err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Article created", "id": articleID})
	}
}

func (ctrl ArticleController) All(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		data, err := articleModel.All(userID)

		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get articles", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (ctrl ArticleController) One(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			data, err := articleModel.One(userID, uint(id))

			if err != nil {
				context.JSON(http.StatusNotFound, gin.H{"Message": "Article not found", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"data": data})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		}
	}
}

func (ctrl ArticleController) Update(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			var articleForm forms.ArticleForm

			if context.ShouldBindJSON(&articleForm) != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
				context.Abort()
				return
			}

			err := articleModel.Update(userID, uint(id), articleForm)
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be updated", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"message": "Article updated"})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter", "error": err.Error()})
		}
	}
}

func (ctrl ArticleController) Delete(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			err := articleModel.Delete(userID, uint(id))
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be deleted", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"message": "Article deleted"})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		}
	}
}
