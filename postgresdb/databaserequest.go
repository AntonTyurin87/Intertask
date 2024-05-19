package postgresdb

import (
	"database/sql"
	"fmt"
	"intertask/graphqlsh"
)

// Structure for accessing the PostgresQL database.
type Storage struct {
	DB *sql.DB
}

// Creating a new instance of a structure for accessing the PostgresQL database.
func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}

// Returns information about the ability to comment on a post.
func (s *Storage) ReternPostCommentStatus(id int) (bool, error) {

	var сanComment bool

	qeryToDB := fmt.Sprintf(`
		SELECT cancomment 
		FROM posts WHERE id = %d;`, id)

	rows, err := s.DB.Query(qeryToDB)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err = rows.Scan(&сanComment); err != nil {
			return false, err
		}
	}
	return сanComment, nil
}

// Makes a change to the post entry in the PostgresQL database about the ability to comment the post.
func (s *Storage) UpdatePost(correctPost *graphqlsh.Post) (*graphqlsh.Post, error) {

	var err error
	var rows *sql.Rows
	var result graphqlsh.Post

	qeryToDB := fmt.Sprintf(`
		UPDATE posts
			SET cancomment =
				CASE
					WHEN 
						(SELECT userid FROM posts WHERE id = %d) = %d 
					THEN %t
				END
			WHERE id = %d 
		RETURNING id, text, userid, cancomment;`,
		correctPost.ID, correctPost.UserID,
		correctPost.CanComment,
		correctPost.ID)

	//Update data about a post in the PostgresQL database.
	rows, err = s.DB.Query(qeryToDB)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err = rows.Scan(&result.ID, &result.Text, &result.UserID, &result.CanComment); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// Creates a record of a new comment to post in the PostgresQL database.
func (s *Storage) CreateNewComment(newComment *graphqlsh.Comment) (*graphqlsh.Comment, error) {

	var err error
	var rows *sql.Rows
	var result graphqlsh.Comment
	//var comentstatus string

	/*
		qeryToDB := fmt.Sprintf(`
				SELECT cancomment
				FROM posts WHERE
				id = %d;`, newComment.PostID)
	*/

	qeryToDB := fmt.Sprintf(`
			SELECT id, text, userid, cancomment 
			FROM posts WHERE id = %d;`,
		newComment.PostID)

	// Requests the status of the post for the possibility of commenting.
	rows, err = s.DB.Query(qeryToDB)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postForComment graphqlsh.Post

	for rows.Next() {
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err = rows.Scan(&postForComment.ID, &postForComment.Text, &postForComment.UserID, &postForComment.CanComment); err != nil {
			return nil, err
		}
	}

	if !postForComment.CanComment {
		return &result, err
	}

	// Trim the string to the maximum character value
	if len(newComment.Text) > 2000 {
		newComment.Text = newComment.Text[:2000]
	}

	var insertToDB string

	// Generates a query string depending on the ID value of the parent object.
	if newComment.PerentID != 0 {
		insertToDB = fmt.Sprintf(`
				INSERT INTO
					comments (userid, text, postid, perentid)
						VALUES (%d,  '%s', %d, %d)
				RETURNING id, userid, text, postid, perentid;`,
			newComment.UserID, newComment.Text,
			newComment.PostID, newComment.PerentID)
	} else {
		insertToDB = fmt.Sprintf(`
				INSERT INTO
					comments (userid, text, postid)
						VALUES (%d,  '%s', %d)
				RETURNING id, userid, text, postid, perentid;`,
			newComment.UserID, newComment.Text,
			newComment.PostID)
	}

	// Writes data about a new comment to post in the PostgresQL database.
	rows, err = s.DB.Query(insertToDB)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err = rows.Scan(&result.ID, &result.UserID, &result.Text, &result.PostID, &result.PerentID); err != nil {
			return nil, err
		}
	}
	return &result, nil
}

// Creates a record of a new post in the PostgresQL database.
func (s *Storage) CreateNewPost(newPost *graphqlsh.Post) (*graphqlsh.Post, error) {

	var err error
	var rows *sql.Rows
	var result graphqlsh.Post

	insertToDB := fmt.Sprintf(`
		INSERT INTO 
			posts (userid, text, cancomment) 
				VALUES (%d, '%s', %t)
		RETURNING id, text, userid, cancomment;`,
		newPost.UserID, newPost.Text, newPost.CanComment)

	// Writes data about a new post in the PostgresQL database.
	rows, err = s.DB.Query(insertToDB)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err = rows.Scan(&result.ID, &result.Text, &result.UserID, &result.CanComment); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// Gets all posts from the PostgresQL database.
func (s *Storage) FetchAllPosts(limit int, offset int) ([]graphqlsh.Post, error) {

	var err error
	var rows *sql.Rows
	var result []graphqlsh.Post

	qeryToDB := fmt.Sprintf(`
		SELECT userid, id, text, cancomment 
			FROM posts limit %d offset %d;`,
		limit, offset)

	// Retrieves data about posts in a PostgresQL database.
	rows, err = s.DB.Query(qeryToDB)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b graphqlsh.Post
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err := rows.Scan(&b.UserID, &b.ID, &b.Text, &b.CanComment); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	return result, nil
}

// Get a post and comments to it by ID from the PostgresQL database.
func (s *Storage) FetchPostByiD(id int) (*graphqlsh.Post, error) {

	var err error
	var rows *sql.Rows
	var result graphqlsh.Post

	// Retrieves data about post and comments to it by ID from the PostgresQL database.
	qeryToDB := fmt.Sprintf(`
		SELECT id, text, userid, cancomment 
		FROM posts WHERE id = %d;`, id)

	rows, err = s.DB.Query(qeryToDB)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err = rows.Scan(&result.ID, &result.Text, &result.UserID, &result.CanComment); err != nil {
			return nil, err
		}
	}
	return &result, nil
}

// Get comments for a specific post from the PostgresQL database.
func (s *Storage) FetchCommentsByPostID(id, limit, offset int) ([]graphqlsh.Comment, error) {

	var err error
	var rows *sql.Rows
	var result []graphqlsh.Comment
	var comentstatus string

	qeryToDB := fmt.Sprintf(`
			SELECT cancomment
			FROM posts WHERE id = %d;`, id)

	// Requests the status of the post for the possibility of commenting.
	err = s.DB.QueryRow(qeryToDB).Scan(&comentstatus)

	// If comments cannot be made, then an empty value is returned.
	if err != nil || comentstatus == "false" {
		return nil, err
	}

	qeryToDB = fmt.Sprintf(`
		SELECT id, userid, text, postid, COALESCE(perentid, 0) 
		FROM comments WHERE postid = %d 
		ORDER BY perentid, id limit %d offset %d;`,
		id, limit, offset)

	rows, err = s.DB.Query(qeryToDB)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var b graphqlsh.Comment
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err := rows.Scan(&b.ID, &b.UserID, &b.Text, &b.PostID, &b.PerentID); err != nil {
			return result, err
		}
		result = append(result, b)
	}

	return result, nil
}
