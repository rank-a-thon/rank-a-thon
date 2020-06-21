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
	ID        int64  `gorm:"column:id;primary_key;auto_increment" json:"id"`
	UserID    int64  `gorm:"column:user_id;not null;unique" json:"-"`
	Title     string `gorm:"column:title" json:"title"`
	Content   string `gorm:"column:content" json:"content"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
	User      User   `gorm:"column:user;foreignkey:UserID" json:"user"`
}

// ArticleModel ...
type ArticleModel struct{}

// Create ...
func (m ArticleModel) Create(userID int64, form forms.ArticleForm) (articleID int64, err error) {
	article := Article{UserID: userID, Title: form.Title, Content: form.Content}
	err = database.GetDB().Table("public.article").Create(&article).Pluck("id", articleID).Error
	//err = database.GetDB().QueryRow(
	//	"INSERT INTO public.article(user_id, title, content) VALUES($1, $2, $3) RETURNING id",
	//	userID, form.Title, form.Content).Scan(&articleID)
	return articleID, err
}

// One ...
func (m ArticleModel) One(userID, id int64) (article Article, err error) {
	err = database.GetDB().Table("public.article").
		Where("article.user_id = ? AND article.id = ?", userID, id).
		Joins("left join public.user on article.user_id = user.id").
		Take(&article).Error
	//err = database.GetDB().SelectOne(
	//	&article,
	//	"SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 AND a.id=$2 LIMIT 1",
	//	userID, id)
	return article, err
}

// All ...
func (m ArticleModel) All(userID int64) (articles []Article, err error) {
	err = database.GetDB().Table("public.article").
		Joins("left join public.user on article.user_id = user.id").
		Order("article.id desc").
		Find(&articles).Error
	//_, err = database.GetDB().Select(
	//	&articles,
	//	"SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC",
	//	userID)
	return articles, err
}

// Update ...
func (m ArticleModel) Update(userID int64, id int64, form forms.ArticleForm) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("article not found")
	}
	err = database.GetDB().Table("public.article").Model(&m).
		Where("id = ?", id).
		Updates(map[string]interface{}{"title": form.Title, "content": form.Content}).Error
	//_, err = database.GetDB().Exec("UPDATE public.article SET title=$2, content=$3 WHERE id=$1", id, form.Title, form.Content)

	return err
}

// Delete ...
func (m ArticleModel) Delete(userID, id int64) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("Article not found")
	}
	err = database.GetDB().Table("public.article").Where("id = ?", id).Delete(Article{}).Error
	//_, err = database.GetDB().Exec("DELETE FROM public.article WHERE id=$1", id)

	return err
}
