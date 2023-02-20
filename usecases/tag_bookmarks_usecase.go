package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
)

type TagBookmarksUsecase struct {
	tags         TagsRepository
	tagBookmarks TagBookmarksRepository
}

func NewTagBookmarksUsecase(tags TagsRepository, tagBookmarks TagBookmarksRepository) *TagBookmarksUsecase {
	return &TagBookmarksUsecase{tags, tagBookmarks}
}

type TagBookmarksListRequest struct {
	TagID string `uri:"tag_id" binding:"required,uuid"`
}

func (u *TagBookmarksUsecase) List(ctx context.Context, req *TagBookmarksListRequest) ([]*domain.Bookmark, error) {
	tagID, err := uuid.Parse(req.TagID)
	if err != nil {
		return nil, err
	}

	if _, err := u.tags.Get(ctx, &tagID); err != nil {
		return nil, err
	}

	return u.tagBookmarks.List(ctx, &tagID)
}
