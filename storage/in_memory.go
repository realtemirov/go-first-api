package storage

import (
	"errors"
	"time"

	"github.com/realtemirov/learn/Bootcamp-2022-Udevs/project-6/models"
)

var ArticleStorage map[string]models.Article

func init() {
	ArticleStorage = make(map[string]models.Article)

	now := time.Now()

	ArticleStorage["1"] = models.Article{
		ID: "1",
		Content: models.Content{
			Title: "Title",
			Body:  "Body",
		},
		Author: models.Person{
			Firstname: "John1",
			Lastname:  "Doe1",
		},
		CreatedAt: &now,
	}

	ArticleStorage["2"] = models.Article{
		ID: "2",
		Content: models.Content{
			Title: "Title",
			Body:  "Body",
		},
		Author: models.Person{
			Firstname: "John2",
			Lastname:  "Doe2",
		},
		CreatedAt: &now,
	}
}

func CreateArticle(entity models.Article) error {

	//TODO

	return nil
}

func GetArticleList(search string) (resp []models.Article) {
	for _, v := range ArticleStorage {
		resp = append(resp, v)
	}
	return resp
}

func UpdateArticle(entity models.Article) error {

	val, ok := ArticleStorage[entity.ID]

	if !ok {
		return errors.New("not found")
	}
	now := time.Now()
	val.Author = entity.Author
	val.Content = entity.Content
	val.UpdatedAt = &now

	ArticleStorage[val.ID] = val
	return nil
}
