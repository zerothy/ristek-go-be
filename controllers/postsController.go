package controllers

import (
	"github.com/gin-gonic/gin"
	"server/initializers"
	"server/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get ID off URl
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the URL
	id := c.Param("id")

	// Get the data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respont with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off URL
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with it
	c.Status(200)
}
