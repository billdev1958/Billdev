package db

import "time"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"names"`
	Email string `json:"emails"`
}

type Post struct {
	ID         int
	Title      string
	Body       string
	Created_at time.Time
}

type CreatePostRequest struct {
	Title string
	Body  string
}

func NewPost(title, body string) (*Post, error) {
	return &Post{
		Title:      title,
		Body:       body,
		Created_at: time.Now().UTC(),
	}, nil
}
