package postgresdb

import (
	"database/sql"
	"fmt"
	blogInterface "intertask/cmd/bloginterface"
	"strconv"
	//"intertask/cmd/bloginterface"
)

// Blog Interface
type Blog interface {
	FetchAllPosts(limit, offset int) ([]blogInterface.Post, error)
	FetchPostByiD(id int) (*blogInterface.Post, error)
	FetchCommentsByPostID(id, limit, offset int) ([]blogInterface.Comment, error)
}

// Blog interface function to retrieve all posts
func AllPosts(b Blog, limit, offset int) ([]blogInterface.Post, error) {
	return b.FetchAllPosts(limit, offset) // From DB
	// From memory
}

// Blog interface function to retrieve post by ID
func PostById(b Blog, id int) (*blogInterface.Post, error) {
	return b.FetchPostByiD(id) // From DB
	// From memory
}

// Blog interface function to retrieve comments by PostID
func CommentsByPostID(b Blog, id, limit, offset int) ([]blogInterface.Comment, error) {
	return b.FetchCommentsByPostID(id, limit, offset)
	// From memory
}

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}

func (s *Storage) FetchAllPosts(limit int, offset int) ([]blogInterface.Post, error) {
	rows, err := s.DB.Query("SELECT postauthorid, id, text, cancomment FROM posts limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset) + ";")
	//rows, err := s.DB.Query("SELECT * FROM posts limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset) + ";")

	//fmt.Println("1")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []blogInterface.Post

	for rows.Next() {
		var b blogInterface.Post
		if err := rows.Scan(&b.PostAuthorID, &b.PID, &b.Text, &b.CanComment); err != nil {
			return nil, err
		}
		posts = append(posts, b)
	}
	//fmt.Println("2")
	fmt.Println(posts)

	return posts, nil
}

func (s *Storage) FetchPostByiD(id int) (*blogInterface.Post, error) {

	var err error

	rows, err := s.DB.Query("SELECT id, text, postauthorid, cancomment FROM posts WHERE id = " + strconv.Itoa(id) + ";")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result blogInterface.Post

	for rows.Next() {
		if err = rows.Scan(&result.PID, &result.Text, &result.PostAuthorID, &result.CanComment); err != nil {
			return nil, err
		}
	}
	return &result, nil
}

func (s *Storage) FetchCommentsByPostID(id, limit, offset int) ([]blogInterface.Comment, error) {

	rows, err := s.DB.Query("SELECT id, userid, text, postid, COALESCE(perentid, 0) FROM comments WHERE postid = " + strconv.Itoa(id) + " ORDER BY perentid, id limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset) + ";") //+ " ORDERED BY perentid, id

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []blogInterface.Comment

	for rows.Next() {
		var b blogInterface.Comment
		if err := rows.Scan(&b.CID, &b.UserID, &b.Text, &b.PostID, &b.PerentID); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	return result, nil
}
