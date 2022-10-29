package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/learn/Bootcamp-2022-Udevs/project-6/models"
)

var ArticleStorage map[int]models.Article

func init() {
	ArticleStorage = make(map[int]models.Article)

	now := time.Now()

	ArticleStorage[1] = models.Article{
		ID: 1,
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

	ArticleStorage[2] = models.Article{
		ID: 2,
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

func GetArticleList(c *gin.Context) {
	var list []models.Article
	list = make([]models.Article, 0)

	for _, v := range ArticleStorage {
		list = append(list, v)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleList",
		"data":    list,
	})
}

func GetArticleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := ArticleStorage[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"Not Foudf": id})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleListByid",
		"data":    ArticleStorage[id],
	})
}

func CreateArticle(c *gin.Context) {
	var json interface{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "PostArticle",
		"json_body": json,
	})
}

func UpdateArticle(c *gin.Context) {
	var data models.Article
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := ArticleStorage[data.ID]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"Not Foudf": data.ID})
		return
	}

	ArticleStorage[data.ID] = data

	c.JSON(http.StatusOK, gin.H{
		"message":   "PutArticle",
		"json_body": data,
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETEArticle",
		"id":      id,
	})
}
