package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

type TeamController struct{}

var teamModel = new(models.TeamModel)
var teamInviteModel = new(models.TeamInviteModel)

func (ctrl TeamController) Create(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		var teamForm forms.TeamForm

		event, err := fetchAndValidateEvent(context)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid event name"})
			context.Abort()
			return
		}

		if context.ShouldBindJSON(&teamForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}

		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		} else if models.JsonStringToStringUintMap(user.TeamIDForEvent)[event] != 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Team could not be created, user already has team"})
			context.Abort()
			return
		}

		teamID, err := teamModel.Create(userID, teamForm, models.Event(event))
		if teamID == 0 && err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Team could not be created", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Team created", "id": teamID})
	}
}

func (ctrl TeamController) One(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		event, err := fetchAndValidateEvent(context)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid event name"})
			context.Abort()
			return
		}
		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		}

		teamID := models.JsonStringToStringUintMap(user.TeamIDForEvent)[event]
		if teamID == 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Team does not have team"})
			context.Abort()
			return
		}

		data, err := teamModel.One(teamID)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Team not found", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (ctrl TeamController) All(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		data, err := teamModel.All(userID)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Teams not found", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (ctrl TeamController) Update(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		event, err := fetchAndValidateEvent(context)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid event name"})
			context.Abort()
			return
		}

		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		}
		teamID := models.JsonStringToStringUintMap(user.TeamIDForEvent)[event]
		if teamID == 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Team does not have team"})
			context.Abort()
			return
		}

		var teamForm forms.TeamForm
		if context.ShouldBindJSON(&teamForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}

		err = teamModel.Update(teamID, teamForm)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Team could not be updated", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Team updated"})
	}
}

func (ctrl TeamController) SendInvite(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		}

		var teamInviteForm forms.TeamInviteForm
		if context.ShouldBindJSON(&teamInviteForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}

		if !models.Events[teamInviteForm.Event] {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Event does not exist"})
			context.Abort()
			return
		}

		teamID := userModel.GetTeamIDForEventMap(user)[teamInviteForm.Event]
		if teamID == 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "User does not have team"})
			context.Abort()
			return
		}

		invitedUser, err := userModel.GetUserByEmail(teamInviteForm.Email)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		}

		invitedUserTeamID := userModel.GetTeamIDForEventMap(invitedUser)[teamInviteForm.Event]
		if invitedUserTeamID == teamID {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Invited user already belongs to team"})
			context.Abort()
			return
		}

		inviteID, err := teamInviteModel.Create(invitedUser.ID, teamID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Team invited could not be created", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Team invite created", "invite_id": inviteID})
	}
}

func (ctrl TeamController) GetInvites(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		data, err := teamInviteModel.All(userID)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Team invites not found", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (ctrl TeamController) AcceptInvite(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		if teamID, err := strconv.ParseUint(context.Query("teamid"), 10, 64); err == nil {
			err = teamModel.AddTeamMember(userID, uint(teamID))
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not add team member", "error": err.Error()})
				context.Abort()
				return
			}

			err = teamInviteModel.Delete(userID, uint(teamID))
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete team invite", "error": err.Error()})
				context.Abort()
				return
			}
			context.JSON(http.StatusOK, gin.H{"message": "Team invite accepted"})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
			context.Abort()
			return
		}

	}
}

func (ctrl TeamController) DeclineInvite(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		if teamID, err := strconv.ParseUint(context.Query("teamid"), 10, 64); err == nil {
			err = teamInviteModel.Delete(userID, uint(teamID))
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete team invite", "error": err.Error()})
				context.Abort()
			}
			context.JSON(http.StatusOK, gin.H{"message": "Team invite declined"})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
			context.Abort()
			return
		}
	}
}

func (ctrl TeamController) RemoveTeamMember(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		if deleteUserID, err := strconv.ParseUint(context.Query("delete-user-id"), 10, 64); err == nil {
			event, err := fetchAndValidateEvent(context)
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid event name"})
				context.Abort()
				return
			}
			user, err := userModel.One(userID)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
				context.Abort()
				return
			}
			teamID := models.JsonStringToStringUintMap(user.TeamIDForEvent)[event]
			if teamID == 0 {
				context.JSON(http.StatusNotAcceptable, gin.H{"message": "User does not have team"})
				context.Abort()
				return
			}

			err = teamModel.RemoveTeamMember(uint(deleteUserID), teamID)
			err = userModel.UpdateTeamForUser(uint(deleteUserID), 0, models.Event(event))
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Team member could not be removed", "error": err.Error()})
				context.Abort()
				return
			}
			context.JSON(http.StatusOK, gin.H{"message": "Team member removed"})
		}
	}
}

func (ctrl TeamController) Delete(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		event, err := fetchAndValidateEvent(context)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid event name"})
			context.Abort()
			return
		}
		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		}
		teamID := models.JsonStringToStringUintMap(user.TeamIDForEvent)[event]
		if teamID == 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "User does not have team"})
			context.Abort()
			return
		}

		err = teamModel.Delete(teamID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Team could not be deleted", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
	}
}

func getTeamIDForEvent(context *gin.Context, userID uint) (teamID uint, err error){
	event, err := fetchAndValidateEvent(context)
	if err != nil {
		return 0, errors.New("invalid event name")
	}

	user, err := userModel.One(userID)
	if err != nil {
		return 0, errors.New("unable to fetch user")
	}
	teamIDForEvent := userModel.GetTeamIDForEventMap(user)
	return teamIDForEvent[event], nil
}