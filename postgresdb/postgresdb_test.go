package postgresdb

import (
	"database/sql"
	"fmt"
	"intertask/graphqlsh"
	"os"
	"testing"

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

	// Erase database tables.
	res1, err := db.Exec("DROP TABLE IF EXISTS users; DROP TABLE IF EXISTS posts; DROP TABLE IF EXISTS comments;")
	if err != nil {
		fmt.Println(res1 == nil, err)
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
	res3, err := db.Exec(stringSQL)
	if err != nil {
		fmt.Println(res3 == nil, err)
		return nil, nil, err
	}

	return db, storage, err
}

// Test function thats updates the status for the ability to comment on the post.
func TestUpdatePost(t *testing.T) {

	db, storage, err := DBGenerate()
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database is not connected")
	}

	// In Data
	qeryToDB := "SELECT id, text, userid, cancomment FROM posts WHERE id = 1;"

	// Query to the model database.
	rows, err := db.Query(qeryToDB)
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database request failed.")
	}

	var beforeUpdate graphqlsh.Post //As In Data
	var canCommentNow bool

	for rows.Next() {
		if err = rows.Scan(&beforeUpdate.ID, &beforeUpdate.Text, &beforeUpdate.UserID, &canCommentNow); err != nil {
			fmt.Print(err)
			t.Errorf("The database request failed.")
		}
	}
	defer db.Close()

	beforeUpdate.CanComment = !canCommentNow

	// Expected response after the UpdatePost function works.
	exeptedUpdate := graphqlsh.Post{
		ID:         beforeUpdate.ID,
		UserID:     beforeUpdate.UserID,
		Text:       beforeUpdate.Text,
		CanComment: !canCommentNow,
	}

	//Out Data
	afterUpdate, err := storage.UpdatePost(&beforeUpdate)
	if err != nil {
		fmt.Println(err)
		t.Errorf("The function being tested does not work!")
	}

	// Comparison of structure fields.
	assert.Equal(t, beforeUpdate.ID, afterUpdate.ID, exeptedUpdate.ID)
	assert.Equal(t, beforeUpdate.UserID, afterUpdate.UserID, exeptedUpdate.UserID)
	assert.Equal(t, beforeUpdate.Text, afterUpdate.Text, exeptedUpdate.Text)
	assert.Equal(t, beforeUpdate.CanComment, afterUpdate.CanComment, exeptedUpdate.CanComment)

}

// Test function thats creates a record of a new comment to post in the PostgresQL database.
func TestCreateNewComment(t *testing.T) {

	db, storage, err := DBGenerate()
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database is not connected.")
	}

	// In Data
	commentForTruePost := graphqlsh.Comment{
		Text:     "New unique comment!",
		UserID:   2,
		PostID:   3,
		PerentID: 1,
	}

	// Out Data
	actualCommentForTruePost, err := storage.CreateNewComment(&commentForTruePost)
	if err != nil {
		fmt.Println(err)
		t.Errorf("New comment creation failed.")
	}

	// in Storage Data
	qeryToDB := fmt.Sprintf(`
		SELECT id, text, userid, postid 
		FROM comments WHERE text ='%s' AND userid = %d AND postid = %d;`,
		commentForTruePost.Text, commentForTruePost.UserID, commentForTruePost.PostID)

	rows, err := db.Query(qeryToDB)
	if err != nil {
		fmt.Println(err)
		t.Errorf("The request for a new comment from the model database failed.")
	}
	defer rows.Close()

	var inStorageCommentForTruePost graphqlsh.Comment

	for rows.Next() {
		if err = rows.Scan(&inStorageCommentForTruePost.ID, &inStorageCommentForTruePost.Text, &inStorageCommentForTruePost.UserID, &inStorageCommentForTruePost.PostID); err != nil {
			fmt.Println(err)
			t.Errorf("Writing new comment query results from the model database failed.")
		}
	}

	assert.Equal(t, actualCommentForTruePost.ID, inStorageCommentForTruePost.ID)
	assert.Equal(t, commentForTruePost.UserID, actualCommentForTruePost.UserID, inStorageCommentForTruePost.UserID)
	assert.Equal(t, commentForTruePost.Text, actualCommentForTruePost.Text, inStorageCommentForTruePost.Text)
	assert.Equal(t, commentForTruePost.PostID, actualCommentForTruePost.PostID, inStorageCommentForTruePost.PostID)
}

