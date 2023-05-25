package controllers

import (
	"crud/initializers"
	"crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func PostCreate(c *gin.Context) {
	var body Post
	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body Post

	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
