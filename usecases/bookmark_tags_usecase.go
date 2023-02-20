package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
)

type BookmarkTagsUsecase struct {
	bookmarks    BookmarksRepository
	tags         TagsRepository
	bookmarkTags BookmarkTagsRepository
}

func NewBookmarkTagsUsecase(bookmarks BookmarksRepository, tags TagsRepository, bookmarkTags BookmarkTagsRepository) *BookmarkTagsUsecase {
	return &BookmarkTagsUsecase{bookmarks, tags, bookmarkTags}
}

type BookmarkTagsListRequest struct {
	BookmarkID string `uri:"bookmark_id" binding:"required,uuid"`
}

func (u *BookmarkTagsUsecase) List(ctx context.Context, req *BookmarkTagsListRequest) ([]*domain.Tag, error) {
	bookmarkID, err := uuid.Parse(req.BookmarkID)
	if err != nil {
		return nil, err
	}

	if _, err := u.bookmarks.Get(ctx, &bookmarkID); err != nil {
		return nil, err
	}

	return u.bookmarkTags.List(ctx, &bookmarkID)
}

type BookmarkTagsAddRequest struct {
	BookmarkID string `uri:"bookmark_id" binding:"required,uuid"`
	TagID      string `uri:"tag_id" binding:"required,uuid"`
}

func (u *BookmarkTagsUsecase) Add(ctx context.Context, req *BookmarkTagsAddRequest) ([]*domain.Tag, error) {
	bookmarkID, err := uuid.Parse(req.BookmarkID)
	if err != nil {
		return nil, err
	}
	tagID, err := uuid.Parse(req.TagID)
	if err != nil {
		return nil, err
	}

	_, err = u.bookmarks.Get(ctx, &bookmarkID)
	if err != nil {
		return nil, err
	}
	_, err = u.tags.Get(ctx, &tagID)
	if err != nil {
		return nil, err
	}

	if err := u.bookmarkTags.Add(ctx, &bookmarkID, &tagID); err != nil {
		return nil, err
	}

	return u.bookmarkTags.List(ctx, &bookmarkID)
}

type BookmarkTagsRemoveRequest struct {
	BookmarkID string `uri:"bookmark_id" binding:"required,uuid"`
	TagID      string `uri:"tag_id" binding:"required,uuid"`
}

func (u *BookmarkTagsUsecase) Remove(ctx context.Context, req *BookmarkTagsRemoveRequest) ([]*domain.Tag, error) {
	bookmarkID, err := uuid.Parse(req.BookmarkID)
	if err != nil {
		return nil, err
	}
	tagID, err := uuid.Parse(req.TagID)
	if err != nil {
		return nil, err
	}

	_, err = u.bookmarks.Get(ctx, &bookmarkID)
	if err != nil {
		return nil, err
	}
	_, err = u.tags.Get(ctx, &tagID)
	if err != nil {
		return nil, err
	}

	if err := u.bookmarkTags.Remove(ctx, &bookmarkID, &tagID); err != nil {
		return nil, err
	}

	return u.bookmarkTags.List(ctx, &bookmarkID)
}
