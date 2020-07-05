package models

import (
    "errors"
    "fmt"

    "github.com/jinzhu/gorm"
    "github.com/rank-a-thon/rank-a-thon/api/database"
)

type SubmissionRanking struct {
    gorm.Model
    SubmissionID            uint    `gorm:"column:submission_id;not null" json:"submission_id"`
    // Rankings are position from 1 to num_submissions
    MainRanking		        uint    `gorm:"column:main_ranking;default:0" json:"main_ranking"`
    AnnoyingRanking         uint    `gorm:"column:annoying_ranking;default:0" json:"annoying_ranking"`
    EntertainRanking        uint	`gorm:"column:entertaining_ranking;default:0" json:"entertaining_ranking"`
    BeautifulRanking        uint	`gorm:"column:beautiful_ranking;default:0" json:"beautiful_ranking"`
    SociallyUsefulRanking   uint    `gorm:"column:socially_useful_ranking;default:0" json:"socially_useful_ranking"`
    HardwareRanking    	    uint    `gorm:"column:hardware_ranking;default:0" json:"hardware_ranking"`
    AwesomelyUselessRanking uint    `gorm:"column:awesomely_useless_ranking;default:0" json:"awesomely_useless_ranking"`
}

type SubmissionRankingModel struct{}

// Create ...
// When this is created, the rankings are not assigned yet and are 0
func (m SubmissionRankingModel) Create(
    submissionID uint, mainRanking uint, annoyingRanking uint, entertainRanking uint, beautifulRanking uint,
    sociallyUsefulRanking uint, hardwareRanking uint, awesomelyUselessRanking uint) (submissionRankingID uint, err error) {

    submissionRanking := SubmissionRanking{
        SubmissionID:            submissionID,
        MainRanking:             mainRanking,
        AnnoyingRanking:         annoyingRanking,
        EntertainRanking:        entertainRanking,
        BeautifulRanking:        beautifulRanking,
        SociallyUsefulRanking:   sociallyUsefulRanking,
        HardwareRanking:         hardwareRanking,
        AwesomelyUselessRanking: awesomelyUselessRanking,
    }
    err = database.GetDB().Table("public.submission_rankings").Create(&submissionRanking).Error
    return submissionRanking.ID, err
}

// One ...
func (m SubmissionRankingModel) OneBySubmissionID(submissionID uint) (submissionRanking Evaluation, err error) {
    err = database.GetDB().Table("public.submission_rankings").
        Where("submission_rankings.submission_id = ?", submissionID).
        Take(&submissionRanking).Error
    return submissionRanking, err
}

// Get all evaluations assigned to a judge
func (m SubmissionRankingModel) All() (submissionRankings []SubmissionRanking, err error) {
    err = database.GetDB().Table("public.submission_rankings").
        Order("submission_rankings.id desc").
        Find(&submissionRankings).Error
    return submissionRankings, err
}

// Delete ...
func (m SubmissionRankingModel) DeleteBySubmissionID(submissionID uint) (err error) {
    _, err = m.OneBySubmissionID(submissionID)

    if err != nil {
        return errors.New(fmt.Sprintf("error deleting: submission ranking %d not found", submissionID))
    }
    err = database.GetDB().Table("public.submission_rankings").
        Where("submission_id = ?", submissionID).Delete(Evaluation{}).Error
    return err
}

