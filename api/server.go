package api

import (
	"log"
	"net/http"

	"github.com/billdev1958/Billdev.git/db"
	"github.com/gorilla/mux"
)

type MyHandler struct{}

type Server struct {
	store      db.Storage
	listenAddr string
}

func NewServer(listenAddr string, store db.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", makeHTTPHandleFunc(s.handlePosts))
	router.HandleFunc("/posts/{id}", makeHTTPHandleFunc(s.handleGetPostByID))
	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
