package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Methods store
type Storage interface {
	CreatePost(*Post) error
	GetPosts() ([]*Post, error)
	GetPostByID(int) (*Post, error)
	DeletePost(int) error
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStore{
		db: db,
	}, nil

}

// Create one post
func (s *PostgreStore) CreatePost(post *Post) error {
	query := `INSERT INTO posts
	(title, body, created_at)
	values ($1, $2, $3)`

	_, err := s.db.Query(
		query,
		post.Title,
		post.Body,
		post.Created_at,
	)
	if err != nil {
		return err
	}
	return nil
}

// Method get post by id
func (s *PostgreStore) GetPostByID(id int) (*Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoPosts(rows)
	}
	return nil, fmt.Errorf("post %d notf found", id)

}

// Method get users all
func (s *PostgreStore) GetPosts() ([]*Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	posts := []*Post{}
	for rows.Next() {
		post, err := scanIntoPosts(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostgreStore) DeletePost(id int) error {
	_, err := s.db.Query("DELETE FROM posts WHERE id = $1", id)
	return err

}

func scanIntoPosts(rows *sql.Rows) (*Post, error) {
	post := new(Post)
	err := rows.Scan(
		&post.ID,
		&post.Title,
		&post.Body,
		&post.Created_at)
	return post, err

}
