package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	repo *PostRepository
}

func NewPostHandler(repo *PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (h *PostHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/posts", h.createBlogPost)
	r.PUT("/posts/:id", h.updateBlogPost)
	r.DELETE("/posts/:id", h.deleteBlogPost)
	r.GET("/posts", h.getBlogPost)
	r.GET("/posts/:id", h.getBlogPostById)
}

func (h *PostHandler) createBlogPost(c *gin.Context) {
	var payload PostDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	var post Post
	if err := h.repo.CreatePost(payload, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) updateBlogPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}
	var payload PostDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	var post Post
	if err := h.repo.UpdatePost(id, payload, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating post"})
		return
	}

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) deleteBlogPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	ok, err := h.repo.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting post"})
		return
	}

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *PostHandler) getBlogPost(c *gin.Context) {
	queryTerm := c.Query("term")
	posts := []Post{}
	if err := h.repo.GetPosts(queryTerm, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) getBlogPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	var post Post
	if err := h.repo.GetPostById(id, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error searching post"})
		return
	}

	if post.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, post)
}
