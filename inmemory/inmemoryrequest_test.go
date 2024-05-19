package inmemory

import (
	"intertask/graphqlsh"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Filling the model structure with data for subsequent testing of functions.
// Can be copied into each function to produce independent tests.

// Returns a store containing model data.
func inMemoryModelDataCreator() *InMemoryStorage {

	var inMe []InMemoryType

	post1 := InMemoryType{
		ID:         1,
		Text:       "Post Tex1 1 with no comments!",
		UserID:     1,
		CanComment: true,
	}

	post2 := InMemoryType{
		ID:         2,
		Text:       "Post Tex1 2 with 3 comments!",
		UserID:     2,
		CanComment: true,
	}

	// The post that cannot be commented on.
	post3 := InMemoryType{
		ID:         3,
		Text:       "Post Tex1 3 with CanComment = false!",
		UserID:     2,
		CanComment: false,
	}

	comment1 := InMemoryType{
		ID:       4,
		Text:     "New comment 1 to post 1!",
		UserID:   2,
		PostID:   1,
		PerentID: 0,
	}

	comment2 := InMemoryType{
		ID:       5,
		Text:     "New comment 2 to post 1!",
		UserID:   1,
		PostID:   1,
		PerentID: 0,
	}

	// Comment on comment.
	comment3 := InMemoryType{
		ID:       6,
		Text:     "New comment 21 to post 1, to comment 5!",
		UserID:   1,
		PostID:   1,
		PerentID: 5,
	}

	inMe = append(inMe, post1, post2, post3, comment1, comment2, comment3)

	// Model data storage.
	storage := NewInMemory(inMe)

	return storage
}

// Test function thats Returns information about the ability to comment on a post.
func TestReternPostCommentStatus(t *testing.T) {
	// The model data already has a slice with 6 records.
	// The next entry will go under number 7.
	// The next ID start by 10
	storage := inMemoryModelDataCreator()

	//In Data
	id := 3

	expected := false

	//Out Data
	actual, err := storage.ReternPostCommentStatus(id)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	assert.Equal(t, expected, actual)

}

// Test function thats "Makes a change to the post entry about the ability to comment the post in memory."
func TestUpdatePost(t *testing.T) {

	// The model data already has a slice with 6 records.
	// The next entry will go under number 7.
	// The next ID start by 10
	storage := inMemoryModelDataCreator()

	//In Data
	notEmptyPost := graphqlsh.Post{
		ID:         3,
		CanComment: true,
	}

	expectedNotEmpty := graphqlsh.Post{
		ID:         3,
		CanComment: true,
	}

	actualNotEmpty, err := storage.UpdatePost(&notEmptyPost)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	assert.Equal(t, expectedNotEmpty.ID, actualNotEmpty.ID, storage.InMemory[2].ID)
	assert.Equal(t, expectedNotEmpty.CanComment, actualNotEmpty.CanComment, storage.InMemory[2].CanComment)
}

// Test function thats "Creates a record of a new comment to post in memory."
func TestCreateNewComment(t *testing.T) {

	// The model data already has a slice with 6 records.
	// The next entry will go under number 7.
	// The next ID start by 10
	storage := inMemoryModelDataCreator()

	//In Data
	notEmptyComment := graphqlsh.Comment{
		Text:     "New comment!",
		UserID:   2,
		PostID:   1,
		PerentID: 2,
	}

	falseToComment := graphqlsh.Comment{
		Text:   "New comment!",
		UserID: 2,
		PostID: 3,
	}

	//Out Data
	expectedNotEmpty := graphqlsh.Comment{
		ID:       10,
		Text:     "New comment!",
		UserID:   2,
		PostID:   1,
		PerentID: 2,
	}

	expectedFalseToComment := graphqlsh.Comment{}

	actualNotEmpty, err := storage.CreateNewComment(&notEmptyComment)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	actualFalseToComment, err := storage.CreateNewComment(&falseToComment)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	assert.Equal(t, expectedNotEmpty.ID, actualNotEmpty.ID, storage.InMemory[5].ID)
	assert.Equal(t, expectedNotEmpty.UserID, actualNotEmpty.UserID, storage.InMemory[5].UserID)
	assert.Equal(t, expectedNotEmpty.Text, actualNotEmpty.Text, storage.InMemory[5].Text)
	assert.Equal(t, expectedNotEmpty.PostID, actualNotEmpty.PostID, storage.InMemory[5].PostID)
	assert.Equal(t, expectedNotEmpty.PerentID, actualNotEmpty.PerentID, storage.InMemory[5].PerentID)

	assert.Equal(t, expectedFalseToComment.ID, actualFalseToComment.ID)
	assert.Equal(t, expectedFalseToComment.UserID, actualFalseToComment.UserID)
	assert.Equal(t, expectedFalseToComment.Text, actualFalseToComment.Text)
	assert.Equal(t, expectedFalseToComment.PostID, actualFalseToComment.PostID)
	assert.Equal(t, expectedFalseToComment.PerentID, actualFalseToComment.PerentID)

}

// Test function thats "Creates a record of a new post in memory."
func TestCreateNewPost(t *testing.T) {

	// The model data already has a slice with 6 records.
	// The next entry will go under number 7.
	// The next ID start by 10
	storage := inMemoryModelDataCreator()

	//In Data
	notEmptyPost := graphqlsh.Post{
		ID:         11,
		Text:       "New text!",
		UserID:     2,
		CanComment: true,
	}

	//Out Data
	expectedNotEmpty := graphqlsh.Post{
		ID:         11,
		Text:       "New text!",
		UserID:     2,
		CanComment: true,
	}

	actualNotEmpty, err := storage.CreateNewPost(&notEmptyPost)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	// Comparing the fields of structures from the expected structure obtained from the structure function and from the structure instance in memory.
	assert.Equal(t, expectedNotEmpty.ID, actualNotEmpty.ID, storage.InMemory[5].ID)
	assert.Equal(t, expectedNotEmpty.UserID, actualNotEmpty.UserID, storage.InMemory[5].UserID)
	assert.Equal(t, expectedNotEmpty.Text, actualNotEmpty.Text, storage.InMemory[5].Text)
	assert.Equal(t, expectedNotEmpty.CanComment, actualNotEmpty.CanComment, storage.InMemory[5].CanComment)

}

// Test function thats "Gets all posts from memory."
func TestFetchAllPosts(t *testing.T) {

	// The model data already has a slice with 6 records.
	// The next entry will go under number 7.
	// The next ID start by 10
	storage := inMemoryModelDataCreator()

	//In Data
	limit := 3
	offset := 1

	//Out Data
	var exeptedPostsSlice []graphqlsh.Post

	post1 := graphqlsh.Post{
		ID:         storage.InMemory[0].ID,
		UserID:     storage.InMemory[0].UserID,
		Text:       storage.InMemory[0].Text,
		CanComment: storage.InMemory[0].CanComment,
	}

	post2 := graphqlsh.Post{
		ID:         storage.InMemory[1].ID,
		UserID:     storage.InMemory[1].UserID,
		Text:       storage.InMemory[1].Text,
		CanComment: storage.InMemory[1].CanComment,
	}

	post3 := graphqlsh.Post{
		ID:         storage.InMemory[2].ID,
		UserID:     storage.InMemory[2].UserID,
		Text:       storage.InMemory[2].Text,
		CanComment: storage.InMemory[2].CanComment,
	}

	exeptedPostsSlice = append(exeptedPostsSlice, post1, post2, post3)

	actualPostsSlice, err := storage.FetchAllPosts(limit, offset)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	assert.Equal(t, actualPostsSlice[0].ID, exeptedPostsSlice[0].ID, storage.InMemory[0].ID)
	assert.Equal(t, actualPostsSlice[0].UserID, exeptedPostsSlice[0].UserID, storage.InMemory[0].UserID)
	assert.Equal(t, actualPostsSlice[0].Text, exeptedPostsSlice[0].Text, storage.InMemory[0].Text)
	assert.Equal(t, actualPostsSlice[0].CanComment, exeptedPostsSlice[0].CanComment, storage.InMemory[0].CanComment)

	assert.Equal(t, actualPostsSlice[1].ID, exeptedPostsSlice[1].ID, storage.InMemory[1].ID)
	assert.Equal(t, actualPostsSlice[1].UserID, exeptedPostsSlice[1].UserID, storage.InMemory[1].UserID)
	assert.Equal(t, actualPostsSlice[1].Text, exeptedPostsSlice[1].Text, storage.InMemory[1].Text)
	assert.Equal(t, actualPostsSlice[1].CanComment, exeptedPostsSlice[1].CanComment, storage.InMemory[1].CanComment)

	assert.Equal(t, actualPostsSlice[2].ID, exeptedPostsSlice[2].ID, storage.InMemory[2].ID)
	assert.Equal(t, actualPostsSlice[2].UserID, exeptedPostsSlice[2].UserID, storage.InMemory[2].UserID)
	assert.Equal(t, actualPostsSlice[2].Text, exeptedPostsSlice[2].Text, storage.InMemory[2].Text)
	assert.Equal(t, actualPostsSlice[2].CanComment, exeptedPostsSlice[2].CanComment, storage.InMemory[2].CanComment)
}

// Test function thats "Get a post and comments to it by ID from memory."
func TestFetchPostByiD(t *testing.T) {

	// The model data already has a slice with 6 records.
	// The next entry will go under number 7.
	// The next ID start by 10
	storage := inMemoryModelDataCreator()

	//In Data
	id := 1

	//Out Data
	exeptedPost := graphqlsh.Post{
		ID:         storage.InMemory[0].ID,
		UserID:     storage.InMemory[0].UserID,
		Text:       storage.InMemory[0].Text,
		CanComment: storage.InMemory[0].CanComment,
	}

	actualPost, err := storage.FetchPostByiD(id)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	assert.Equal(t, exeptedPost.ID, actualPost.ID, storage.InMemory[0].ID)
	assert.Equal(t, exeptedPost.UserID, actualPost.UserID, storage.InMemory[0].UserID)
	assert.Equal(t, exeptedPost.Text, actualPost.Text, storage.InMemory[0].Text)
	assert.Equal(t, exeptedPost.CanComment, actualPost.CanComment, storage.InMemory[0].CanComment)

}

// Test function thats "Get comments for a specific post from memory."
func TestFetchCommentsByPostID(t *testing.T) {

	// The model data already has a slice with 6 records.
	// The next entry will go under number 7.
	// The next ID start by 10
	storage := inMemoryModelDataCreator()

	//In Data
	id := 1
	limit := 3
	offset := 1

	//Out Data
	var exeptedCommentSlice []graphqlsh.Comment

	comment1 := graphqlsh.Comment{
		ID:       storage.InMemory[3].ID,
		UserID:   storage.InMemory[3].UserID,
		Text:     storage.InMemory[3].Text,
		PostID:   storage.InMemory[3].PostID,
		PerentID: storage.InMemory[3].PerentID,
	}

	comment2 := graphqlsh.Comment{
		ID:       storage.InMemory[4].ID,
		UserID:   storage.InMemory[4].UserID,
		Text:     storage.InMemory[4].Text,
		PostID:   storage.InMemory[4].PostID,
		PerentID: storage.InMemory[4].PerentID,
	}

	comment3 := graphqlsh.Comment{
		ID:       storage.InMemory[5].ID,
		UserID:   storage.InMemory[5].UserID,
		Text:     storage.InMemory[5].Text,
		PostID:   storage.InMemory[5].PostID,
		PerentID: storage.InMemory[5].PerentID,
	}

	exeptedCommentSlice = append(exeptedCommentSlice, comment1, comment2, comment3)

	actualComment, err := storage.FetchCommentsByPostID(id, limit, offset)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	assert.Equal(t, exeptedCommentSlice[0].ID, actualComment[0].ID, storage.InMemory[3].ID)
	assert.Equal(t, exeptedCommentSlice[0].UserID, actualComment[0].UserID, storage.InMemory[3].UserID)
	assert.Equal(t, exeptedCommentSlice[0].Text, actualComment[0].Text, storage.InMemory[3].Text)
	assert.Equal(t, exeptedCommentSlice[0].PostID, actualComment[0].PostID, storage.InMemory[3].PostID)
	assert.Equal(t, exeptedCommentSlice[0].PerentID, actualComment[0].PerentID, storage.InMemory[3].PerentID)

	assert.Equal(t, exeptedCommentSlice[1].ID, actualComment[1].ID, storage.InMemory[4].ID)
	assert.Equal(t, exeptedCommentSlice[1].UserID, actualComment[1].UserID, storage.InMemory[4].UserID)
	assert.Equal(t, exeptedCommentSlice[1].Text, actualComment[1].Text, storage.InMemory[4].Text)
	assert.Equal(t, exeptedCommentSlice[1].PostID, actualComment[1].PostID, storage.InMemory[4].PostID)
	assert.Equal(t, exeptedCommentSlice[1].PerentID, actualComment[1].PerentID, storage.InMemory[4].PerentID)

}
