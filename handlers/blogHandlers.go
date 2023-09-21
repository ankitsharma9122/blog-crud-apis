package handlers

import (
	"context"
	"net/http"
	"project-ankit/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	collection = models.GetBlogPostCollection()
}

func CreateBlogPost(c *gin.Context) {
	var post models.BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.CreatedAt = time.Now()
	result, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Blog post created successfully", "post": result})
}

func GetBlogPosts(c *gin.Context) {
	var posts []models.BlogPost
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var post models.BlogPost
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog posts retrieved successfully", "posts": posts})
}

func GetBlogPost(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var post models.BlogPost
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog post retrieved successfully", "post": post})
}

func UpdateBlogPost(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var post models.BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPost := bson.M{
		"$set": bson.M{
			"title":   post.Title,
			"content": post.Content,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, updatedPost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog post updated successfully"})
}

func DeleteBlogPost(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully"})
}
