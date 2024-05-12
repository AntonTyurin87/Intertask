package blogInterface

type Post struct {
	PID          int    `json:"id"`
	Text         string `json:"text"`
	PostAuthorID int    `json:"postauthorid"`
	CanComment   bool   `json:"cancomment"`
}

type Comment struct {
	CID      int    `json:"id"`
	UserID   int    `json:"uid"`
	Text     string `json:"text"`
	PostID   int    `json:"pid"`
	PerentID int    `json:"peid"`
}

/*
// Blog Interface
type Blog interface {
	FetchPostByiD(id int) (*Post, error)
}

// Blog interface function to retrieve posts by ID
func PostById(b Blog, id int) {
	b.FetchPostByiD(id) // From DB
}
*/
