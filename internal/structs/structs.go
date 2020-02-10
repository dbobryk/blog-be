package structs

import "time"

// BlogPost defines the structure for a blog post
type BlogPost struct {
	PostID    string    `json:"post_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Author    string    `json:"author,omitempty"`
	Content   string    `json:"content,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	Updated   time.Time `json:"updated,omitempty"`
	Published bool      `json:"published,omitempty"`
}
