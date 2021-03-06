package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"

	"golang.org/x/crypto/bcrypt"
)

// Team ...
type User struct {
	gorm.Model
	Email          string   `gorm:"column:email;not null;unique" json:"email"`
	Password       string   `gorm:"column:password" json:"-"`
	Name           string   `gorm:"column:name" json:"name"`
	UserType       UserType `gorm:"column:user_type;default:0" json:"user_type"`
	TeamIDForEvent string   `gorm:"column:team_id_for_event;default:'{}'" json:"team_id_for_event"`
}

// UserModel ...
type UserModel struct{}

var authModel = new(AuthModel)

var judgeEmails = readEmailsFromCsv("secrets/judges.csv")
var superUserEmails = readEmailsFromCsv("secrets/superusers.csv")

// Login ...
func (m UserModel) Login(form forms.LoginForm) (user User, token Token, err error) {
	err = database.GetDB().
		Table("public.users").Where("email = ?", strings.ToLower(form.Email)).Take(&user).Error

	if err != nil {
		return user, token, err
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, errors.New("invalid password")
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		return user, token, err
	}

	err = authModel.CreateAuth(user.ID, tokenDetails)
	if err == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, err
}

// Register ...
func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
	db := database.GetDB()

	// Check if the user exists in database NOTE this will fail if user is deleted and marked with deleted at column

	var count int
	err = db.Table("public.users").
		Where("email = ?", strings.ToLower(form.Email)).Select("count(id)").Count(&count).Error

	if err != nil {
		return user, err
	}

	if count > 0 {
		return user, errors.New("user already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err) //Something really went wrong here...
	}

	//Create the user and return back the user ID
	var userType UserType
	if m.isEmailFromSuperUser(form.Email) {
		userType = 3
	} else if m.isEmailFromJudge(form.Email) {
		userType = 2
	} else {
		userType = 0
	}
	user = User{Email: form.Email, Password: string(hashedPassword), Name: form.Name, UserType: userType}
	err = db.Table("public.users").Create(&user).Scan(&user).Error

	return user, err
}

// One ...
func (m UserModel) One(userID uint) (user User, err error) {
	err = database.GetDB().Table("public.users").
		Where("id = ?", userID).Take(&user).Error
	return user, err
}

func (m UserModel) GetUserByEmail(email string) (user User, err error) {
	err = database.GetDB().Table("public.users").
		Where("email = ?", email).Take(&user).Error
	return user, err
}

func (m UserModel) UpdateTeamForUser(userID, teamID uint, event Event) (err error) {
	user, err := m.One(userID)

	if err != nil {
		return errors.New(fmt.Sprintf("user %d not found", userID))
	}

	teamIDForEvent := JsonStringToStringUintMap(user.TeamIDForEvent)
	teamIDForEvent[string(event)] = teamID

	err = database.GetDB().Table("public.users").Model(&User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"team_id_for_event": StringUintMapToJsonString(teamIDForEvent),
		}).Error
	return err
}

func (m UserModel) GetTeamIDForEventMap(user User) map[string]uint {
	return JsonStringToStringUintMap(user.TeamIDForEvent)
}

func (m UserModel) IsJudgeForUserID(id uint) (bool, error) {
	user, err := userModel.One(id)
	if err != nil {
		return false, err
	}
	if user.UserType == Judge {
		return true, nil
	} else {
		return false, nil
	}
}

func (m UserModel) GetAllJudges() (judges []User, err error) {
	var users []User
	err = database.GetDB().Table("public.users").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	judges = make([]User, 0, 100)
	for _, user := range users {
		if m.isEmailFromJudge(user.Email) {
			judges = append(judges, user)
		}
	}
	return judges, err
}

func (m UserModel) isEmailFromJudge(email string) bool {
	return judgeEmails[email]
}

func (m UserModel) isEmailFromSuperUser(email string) bool {
	return superUserEmails[email]
}
