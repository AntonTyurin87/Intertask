package inmemory

import (
	"fmt"
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
var RecordCounter = 1

// Structure for accessing to memory.
type InMemoryStorage struct {
	InMemory []InMemoryType
}

// Creating a new instance of a structure for accessing to memory.
func NewInMemory(InMe []InMemoryType) *InMemoryStorage {
	return &InMemoryStorage{InMemory: InMe}
}

// Makes a change to the post entry about the ability to comment the post in memory.
func (i *InMemoryStorage) UpdatePost(correctPost *graphqlsh.Post) (*graphqlsh.Post, error) {

	var err error
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

				i.InMemory = append(i.InMemory[:j], i.InMemory[j+1:]...)
				i.InMemory = append(i.InMemory, postData)
			}

			return &result, err
		}
	}
	return &result, err // If a post to update on is not found, then return an empty value.
}

// Creates a record of a new comment to post in memory.
func (i *InMemoryStorage) CreateNewComment(newComment *graphqlsh.Comment) (*graphqlsh.Comment, error) {

	var err error
	var result graphqlsh.Comment

	//
	for _, value := range i.InMemory {
		if value.ID == newComment.PostID {
			if !value.CanComment {
				return &result, err
			} else {

				var trimText string

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

				return &result, err
			}
		}
	}
	return &result, err // If a post to comment on is not found, then return an empty value.
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

	var err error
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

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result, err
}

// Get a post and comments to it by ID from memory.
func (i *InMemoryStorage) FetchPostByiD(id int) (*graphqlsh.Post, error) {
	var err error
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

	return &result, err
}

// Get comments for a specific post from memory.
func (i *InMemoryStorage) FetchCommentsByPostID(id, limit, offset int) ([]graphqlsh.Comment, error) {

	var err error
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

	sort.Slice(result, func(i, j int) bool {
		return result[i].PerentID < result[j].PerentID && result[i].ID < result[j].ID
	})

	if len(result) < offset-1 {
		var empty []graphqlsh.Comment
		err = fmt.Errorf("There are not so many comments.", err)
		return empty, err
	}

	if len(result) < offset+limit {
		return result[offset:], err
	}

	return result[offset : offset+limit], err
}
