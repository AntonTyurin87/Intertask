package postgresdb

import (
	"database/sql"
	"strconv"
)

/*
type Comment struct {
	CID      int    `json:"id"`
	UserID   int    `json:"uid"`
	Text     string `json:"text"`
	PostID   int    `json:"pid"`
	PerentID int    `json:"peid"`
}

// Structure for working with a database
type Storage struct {
	DB *sql.DB
}

// Constructor for new objects Storage struct
func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}
*/

type Post struct {
	PID          int    `json:"id"`
	Text         string `json:"text"`
	PostAuthorID int    `json:"postauthorid"`
	CanComment   bool   `json:"cancomment"`
}

type Comment struct {
	CID      int    `json:"id"`
	UserID   int    `json:"uid"`
	Text     string `json:"text"`
	PostID   int    `json:"pid"`
	PerentID int    `json:"peid"`
}

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}

func (db *Storage) FetchPostByiD(id int) (*Post, error) {

	var err error

	rows, err := db.DB.Query("SELECT id, text, postauthorid, cancomment FROM posts WHERE id = " + strconv.Itoa(id) + ";")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Post

	for rows.Next() {
		if err = rows.Scan(&result.PID, &result.Text, &result.PostAuthorID, &result.CanComment); err != nil {
			return nil, err
		}
	}
	return &result, nil
}

func (db *Storage) FetchCommentsByPostID(id int) ([]Comment, error) {

	rows, err := db.DB.Query("SELECT id, userid, text, postid, COALESCE(perentid, 0) FROM comments WHERE postid = " + strconv.Itoa(id) + " ORDER BY perentid, id;") //+ " ORDERED BY perentid, id

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Comment

	for rows.Next() {
		var b Comment
		if err := rows.Scan(&b.CID, &b.UserID, &b.Text, &b.PostID, &b.PerentID); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	return result, nil
}
