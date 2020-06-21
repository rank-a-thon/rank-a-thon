package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
)

// Article ...
type Article struct {
	gorm.Model
	UserID    uint  `gorm:"column:user_id;not null" json:"-"`
	Title     string `gorm:"column:title" json:"title"`
	Content   string `gorm:"column:content" json:"content"`
	User      User   `gorm:"column:user;foreignkey:UserID" json:"user"`
}

// ArticleModel ...
type ArticleModel struct{}

// Create ...
func (m ArticleModel) Create(userID uint, form forms.ArticleForm) (articleID uint, err error) {
	article := Article{UserID: userID, Title: form.Title, Content: form.Content}
	err = database.GetDB().Table("public.articles").Create(&article).Error
	return article.ID, err
}

// One ...
func (m ArticleModel) One(userID, id uint) (article Article, err error) {
	err = database.GetDB().Preload("User").Table("public.articles").
		Where("articles.user_id = ? AND articles.id = ?", userID, id).
		//Select("articles.id, articles.title, articles.content, articles.updated_at, articles.created_at").
		Joins("left join public.users on articles.user_id = users.id").
		Take(&article).Error
	return article, err
}

// All ...
func (m ArticleModel) All(userID uint) (articles []Article, err error) {
	err = database.GetDB().Preload("User").Table("public.articles").
		Where("articles.user_id = ?", userID).
		Joins("left join public.users on articles.user_id = users.id").
		Order("articles.id desc").
		Find(&articles).Error
	return articles, err
}

// Update ...
func (m ArticleModel) Update(userID uint, id uint, form forms.ArticleForm) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("article not found")
	}
	err = database.GetDB().Table("public.articles").Model(&Article{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"title": form.Title, "content": form.Content}).Error
	return err
}

// Delete ...
func (m ArticleModel) Delete(userID, id uint) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("Article not found")
	}
	err = database.GetDB().Table("public.articles").Where("id = ?", id).Delete(Article{}).Error

	return err
}
