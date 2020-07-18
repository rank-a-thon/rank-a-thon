package models

import (
    "errors"
    "fmt"

    "github.com/jinzhu/gorm"
    "github.com/rank-a-thon/rank-a-thon/api/database"
)

type SubmissionRanking struct {
    gorm.Model
    SubmissionID            uint       `gorm:"column:submission_id;not null" json:"submission_id"`
    // Rankings are position from 1 to num_submissions
    MainRanking		        uint       `gorm:"column:main_ranking;default:0" json:"main_ranking"`
    AnnoyingRanking         uint       `gorm:"column:annoying_ranking;default:0" json:"annoying_ranking"`
    EntertainRanking        uint	   `gorm:"column:entertaining_ranking;default:0" json:"entertaining_ranking"`
    BeautifulRanking        uint       `gorm:"column:beautiful_ranking;default:0" json:"beautiful_ranking"`
    SociallyUsefulRanking   uint       `gorm:"column:socially_useful_ranking;default:0" json:"socially_useful_ranking"`
    HardwareRanking    	    uint       `gorm:"column:hardware_ranking;default:0" json:"hardware_ranking"`
    AwesomelyUselessRanking uint       `gorm:"column:awesomely_useless_ranking;default:0" json:"awesomely_useless_ranking"`

    MainScore		        float64    `gorm:"column:main_score;default:0" json:"main_score"`
    AnnoyingScore           float64    `gorm:"column:annoying_score;default:0" json:"annoying_score"`
    EntertainScore          float64	   `gorm:"column:entertaining_score;default:0" json:"entertaining_score"`
    BeautifulScore          float64    `gorm:"column:beautiful_score;default:0" json:"beautiful_score"`
    SociallyUsefulScore     float64    `gorm:"column:socially_useful_score;default:0" json:"socially_useful_score"`
    HardwareScore    	    float64    `gorm:"column:hardware_score;default:0" json:"hardware_score"`
    AwesomelyUselessScore   float64    `gorm:"column:awesomely_useless_score;default:0" json:"awesomely_useless_score"`
}

type SubmissionRankingModel struct{}

// Create ...
// When this is created, the rankings are not assigned yet and are 0
func (m SubmissionRankingModel) Create(
    submissionID uint, scoresArray []float64) (submissionRankingID uint, err error) {
    if len(scoresArray) != NumberOfRatings {
        return 0, errors.New("scores array length mismatch")
    }

    submissionRanking := SubmissionRanking{
        SubmissionID:          submissionID,
        MainScore:             scoresArray[0],
        AnnoyingScore:         scoresArray[1],
        EntertainScore:        scoresArray[2],
        BeautifulScore:        scoresArray[3],
        SociallyUsefulScore:   scoresArray[4],
        HardwareScore:         scoresArray[5],
        AwesomelyUselessScore: scoresArray[6],
    }
    err = database.GetDB().Table("public.submission_rankings").Create(&submissionRanking).Error
    return submissionRanking.ID, err
}

// Update ...
func (m SubmissionRankingModel) Update(submissionID uint, form map[string]uint) (err error) {
    _, err = m.OneBySubmissionID(submissionID)

    if err != nil {
        return errors.New(fmt.Sprintf("submission %d not found", submissionID))
    }
    err = database.GetDB().Table("public.submission_rankings").Model(&Team{}).
        Where("id = ?", submissionID).
        Updates(form).Error
    return err
}

// One ...
func (m SubmissionRankingModel) OneBySubmissionID(submissionID uint) (submissionRanking Evaluation, err error) {
    err = database.GetDB().Table("public.submission_rankings").
        Where("submission_rankings.submission_id = ?", submissionID).
        Take(&submissionRanking).Error
    return submissionRanking, err
}

// All ...
func (m SubmissionRankingModel) All() (submissionRankings []SubmissionRanking, err error) {
    err = database.GetDB().Table("public.submission_rankings").
        Order("submission_rankings.id desc").
        Find(&submissionRankings).Error
    return submissionRankings, err
}

// Sorted by category score descending order
func (m SubmissionRankingModel) AllByCategory(category string) (submissionRankings []SubmissionRanking, err error) {
    err = database.GetDB().Table("public.submission_rankings").
        Order(fmt.Sprintf("submission_rankings.%s_score desc", category)).
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

