package inmemory

import (
	"intertask/graphqlsh"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewPost(t *testing.T) {

	var InMe []InMemoryType
	storage := NewInMemory(InMe)

	//In Data
	notEmptyPost := graphqlsh.Post{
		ID:         1,
		Text:       "New text!",
		UserID:     2,
		CanComment: true,
	}

	//Out Data
	expectedNotEmpty := graphqlsh.Post{
		ID:         1,
		Text:       "New text!",
		UserID:     2,
		CanComment: true,
	}

	actualNotEmpty, err := storage.CreateNewPost(&notEmptyPost)

	if err != nil {
		t.Errorf("Should not produce an error")
	}

	// Comparing the fields of structures from the expected structure obtained from the structure function and from the structure instance in memory.
	assert.Equal(t, expectedNotEmpty.ID, actualNotEmpty.ID, storage.InMemory[0].ID)
	assert.Equal(t, expectedNotEmpty.UserID, actualNotEmpty.UserID, storage.InMemory[0].UserID)
	assert.Equal(t, expectedNotEmpty.Text, actualNotEmpty.Text, storage.InMemory[0].Text)
	assert.Equal(t, expectedNotEmpty.CanComment, actualNotEmpty.CanComment, storage.InMemory[0].CanComment)

}
