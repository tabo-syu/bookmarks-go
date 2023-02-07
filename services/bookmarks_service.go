package services

import "github.com/tabo-syu/bookmarks/gateways"

type BookmarksService struct {
	bookmarks gateways.BookmarksRepository
}

func NewBookmarksService(bookmarks gateways.BookmarksRepository) *BookmarksService {
	return &BookmarksService{bookmarks}
}
