package domain

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Bookmark struct {
	ID          uuid.UUID `json:"id"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Comments []*Comment `json:"comments,omitempty"`
	Tags     []*Tag     `json:"tags,omitempty"`
}

type BookmarkCreateRequest struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewBookmark(urlStr string, title string, description string) (*Bookmark, error) {
	if _, err := url.ParseRequestURI(urlStr); err != nil {
		return nil, NewValidationError("Bookmark", "Url", err)
	}

	return &Bookmark{
		Url:         urlStr,
		Title:       title,
		Description: description,
	}, nil
}

type Comment struct {
	ID         uuid.UUID `json:"id"`
	BookmarkID uuid.UUID `json:"-"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Bookmark *Bookmark `json:"bookmark,omitempty"`
}

type Tag struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Bookmarks []*Bookmark `json:"bookmarks,omitempty"`
}
