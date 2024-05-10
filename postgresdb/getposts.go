package postgresdb

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

// Structure for working by posts
type Post struct {
	PID          int    `json:"id"`
	Text         string `json:"ptext"`
	PostAuthorID int    `json:"uid"`
	CanComment   bool   `json:"cancomment"`
}

// Structure for working with a database
type Storage struct {
	db *sql.DB
}

// Constructor for new objects Storage struct
func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (db *Storage) GetPosts(limit int, offset int) ([]Post, error) {
	var posts []Post
	rows, err := db.db.Query("SELECT id, text, postauthorid, cancomment FROM posts limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset))
	//rows, err := db.db.Query("SELECT * FROM posts limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Post
		if err := rows.Scan(&b.PID, &b.Text, &b.PostAuthorID, &b.CanComment); err != nil {
			return nil, err
		}
		posts = append(posts, b)
	}
	return posts, nil
}
