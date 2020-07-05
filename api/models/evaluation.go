package models

import (
    "errors"
    "fmt"

    "github.com/jinzhu/gorm"
    "github.com/rank-a-thon/rank-a-thon/api/database"
    "github.com/rank-a-thon/rank-a-thon/api/forms"
)

// Evaluation ...
type Evaluation struct {
    gorm.Model
    JudgeID    	      	    uint  `gorm:"column:judge_id;not null" json:"judge_id"`
    SubmissionID            uint  `gorm:"column:submission_id;not null" json:"submission_id"`
    // Rankings are integers 1-10 when set
    MainRating              uint  `gorm:"column:main_rating;default:0" json:"main_rating"`
    AnnoyingRating          uint  `gorm:"column:annoying_rating;default:0" json:"annoying_rating"`
    EntertainRating         uint  `gorm:"column:entertaining_rating;default:0" json:"entertaining_rating"`
    BeautifulRating         uint  `gorm:"column:beautiful_rating;default:0" json:"beautiful_rating"`
    SociallyUsefulRating    uint  `gorm:"column:socially_useful_rating;default:0" json:"socially_useful_rating"`
    HardwareRating          uint  `gorm:"column:hardware_rating;default:0" json:"hardware_rating"`
    AwesomelyUselessRating  uint  `gorm:"column:awesomely_useless_rating;default:0" json:"awesomely_useless_rating"`
}

// EvaluationModel ...
type EvaluationModel struct{}

// Create ...
// When this is created, the ratings are not assigned yet and are 0
func (m EvaluationModel) Create(judgeID uint, submissionID uint) (evaluationID uint, err error) {
    evaluation := Evaluation{JudgeID: judgeID, SubmissionID: submissionID}
    err = database.GetDB().Table("public.evaluations").Create(&evaluation).Error
    return evaluation.ID, err
}

// One ...
func (m EvaluationModel) One(id uint) (evaluation Evaluation, err error) {
    err = database.GetDB().Table("public.evaluations").
        Where("evaluations.id = ?", id).
        Take(&evaluation).Error
    return evaluation, err
}

// Get all evaluations assigned to a judge
func (m EvaluationModel) AllForJudge(judgeID uint) (evaluations []Evaluation, err error) {
    err = database.GetDB().Table("public.evaluations").
        Where("evaluations.judge_id = ?", judgeID).
        Order("evaluations.id desc").
        Find(&evaluations).Error
    return evaluations, err
}

// Get all evaluations assigned to a submission
func (m EvaluationModel) AllForSubmission(judgeID uint) (evaluations []Evaluation, err error) {
    err = database.GetDB().Table("public.evaluations").
        Where("evaluations.judge_id = ?", judgeID).
        Order("evaluations.id desc").
        Find(&evaluations).Error
    return evaluations, err
}

// Update ...
func (m EvaluationModel) Update(id uint, form forms.EvaluationForm) (err error) {
    _, err = m.One(id)
    if err != nil {
        return errors.New(fmt.Sprintf("error updating: evaluation %d not found", id))
    }
    err = database.GetDB().Table("public.evaluations").Model(&Evaluation{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "MainRating": form.MainRating,
            "AnnoyingRating": form.AnnoyingRating,
            "EntertainRating": form.EntertainRating,
            "BeautifulRating": form.BeautifulRating,
            "SociallyUsefulRating": form.SociallyUsefulRating,
            "HardwareRating": form.HardwareRating,
            "AwesomelyUselessRating": form.AwesomelyUselessRating}).Error
    return err
}

// Delete ...
func (m EvaluationModel) Delete(id uint) (err error) {
    _, err = m.One(id)

    if err != nil {
        return errors.New(fmt.Sprintf("error deleting: evaluation %d not found", id))
    }
    err = database.GetDB().Table("public.evaluations").Where("id = ?", id).Delete(Evaluation{}).Error

    return err
}

func (m EvaluationModel) HaveAllEvaluationsCompleted() (result bool, err error) {
    return false, nil
}
