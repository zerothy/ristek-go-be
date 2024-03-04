package controllers

import (
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Amount   int
		Types    string
		Category string
		Dates    string
		Note     string
	}

	c.Bind(&body)

	// Create post
	post := models.Post{Amount: body.Amount, Types: body.Types, Category: body.Category, Dates: body.Dates, Note: body.Note}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})

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
		Amount   int
		Types    string
		Category string
		Dates    string
		Note     string
	}

	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Amount:   body.Amount,
		Types:    body.Types,
		Category: body.Category,
		Dates:    body.Dates,
		Note:     body.Note,
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
	result := initializers.DB.Delete(&models.Post{}, id)

	// Check if it was found
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})

		return
	}

	// Check for errors
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	// Respond with it
	c.JSON(200, gin.H{
		"message": "Deleted post with id: " + id,
	})
}

// This is the function to sort by dates
func PostsIndexByType(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Create a map to store posts by date
	postsByDate := make(map[string][]models.Post)

	// Group posts by date
	for _, post := range posts {
		postsByDate[post.Dates] = append(postsByDate[post.Dates], post)
	}

	// Respond with the posts grouped by date
	c.JSON(200, gin.H{
		"postsByDate": postsByDate,
	})
}
