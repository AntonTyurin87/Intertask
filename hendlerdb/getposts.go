package hendlerdb

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

type Post struct {
	PID          int    `json:"id"`
	Text         string `json:"ptext"`
	PostAuthorID int    `json:"uid"`
	CanComment   bool   `json:"cancomment"`
}

func GetPosts(limit int, offset int, db *sql.DB) ([]Post, error) {
	var posts []Post
	//rows, err := db.Query("SELECT pid, ptext, postAuthorID, cancomment FROM posts limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset))
	rows, err := db.Query("SELECT * FROM posts limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset))

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
