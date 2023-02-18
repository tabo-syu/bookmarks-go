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
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Bookmark *Bookmark `json:"bookmark,omitempty"`
}

func NewComment(bookmarkID *uuid.UUID, body string) (*Comment, error) {
	comment := Comment{
		BookmarkID: *bookmarkID,
		Body:       body,
	}

	err := validate.Struct(comment)
	if err != nil {
		return nil, NewValidationError("Comment", err)
	}

	return &comment, nil
}

type Tag struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color" validate:"hexcolor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Bookmarks []*Bookmark `json:"bookmarks,omitempty"`
}

func NewTag(name string, color string) (*Tag, error) {
	tag := Tag{
		Name:  name,
		Color: color,
	}

	err := validate.Struct(tag)
	if err != nil {
		return nil, NewValidationError("Tag", err)
	}

	return &tag, nil
}
