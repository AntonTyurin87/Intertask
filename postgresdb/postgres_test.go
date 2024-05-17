package postgresdb

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUpdatePost(t *testing.T) {
	/*
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		service := NewStorage(db)

		rows := sqlmock.NewRows([]string{"id", "text", "userid", "cancomment"}).AddRow(1, "text", 1, true)

			qeryToDB := `
				UPDATE posts
					SET cancomment =
						CASE
							WHEN
								(SELECT userid FROM posts WHERE id = 1) = 1
							THEN true
						END
					WHERE id = 1;`

		mokQuery := "SELECT id, text, userid, cancomment FROM posts WHERE id = 1;"

		mock.ExpectQuery(mokQuery).WillReturnRows(rows)

		var correctPost = graphqlsh.Post{
			ID:         1,
			UserID:     1,
			CanComment: true,
		}

		ww := &correctPost

		posts, err := service.UpdatePost(&correctPost)
		if err != nil {
			t.Fail()
		}

		post := posts
		if post.CanComment != true {
			t.Log("мы что-то не так сделали")
			t.Fail()
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	*/
}

func TestCreateNewComment(t *testing.T) {
}

func TestCreateNewPost(t *testing.T) {
}

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