// Test function thats creates a record of a new post in the PostgresQL database.
func TestCreateNewPost(t *testing.T) {

	db, storage, err := DBGenerate()
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database is not connected.")
	}

	// In Data
	notEmptyPost := graphqlsh.Post{
		Text:       "New unique text!",
		UserID:     2,
		CanComment: true,
	}

	// Out Data
	actualNotEmpty, err := storage.CreateNewPost(&notEmptyPost)
	if err != nil {
		fmt.Println(err)
		t.Errorf("New post creation failed.")
	}

	// in Storage Data
	qeryToDB := fmt.Sprintf(`
		SELECT id, text, userid, cancomment 
		FROM posts WHERE text ='%s' AND userid = %d AND cancomment = %t;`,
		notEmptyPost.Text, notEmptyPost.UserID, notEmptyPost.CanComment)

	rows, err := db.Query(qeryToDB)
	if err != nil {
		fmt.Println(err)
		t.Errorf("The request for a new post from the model database failed.")
	}
	defer rows.Close()

	var inStorageNotEmpty graphqlsh.Post
	for rows.Next() {
		if err = rows.Scan(&inStorageNotEmpty.ID, &inStorageNotEmpty.Text, &inStorageNotEmpty.UserID, &inStorageNotEmpty.CanComment); err != nil {
			fmt.Println(err)
			t.Errorf("Writing new post query results from the model database failed.")
		}
	}

	assert.Equal(t, actualNotEmpty.ID, inStorageNotEmpty.ID)
	assert.Equal(t, notEmptyPost.UserID, actualNotEmpty.UserID, inStorageNotEmpty.UserID)
	assert.Equal(t, notEmptyPost.Text, actualNotEmpty.Text, inStorageNotEmpty.Text)
	assert.Equal(t, notEmptyPost.CanComment, actualNotEmpty.CanComment, inStorageNotEmpty.CanComment)

}

// Test function thats gets all posts from the PostgresQL database.
func TestFetchAllPosts(t *testing.T) {

	db, storage, err := DBGenerate()
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database is not connected.")
	}

	// In Data
	limit := 3
	offset := 1

	// Expected result
	var expectedResult []graphqlsh.Post

	qeryToDB := fmt.Sprintf(`
		SELECT userid, id, text, cancomment 
			FROM posts limit %d offset %d;`,
		limit, offset)

	rows, err := db.Query(qeryToDB)

	if err != nil {
		fmt.Print(err)
		t.Errorf("The database did not process the request.")
	}
	defer rows.Close()

	for rows.Next() {
		var b graphqlsh.Post
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err := rows.Scan(&b.UserID, &b.ID, &b.Text, &b.CanComment); err != nil {
			fmt.Print(err)
			t.Errorf("The query result could not be written.")
		}
		expectedResult = append(expectedResult, b)
	}

	//Out Data
	actualResult, err := storage.FetchAllPosts(limit, offset)

	if err != nil {
		fmt.Println(err)
		t.Errorf("New posts slice creation failed.")
	}

	assert.Equal(t, expectedResult[0].ID, actualResult[0].ID)
	assert.Equal(t, expectedResult[0].UserID, actualResult[0].UserID)
	assert.Equal(t, expectedResult[0].Text, actualResult[0].Text)
	assert.Equal(t, expectedResult[0].CanComment, actualResult[0].CanComment)

	assert.Equal(t, expectedResult[1].ID, actualResult[1].ID)
	assert.Equal(t, expectedResult[1].UserID, actualResult[1].UserID)
	assert.Equal(t, expectedResult[1].Text, actualResult[1].Text)
	assert.Equal(t, expectedResult[1].CanComment, actualResult[1].CanComment)

	assert.Equal(t, expectedResult[2].ID, actualResult[2].ID)
	assert.Equal(t, expectedResult[2].UserID, actualResult[2].UserID)
	assert.Equal(t, expectedResult[2].Text, actualResult[2].Text)
	assert.Equal(t, expectedResult[2].CanComment, actualResult[2].CanComment)
}

