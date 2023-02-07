package domain

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Bookmark struct {
	ID          uuid.UUID
	Url         string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Comments []*Comment
	Tags     []*Tag
}

type BookmarkInput struct {
	Url         string
	Title       string
	Description string
}

func NewBookmark(param *BookmarkInput) (*Bookmark, error) {
	if _, err := url.Parse(param.Url); err != nil {
		return nil, NewValidationError("Bookmark", "Url", err)
	}

	return &Bookmark{
		Url:         param.Url,
		Title:       param.Title,
		Description: param.Description,
	}, nil
}

type Comment struct {
	ID         uuid.UUID
	BookmarkID uuid.UUID
	Comment    string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Bookmark *Bookmark
}

type Tag struct {
	ID        uuid.UUID
	Name      string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time

	Bookmarks []*Bookmark
}
