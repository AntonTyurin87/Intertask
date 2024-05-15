package postgresdb

import (
	"database/sql"
	"fmt"
	blogInterface "intertask/cmd/bloginterface"
	"strconv"
	"strings"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}

func (s *Storage) CreateNotification(comment int) ([]blogInterface.UserSubscription, error) {
	var err error
	var rows *sql.Rows
	var result []blogInterface.UserSubscription

	rows, err = s.DB.Query("SELECT * FROM usersubscription WHERE pid = " + strconv.Itoa(comment) + " AND confirmation = true;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b blogInterface.UserSubscription
		if err := rows.Scan(&b.ID, &b.UserID, &b.PostID, &b.Сonfirmation); err != nil {
			return result, err
		}
		result = append(result, b)
	}

	return result, nil
}

func (s *Storage) CreateUserSubscription(newSubscription *blogInterface.UserSubscription) (*blogInterface.UserSubscription, error) {

	//!!!Решить проблему с подпиской много раз на один пост!!!!!

	var err error
	var rows *sql.Rows
	var sqlString string

	if !newSubscription.Сonfirmation {
		rows, err = s.DB.Query(`DELETE FROM usersubscription 
															WHERE 			
																uid = ` + strconv.Itoa(newSubscription.UserID) + ` AND 
																pid = ` + strconv.Itoa(newSubscription.PostID) + ` RETURNING id, uid, pid, (confirmation = false);`)
	} else {

		qeryToDB := fmt.Sprint(`INSERT INTO usersubscription (uid, pid, confirmation) 
		VALUES (` + strconv.Itoa(newSubscription.UserID) + `,
				` + strconv.Itoa(newSubscription.PostID) + `,
				` + strconv.FormatBool(newSubscription.Сonfirmation) + `)
				RETURNING id, uid, pid, confirmation;`)

		sqlString = strings.Join(strings.Fields(qeryToDB), " ")

		rows, err = s.DB.Query(sqlString)

	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result blogInterface.UserSubscription

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.UserID, &result.PostID, &result.Сonfirmation); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (s *Storage) CorrectPost(correctPost *blogInterface.Post) (*blogInterface.Post, error) {
	qeryToDB := fmt.Sprint(`
		UPDATE posts
			SET cancomment =
				CASE
					WHEN (SELECT postauthorid FROM posts 
						WHERE id = ` + strconv.Itoa(correctPost.PID) + `) 
						= ` + strconv.Itoa(correctPost.PostAuthorID) + ` 
					THEN ` + strconv.FormatBool(correctPost.CanComment) + `
				END
			WHERE id = ` + strconv.Itoa(correctPost.PID) + ` 
			
			RETURNING id, text, postauthorid, cancomment;`)

	sqlString := strings.Join(strings.Fields(qeryToDB), " ")

	rows, err := s.DB.Query(sqlString)

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

func (s *Storage) CreateNewComment(newComment *blogInterface.Comment) (*blogInterface.Comment, error) {

	var result blogInterface.Comment
	var err error
	var comentstatus string
	var rows *sql.Rows

	err = s.DB.QueryRow("SELECT cancomment FROM posts WHERE id = " + strconv.Itoa(newComment.PostID) + ";").Scan(&comentstatus)

	if err != nil {
		return nil, err
	}

	if comentstatus == "false" {
		return nil, err
	}

	if newComment.PerentID != 0 {
		rows, err = s.DB.Query("INSERT INTO comments (userid, text, postid, perentid) VALUES (" + strconv.Itoa(newComment.UserID) + ",  '" + fmt.Sprintf(newComment.Text) + "', " + strconv.Itoa(newComment.PostID) + ",  " + strconv.Itoa(newComment.PerentID) + ") RETURNING id, userid, text, postid, perentid;") //.Scan(&result.CID, &result.UserID, &result.Text, &result.PostID, &result.PerentID) //RETURNING id, uid, text, pid, peid;")
	} else {
		rows, err = s.DB.Query("INSERT INTO comments (userid, text, postid) VALUES (" + strconv.Itoa(newComment.UserID) + ",  '" + fmt.Sprintf(newComment.Text) + "', " + strconv.Itoa(newComment.PostID) + ") RETURNING id, userid, text, postid, COALESCE(perentid, 0);") //.Scan(&result.CID, &result.UserID, &result.Text, &result.PostID, &result.PerentID) //RETURNING id, uid, text, pid, peid;")
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		if err = rows.Scan(&result.CID, &result.UserID, &result.Text, &result.PostID, &result.PerentID); err != nil {
			return nil, err
		}
	}
	return &result, nil

}

func (s *Storage) CreateNewPost(newPost *blogInterface.Post) (*blogInterface.Post, error) {

	// Надо объединить запросы в базу или сразу реторнить результат INSERTа
	lastInsertId := 0
	err := s.DB.QueryRow("INSERT INTO posts (postauthorid, text, cancomment) VALUES (" + strconv.Itoa(newPost.PostAuthorID) + ",  '" + fmt.Sprintf(newPost.Text) + "', " + strconv.FormatBool(newPost.CanComment) + ") RETURNING id;").Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	rows, err := s.DB.Query("SELECT id, text, postauthorid, cancomment FROM posts WHERE id = " + strconv.Itoa(lastInsertId) + ";")

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

func (s *Storage) FetchAllPosts(limit int, offset int) ([]blogInterface.Post, error) {
	rows, err := s.DB.Query("SELECT postauthorid, id, text, cancomment FROM posts limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset) + ";")

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

	var result []blogInterface.Comment
	var err error
	var comentstatus string

	err = s.DB.QueryRow("SELECT cancomment FROM posts WHERE id = " + strconv.Itoa(id) + ";").Scan(&comentstatus)

	if err != nil {
		fmt.Println("1")
		return result, err
	}

	//fmt.Println(comentstatus)

	if comentstatus == "false" {
		return result, err
	}

	rows, err := s.DB.Query("SELECT id, userid, text, postid, COALESCE(perentid, 0) FROM comments WHERE postid = " + strconv.Itoa(id) + " ORDER BY perentid, id limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset) + ";") //+ " ORDERED BY perentid, id

	if err != nil {
		fmt.Println("1")
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var b blogInterface.Comment
		if err := rows.Scan(&b.CID, &b.UserID, &b.Text, &b.PostID, &b.PerentID); err != nil {
			return result, err
		}
		result = append(result, b)
	}

	return result, nil
}
