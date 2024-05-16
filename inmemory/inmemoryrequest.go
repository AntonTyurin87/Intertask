package inmemory

import (
	"intertask/graphqlsh"
)

// Structure for storing post and comment data.
type InMemoryType struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	UserID     int    `json:"userid"`
	PerentID   int    `json:"perentid"`
	CanComment bool   `json:"cancomment"`
}

// Memory record counter
var RecordCounter = 1

// Structure for accessing to memory.
type InMemoryStorage struct {
	InMemory []InMemoryType
}

// Creating a new instance of a structure for accessing to memory.
func NewInMemory(InMe []InMemoryType) *InMemoryStorage {
	return &InMemoryStorage{InMemory: InMe}
}

// !!!  Вероятно удалить функцию или пересобрать. Пока бесполезна
func (i *InMemoryStorage) CreateNotification(comment int) ([]graphqlsh.UserSubscription, error) {
	return nil, nil
}

// Creates a record about the user's subscription to a post in memory.
func (i *InMemoryStorage) CreateUserSubscription(newSubscription *graphqlsh.UserSubscription) (*graphqlsh.UserSubscription, error) {
	return nil, nil
}

// Makes a change to the post entry about the ability to comment the post in memory.
func (i *InMemoryStorage) UpdatePost(correctPost *graphqlsh.Post) (*graphqlsh.Post, error) {
	return nil, nil
}

// Creates a record of a new comment to post in memory.
func (i *InMemoryStorage) CreateNewComment(newComment *graphqlsh.Comment) (*graphqlsh.Comment, error) {
	return nil, nil
}

// Creates a record of a new post in memory.
func (i *InMemoryStorage) CreateNewPost(newPost *graphqlsh.Post) (*graphqlsh.Post, error) {

	var err error
	var result graphqlsh.Post

	postData := InMemoryType{
		ID:         RecordCounter,
		Text:       newPost.Text,
		UserID:     newPost.UserID,
		CanComment: newPost.CanComment,
	}

	i.InMemory = append(i.InMemory, postData)

	result = graphqlsh.Post{
		ID:         postData.ID,
		Text:       postData.Text,
		UserID:     postData.UserID,
		CanComment: postData.CanComment,
	}

	RecordCounter++

	return &result, err
}

// Gets all posts from memory.
func (i *InMemoryStorage) FetchAllPosts(limit int, offset int) ([]graphqlsh.Post, error) {
	return nil, nil
}

// Get a post and comments to it by ID from memory.
func (i *InMemoryStorage) FetchPostByiD(id int) (*graphqlsh.Post, error) {
	return nil, nil
}

// Get comments for a specific post from memory.
func (i *InMemoryStorage) FetchCommentsByPostID(id, limit, offset int) ([]graphqlsh.Comment, error) {
	return nil, nil
}
