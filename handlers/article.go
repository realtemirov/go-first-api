package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/learn/Bootcamp-2022-Udevs/project-6/models"
	"github.com/realtemirov/learn/Bootcamp-2022-Udevs/project-6/storage"
)

func GetArticleList(c *gin.Context) {

	search := ""

	resp := storage.GetArticleList(search)

	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleList",
		"data":    resp,
	})
}

func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	if _, ok := storage.ArticleStorage[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"Not Foudf": id})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetArticleListByid",
		"data":    storage.ArticleStorage[id],
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

	err := storage.UpdateArticle(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article has been updated",
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETEArticle",
		"id":      id,
	})
}
