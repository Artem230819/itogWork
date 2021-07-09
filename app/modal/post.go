package modal

import "github.com/Artem230819/itogWork/server"

type Post struct {
	Id   int
	Text string
	Date string `db:"date_completion"`
}

func NewPost(text string) *Post {
	return &Post{Text: text}
}

func GetAllPosts() (posts []Post, err error) {
	query := `SELECT * FROM posts`
	err = server.Db.Select(&posts, query)
	return
}

func (p *Post) Add() (err error) {
	query := `INSERT INTO posts (text) VALUES (?)`
	_, err = server.Db.Exec(query, p.Text)
	return
}