// Test function thats get a post and comments to it by ID from the PostgresQL database.
func TestFetchPostByiD(t *testing.T) {

	db, storage, err := DBGenerate()
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database is not connected.")
	}

	// In Data
	id := 3

	// Expected result
	var expectedResult graphqlsh.Post

	qeryToDB := fmt.Sprintf(`
			SELECT userid, id, text, cancomment 
				FROM posts WHERE id = %d;`,
		id)

	rows, err := db.Query(qeryToDB)

	if err != nil {
		fmt.Print(err)
		t.Errorf("The database did not process the request.")
	}
	defer rows.Close()

	for rows.Next() {
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err := rows.Scan(&expectedResult.UserID, &expectedResult.ID, &expectedResult.Text, &expectedResult.CanComment); err != nil {
			fmt.Print(err)
			t.Errorf("The query result could not be written.")
		}
	}

	//Out Data
	actualResult, err := storage.FetchPostByiD(id)

	if err != nil {
		fmt.Println(err)
		t.Errorf("New posts slice creation failed.")
	}

	assert.Equal(t, expectedResult.ID, actualResult.ID)
	assert.Equal(t, expectedResult.UserID, actualResult.UserID)
	assert.Equal(t, expectedResult.Text, actualResult.Text)
	assert.Equal(t, expectedResult.CanComment, actualResult.CanComment)

}

// Test function thats get comments for a specific post from the PostgresQL database.
func TestFetchCommentsByPostID(t *testing.T) {

	db, storage, err := DBGenerate()
	if err != nil {
		fmt.Println(err)
		t.Errorf("The database is not connected.")
	}

	// In Data
	id := 1
	limit := 3
	offset := 1

	// Expected result
	var expectedResult []graphqlsh.Comment

	qeryToDB := fmt.Sprintf(`
		SELECT userid, id, text, postid 
			FROM comments WHERE postid = %d limit %d offset %d;`,
		id, limit, offset)

	rows, err := db.Query(qeryToDB)

	if err != nil {
		fmt.Print(err)
		t.Errorf("The database did not process the request.")
	}
	defer rows.Close()

	for rows.Next() {
		var b graphqlsh.Comment
		// Writes the values ​​obtained from the PostgresQL database to the result.
		if err := rows.Scan(&b.UserID, &b.ID, &b.Text, &b.PostID); err != nil {
			fmt.Print(err)
			t.Errorf("The query result could not be written.")
		}
		expectedResult = append(expectedResult, b)
	}

	//Out Data
	actualResult, err := storage.FetchCommentsByPostID(id, limit, offset)

	if err != nil {
		fmt.Println(err)
		t.Errorf("New posts slice creation failed.")
	}

	assert.Equal(t, expectedResult[0].ID, actualResult[0].ID)
	assert.Equal(t, expectedResult[0].UserID, actualResult[0].UserID)
	assert.Equal(t, expectedResult[0].Text, actualResult[0].Text)
	assert.Equal(t, expectedResult[0].PostID, actualResult[0].PostID)

	assert.Equal(t, expectedResult[1].ID, actualResult[1].ID)
	assert.Equal(t, expectedResult[1].UserID, actualResult[1].UserID)
	assert.Equal(t, expectedResult[1].Text, actualResult[1].Text)
	assert.Equal(t, expectedResult[1].PostID, actualResult[1].PostID)

	assert.Equal(t, expectedResult[2].ID, actualResult[2].ID)
	assert.Equal(t, expectedResult[2].UserID, actualResult[2].UserID)
	assert.Equal(t, expectedResult[2].Text, actualResult[2].Text)
	assert.Equal(t, expectedResult[2].PostID, actualResult[2].PostID)

}
