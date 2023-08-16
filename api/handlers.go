package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/billdev1958/Billdev.git/db"
	"github.com/gorilla/mux"
)

func (s *Server) handlePosts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetPosts(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreatePost(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// handler for post article
func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	req := new(db.CreatePostRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	post, err := db.NewPost(
		req.Title,
		req.Body)
	if err != nil {
		return err
	}
	if err := s.store.CreatePost(post); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, post)
}

// Handler get all posts
func (s *Server) handleGetPosts(w http.ResponseWriter, r *http.Request) error {
	posts, err := s.store.GetPosts()
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, posts)
}

func (s *Server) handleGetPostByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getID(r)
		if err != nil {
			return err
		}
		post, err := s.store.GetPostByID(id)
		if err != nil {
			return err
		}
		return WriteJson(w, http.StatusOK, post)
	}
	if r.Method == "DELETE" {
		return s.handleDeletePost(w, r)

	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *Server) handleDeletePost(w http.ResponseWriter, r *http.Request) error {

	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeletePost(id); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, map[string]int{"deleted": id})
}

// Make json encoder
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// structure for handler functions
type apiFunc func(http.ResponseWriter, *http.Request) error

// error for api func
type apiError struct {
	Error string
}

// Make handler functions for apirest
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

// func for get id params url
func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id givven %s,", idStr)
	}
	return id, nil
}
