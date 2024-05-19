package inmemory

import (
	"intertask/graphqlsh"
	"sort"
)

// Structure for storing post and comment data.
type InMemoryType struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	UserID     int    `json:"userid"`
	PostID     int    `json:"postid"`
	PerentID   int    `json:"perentid"`
	CanComment bool   `json:"cancomment"`
}

// Memory record counter
// Used to form ID.
// ID numbers 1 to 9 are used for testing.
var RecordCounter = 10

// Structure for accessing to memory.
type InMemoryStorage struct {
	InMemory []InMemoryType
}

// Creating a new instance of a structure for accessing to memory.
func NewInMemory(InMe []InMemoryType) *InMemoryStorage {
	return &InMemoryStorage{InMemory: InMe}
}

// Returns information about the ability to comment on a post.
func (i *InMemoryStorage) ReternPostCommentStatus(id int) (bool, error) {

	for _, value := range i.InMemory {
		if value.ID == id {
			return value.CanComment, nil
		}
	}
	return false, nil
}

// Makes a change to the post entry about the ability to comment the post in memory.
func (i *InMemoryStorage) UpdatePost(correctPost *graphqlsh.Post) (*graphqlsh.Post, error) {

	var result graphqlsh.Post

	for j, value := range i.InMemory {
		if value.ID == correctPost.ID {
			if value.CanComment != correctPost.CanComment {

				result = graphqlsh.Post{
					ID:         value.ID,
					Text:       value.Text,
					UserID:     value.ID,
					CanComment: correctPost.CanComment,
				}

				postData := InMemoryType{
					ID:         value.ID,
					Text:       value.Text,
					UserID:     value.ID,
					CanComment: correctPost.CanComment,
				}

				// Removes an outdated record from a record slice and replaces it with a new one.
				i.InMemory = append(i.InMemory[:j], i.InMemory[j+1:]...)
				i.InMemory = append(i.InMemory, postData)
			}

			return &result, nil
		}
	}
	return &result, nil // If a post to update on is not found, then return an empty value.
}

// Creates a record of a new comment to post in memory.
func (i *InMemoryStorage) CreateNewComment(newComment *graphqlsh.Comment) (*graphqlsh.Comment, error) {

	var result graphqlsh.Comment

	for _, value := range i.InMemory {
		if value.ID == newComment.PostID {
			if !value.CanComment {
				return &result, nil
			} else {

				var trimText string
				// If the comment is longer than 2000 characters,
				// then only the first 2000 characters are left and the rest is discarded.
				if len(newComment.Text) > 2000 {
					trimText = newComment.Text[0:2000]
				} else {
					trimText = newComment.Text
				}

				commentData := InMemoryType{
					ID:       RecordCounter,
					PostID:   newComment.PostID,
					PerentID: newComment.PerentID,
					UserID:   newComment.UserID,
					Text:     trimText,
				}

				i.InMemory = append(i.InMemory, commentData)

				result = graphqlsh.Comment{
					ID:       commentData.ID,
					UserID:   commentData.UserID,
					Text:     commentData.Text,
					PostID:   commentData.PostID,
					PerentID: commentData.PerentID,
				}

				RecordCounter++

				return &result, nil
			}
		}
	}
	return &result, nil // If a post to comment on is not found, then return an empty value.
}

// Creates a record of a new post in memory.
func (i *InMemoryStorage) CreateNewPost(newPost *graphqlsh.Post) (*graphqlsh.Post, error) {

	var result graphqlsh.Post

	postData := InMemoryType{
		ID:         RecordCounter,
		Text:       newPost.Text,
		UserID:     newPost.UserID,
		CanComment: newPost.CanComment,
	}

	// Adds a new record to the data slice.
	i.InMemory = append(i.InMemory, postData)

	result = graphqlsh.Post{
		ID:         postData.ID,
		Text:       postData.Text,
		UserID:     postData.UserID,
		CanComment: postData.CanComment,
	}

	RecordCounter++

	return &result, nil
}

// Gets all posts from memory.
func (i *InMemoryStorage) FetchAllPosts(limit int, offset int) ([]graphqlsh.Post, error) {

	var result []graphqlsh.Post
	var toResult graphqlsh.Post

	for _, value := range i.InMemory {
		if value.PerentID == 0 {
			toResult = graphqlsh.Post{
				ID:         value.ID,
				Text:       value.Text,
				UserID:     value.UserID,
				CanComment: value.CanComment,
			}
			result = append(result, toResult)
		}
	}

	// Sorts the post slice in ascending order.
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	// Works with the length of the output slice.
	if len(result) < offset-1 {
		var empty []graphqlsh.Post
		return empty, &MyError{}
	}

	if len(result) < offset+limit {
		return result[offset-1:], nil
	}

	return result[offset-1 : offset+limit], nil
}

// Get a post and comments to it by ID from memory.
func (i *InMemoryStorage) FetchPostByiD(id int) (*graphqlsh.Post, error) {

	var result graphqlsh.Post

	for _, value := range i.InMemory {
		if value.ID == id {
			result = graphqlsh.Post{
				ID:         value.ID,
				Text:       value.Text,
				UserID:     value.UserID,
				CanComment: value.CanComment,
			}
		}
	}

	return &result, nil
}

// Get comments for a specific post from memory.
func (i *InMemoryStorage) FetchCommentsByPostID(id, limit, offset int) ([]graphqlsh.Comment, error) {

	var result []graphqlsh.Comment
	var toResult graphqlsh.Comment

	for _, value := range i.InMemory {
		if value.ID == id && !value.CanComment {
			return result, nil
		}
	}

	for _, value := range i.InMemory {
		if value.PostID == id {

			toResult = graphqlsh.Comment{
				ID:       value.ID,
				UserID:   value.UserID,
				Text:     value.Text,
				PostID:   value.PostID,
				PerentID: value.PerentID,
			}
			result = append(result, toResult)
		}
	}

	// Sorts the post slice in ascending order.
	sort.Slice(result, func(i, j int) bool {
		return result[i].PerentID < result[j].PerentID && result[i].ID > result[j].ID
	})

	// Works with the length of the output slice.
	if len(result) < offset-1 {
		var empty []graphqlsh.Comment
		return empty, &MyError{}
	}

	if len(result) < offset+limit {
		return result[offset-1:], nil
	}

	return result[offset-1 : offset+limit], nil
}

// Structure for creating a non-standard error.
type MyError struct{}

func (m *MyError) Error() string {
	return "There are not so many comments."
}
