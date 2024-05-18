package postgresdb

import (
	"database/sql"
	"fmt"
	"intertask/graphqlsh"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	_ "modernc.org/sqlite"
)

// Returns the database generated for tests and model storage.
func DBGenerate() (*sql.DB, *Storage, error) {

	// Open the database from a file.
	db, err := sql.Open("sqlite", "DBfortests.db")
	if err != nil {
		return nil, nil, err
	}

	storage := NewStorage(db)

	// Erase data from database tables.
	res1, err := db.Exec("DELETE FROM users; DELETE FROM posts; DELETE FROM comments;")
	if err != nil {
		fmt.Println(res1 == nil)
		return nil, nil, err
	}

	// Reads commands to generate model data.
	textSQL, err := os.ReadFile("sql/DBfortestsdata.sql")
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	stringSQL := string(textSQL)

	// Creates model data in the database.
	res2, err := db.Exec(stringSQL)
	if err != nil {
		fmt.Println(res2 == nil, err)
		return nil, nil, err
	}

	return db, storage, err
}

// Test function thats updates the status for the ability to comment on the post.
func TestUpdatePost(t *testing.T) {

	db, storage, err := DBGenerate()
	if err != nil {
		t.Errorf("The database is not connected")
	}

	qeryToDB := "SELECT userid, id, text, cancomment FROM posts WHERE id = 1;"

	// Query to the model database.
	rows, err := db.Query(qeryToDB)
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database request failed.")
	}
	defer db.Close()

	var beforeUpdate graphqlsh.Post
	var canCommentNow bool

	for rows.Next() {
		if err = rows.Scan(&beforeUpdate.ID, &beforeUpdate.Text, &beforeUpdate.UserID, canCommentNow); err != nil {
			t.Errorf("The database request failed.")
		}
	}

	beforeUpdate.CanComment = !canCommentNow

	// Expected response after the UpdatePost function works.
	exeptedUpdate := graphqlsh.Post{
		ID:         beforeUpdate.ID,
		UserID:     beforeUpdate.UserID,
		Text:       beforeUpdate.Text,
		CanComment: !canCommentNow,
	}

	afterUpdate, err := storage.UpdatePost(&beforeUpdate)
	if err != nil {
		t.Errorf("The function being tested does not work!")
	}

	// Comparison of structure fields.
	assert.Equal(t, beforeUpdate.ID, afterUpdate.ID, exeptedUpdate.ID)
	assert.Equal(t, beforeUpdate.UserID, afterUpdate.UserID, exeptedUpdate.UserID)
	assert.Equal(t, beforeUpdate.Text, afterUpdate.Text, exeptedUpdate.Text)
	assert.Equal(t, !beforeUpdate.CanComment, afterUpdate.CanComment, exeptedUpdate.CanComment)

}

/*
func TestCreateNewComment(t *testing.T) {
}

func TestCreateNewPost(t *testing.T) {
}
*/
func TestFetchAllPosts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	service := NewStorage(db)

	rows := sqlmock.NewRows([]string{"userid", "id", "text", "cancomment"}).AddRow(1, 1, "text", true)
	mock.ExpectQuery("SELECT userid, id, text, cancomment FROM posts limit 10 offset 10;").WillReturnRows(rows)

	posts, err := service.FetchAllPosts(10, 10)
	if err != nil {
		t.Fail()
	}

	post := posts[0]
	if post.CanComment != true {
		t.Log("мы что-то не так сделали")
		t.Fail()
	}

	fmt.Println(post)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFetchPostByiD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	service := NewStorage(db)

	rows := sqlmock.NewRows([]string{"id", "text", "userid", "cancomment"}).AddRow(2, "text", 1, true).AddRow(1, "text", 1, true)

	qeryToDB := `
		SELECT id, text, userid, cancomment 
		FROM posts WHERE id = 1;`

	a := mock.ExpectQuery(qeryToDB).WillReturnRows(rows)
	fmt.Println(a)

	posts, err := service.FetchPostByiD(1)
	if err != nil {
		fmt.Println("HERE!")
		fmt.Println(err)
		t.Fail()
	}

	post := posts
	if post.CanComment != true {
		t.Log("мы что-то не так сделали")
		t.Fail()
	}

	fmt.Println(post)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

/*
func TestFetchCommentsByPostID(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	service := NewStorage(db)

	rows1 := sqlmock.NewRows([]string{"cancomment"}).AddRow(true).AddRow(true)
	qeryToDB1 := fmt.Sprintf(`
			SELECT cancomment
			FROM posts WHERE id = 1;`)
	a := mock.ExpectQuery(qeryToDB1).WillReturnRows(rows1)

	fmt.Println(a)

	rows2 := sqlmock.NewRows([]string{"id", "userid", "text", "postid", "perentid"}).AddRow(1, 1, "text", 1, 1) //.AddRow(2, 1, "text", 1, 1)
	qeryToDB := fmt.Sprint(`
	   	SELECT id, userid, text, postid, COALESCE(perentid, 0)
	   	FROM comments WHERE postid = 1
	   	ORDER BY perentid, id limit 1 offset 1;`)
	b := mock.ExpectQuery(qeryToDB).WillReturnRows(rows2)

	fmt.Println(b)

	comments, err := service.FetchCommentsByPostID(1, 1, 1)

	if err != nil {
		fmt.Println("HERE!")
		fmt.Println(err)
		t.Fail()
	}

	comment := comments[0]

	fmt.Println(comment)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
*/
