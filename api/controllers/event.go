package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

// get event name from param in url, check if is valid event
func fetchAndValidateEvent(context *gin.Context) (event string, err error) {
	event = context.Param("event")
	if !models.Events[event] {
		return event, errors.New("invalid event")
	}
	return event, nil
}
