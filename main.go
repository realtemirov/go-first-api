package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/learn/Bootcamp-2022-Udevs/project-6/handlers"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()

	r.GET("/ping", ping)

	api := r.Group("api")
	v1 := api.Group("v1")

	articleRouter := v1.Group("articles")
	{
		articleRouter.GET("/", handlers.GetArticleList)
		articleRouter.GET("/:id", handlers.GetArticleByID)
		articleRouter.POST("/", handlers.CreateArticle)
		articleRouter.PUT("/", handlers.UpdateArticle)
		articleRouter.DELETE("/:id", handlers.DeleteArticle)

	}

	r.Run()

	r.StaticFS("/static", gin.Dir("static", false))

}
