package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/umdalecs/blogging-platform-api/utils"
)

type PostHandler struct {
	Repo *PostRepository
}

func NewPostHandler(repo *PostRepository) *PostHandler {
	return &PostHandler{Repo: repo}
}

func (h *PostHandler) RegisterRoutes(m *http.ServeMux) {
	m.HandleFunc("POST /posts/", h.createBlogPost)
	m.HandleFunc("POST /posts", h.createBlogPost) // fallback for no final slash
	m.HandleFunc("PUT /posts/{id}", h.updateBlogPost)
	m.HandleFunc("DELETE /posts/{id}", h.deleteBlogPost)
	m.HandleFunc("GET /posts/", h.getBlogPost)
	m.HandleFunc("GET /posts", h.getBlogPost) // fallback for no final slash
	m.HandleFunc("GET /posts/{id}", h.getBlogPostById)
}

func (h *PostHandler) createBlogPost(w http.ResponseWriter, r *http.Request) {
	var payload PostDto
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteJsonErr(w, http.StatusBadRequest, fmt.Errorf("invalid body content"))
		return
	}

	post, err := h.Repo.CreatePost(payload)
	if err != nil {
		utils.WriteJsonErr(w, http.StatusInternalServerError, fmt.Errorf("error storing post"))
		return
	}

	utils.WriteJson(w, http.StatusCreated, post)
}

func (h *PostHandler) updateBlogPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJsonErr(w, http.StatusBadRequest, fmt.Errorf("id must be an integer"))
		return
	}
	var payload PostDto
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteJsonErr(w, http.StatusBadRequest, fmt.Errorf("invalid body content"))
		return
	}

	post, err := h.Repo.UpdatePost(id, payload)
	if err != nil {
		utils.WriteJsonErr(w, http.StatusInternalServerError, fmt.Errorf("error updating post"))
		return
	}

	if post.ID == 0 {
		utils.WriteJsonErr(w, http.StatusNotFound, fmt.Errorf("post not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, post)
}

func (h *PostHandler) deleteBlogPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJsonErr(w, http.StatusBadRequest, fmt.Errorf("id must be an integer"))
		return
	}

	ok, err := h.Repo.DeletePost(id)
	if err != nil {
		utils.WriteJsonErr(w, http.StatusInternalServerError, fmt.Errorf("error deleting post"))
		return
	}

	if !ok {
		utils.WriteJsonErr(w, http.StatusNotFound, fmt.Errorf("post not found"))
		return
	}

	utils.WriteEmpty(w, http.StatusNoContent)
}

func (h *PostHandler) getBlogPost(w http.ResponseWriter, r *http.Request) {
	queryTerm := r.URL.Query().Get("term")
	posts, err := h.Repo.GetPosts(queryTerm)
	if err != nil {
		utils.WriteJsonErr(w, http.StatusInternalServerError, fmt.Errorf("error searching posts"))
		return
	}
	utils.WriteJson(w, http.StatusOK, posts)
}

func (h *PostHandler) getBlogPostById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJsonErr(w, http.StatusBadRequest, fmt.Errorf("id must be an integer"))
		return
	}

	post, err := h.Repo.GetPostById(id)
	if err != nil {
		utils.WriteJsonErr(w, http.StatusInternalServerError, fmt.Errorf("error searching post"))
		return
	}

	if post.ID == 0 {
		utils.WriteEmpty(w, http.StatusNotFound)
		return
	}

	utils.WriteJson(w, http.StatusOK, post)
}
