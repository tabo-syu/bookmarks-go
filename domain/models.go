package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

type Bookmark struct {
	ID          uuid.UUID `json:"id"`
	Url         string    `json:"url" validate:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Comments []*Comment `json:"comments,omitempty"`
	Tags     []*Tag     `json:"tags,omitempty"`
}

func NewBookmark(url string, title string, description string) (*Bookmark, error) {
	bookmark := Bookmark{
		Url:         url,
		Title:       title,
		Description: description,
	}

	err := validate.Struct(bookmark)
	if err != nil {
		return nil, NewValidationError("Bookmark", err)
	}

	return &bookmark, nil
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
	Color     string    `json:"color" validate:"hexcolor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Bookmarks []*Bookmark `json:"bookmarks,omitempty"`
}
