package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get data off req body
	var body struct{
		Title string
		Body string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}
 	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	//get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//respond with the posts
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	//get id off the url
	id := c.Param("id")

	//get the data off req body
	var body struct{
		Title string
		Body string
	}

	c.Bind(&body)

	//find the post by we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//respond with the post
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostDelete(c *gin.Context) {
	//get id off the url
	id := c.Param("id")

	//delete the post
	initializers.DB.Delete(&models.Post{}, id)
	

	//respond with the post
	c.Status(200)
}