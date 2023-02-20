// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addTagToBookmark = `-- name: AddTagToBookmark :exec
INSERT INTO bookmark_has_tags
  (bookmark_id, tag_id)
VALUES
  ($1, $2)
`

type AddTagToBookmarkParams struct {
	BookmarkID uuid.UUID
	TagID      uuid.UUID
}

func (q *Queries) AddTagToBookmark(ctx context.Context, arg AddTagToBookmarkParams) error {
	_, err := q.db.ExecContext(ctx, addTagToBookmark, arg.BookmarkID, arg.TagID)
	return err
}

const createBookmark = `-- name: CreateBookmark :one
INSERT INTO bookmarks
  (url, title, description)
VALUES
  ($1, $2, $3)
RETURNING id, url, title, description, created_at, updated_at
`

type CreateBookmarkParams struct {
	Url         string
	Title       string
	Description string
}

func (q *Queries) CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRowContext(ctx, createBookmark, arg.Url, arg.Title, arg.Description)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createComment = `-- name: CreateComment :one
INSERT INTO comments
  (bookmark_id, body)
VALUES
  ($1, $2)
RETURNING id, bookmark_id, body, created_at, updated_at
`

type CreateCommentParams struct {
	BookmarkID uuid.UUID
	Body       string
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.BookmarkID, arg.Body)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.BookmarkID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTag = `-- name: CreateTag :one
INSERT INTO tags
  (name, color)
VALUES
  ($1, $2)
RETURNING id, name, color, created_at, updated_at
`

type CreateTagParams struct {
	Name  string
	Color string
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createTag, arg.Name, arg.Color)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Color,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBookmark = `-- name: DeleteBookmark :exec
DELETE FROM
  bookmarks
WHERE
  id = $1
`

func (q *Queries) DeleteBookmark(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBookmark, id)
	return err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM
  comments
WHERE
  id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM
  tags
WHERE
  id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTag, id)
	return err
}

const findBookmarksByTag = `-- name: FindBookmarksByTag :many
SELECT
  b.id,
  b.url,
  b.title,
  b.description,
  b.created_at,
  b.updated_at
FROM
  bookmarks AS b
LEFT JOIN bookmark_has_tags AS bht ON b.id = bht.bookmark_id
LEFT JOIN tags AS t ON bht.tag_id = t.id
WHERE
  bht.tag_id = $1
ORDER BY
  b.created_at DESC
`

func (q *Queries) FindBookmarksByTag(ctx context.Context, tagID uuid.UUID) ([]Bookmark, error) {
	rows, err := q.db.QueryContext(ctx, findBookmarksByTag, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Bookmark{}
	for rows.Next() {
		var i Bookmark
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findTagsByBookmark = `-- name: FindTagsByBookmark :many
SELECT
  t.id,
  t.name,
  t.color,
  t.created_at,
  t.updated_at
FROM
  tags AS t
LEFT JOIN bookmark_has_tags AS bht ON t.id = bht.tag_id
LEFT JOIN bookmarks AS b ON bht.bookmark_id = b.id
WHERE
  bht.bookmark_id = $1
ORDER BY
  b.created_at DESC
`

func (q *Queries) FindTagsByBookmark(ctx context.Context, bookmarkID uuid.UUID) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, findTagsByBookmark, bookmarkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Color,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookmark = `-- name: GetBookmark :one
SELECT
  id,
  url,
  title,
  description,
  created_at,
  updated_at
FROM
  bookmarks
WHERE
  id = $1
ORDER BY
  created_at DESC
`

func (q *Queries) GetBookmark(ctx context.Context, id uuid.UUID) (Bookmark, error) {
	row := q.db.QueryRowContext(ctx, getBookmark, id)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getComment = `-- name: GetComment :one
SELECT
  id,
  bookmark_id,
  body,
  created_at,
  updated_at
FROM
  comments
WHERE
  id = $1
ORDER BY
  created_at DESC
`

func (q *Queries) GetComment(ctx context.Context, id uuid.UUID) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.BookmarkID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTag = `-- name: GetTag :one
SELECT
  id,
  name,
  color,
  created_at,
  updated_at
FROM
  tags
WHERE
  id = $1
ORDER BY
  created_at DESC
`

func (q *Queries) GetTag(ctx context.Context, id uuid.UUID) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTag, id)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Color,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBookmarks = `-- name: ListBookmarks :many
SELECT
  id,
  url,
  title,
  description,
  created_at,
  updated_at
FROM
  bookmarks
ORDER BY
  created_at DESC
`

func (q *Queries) ListBookmarks(ctx context.Context) ([]Bookmark, error) {
	rows, err := q.db.QueryContext(ctx, listBookmarks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Bookmark{}
	for rows.Next() {
		var i Bookmark
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listComments = `-- name: ListComments :many
SELECT
  id,
  bookmark_id,
  body,
  created_at,
  updated_at
FROM
  comments
WHERE
  bookmark_id = $1
ORDER BY
  created_at DESC
`

func (q *Queries) ListComments(ctx context.Context, bookmarkID uuid.UUID) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listComments, bookmarkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.BookmarkID,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTags = `-- name: ListTags :many
SELECT
  id,
  name,
  color,
  created_at,
  updated_at
FROM
  tags
ORDER BY
  created_at DESC
`

func (q *Queries) ListTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Color,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeTagFromBookmark = `-- name: RemoveTagFromBookmark :exec
DELETE FROM 
  bookmark_has_tags
WHERE
  bookmark_id = $1 AND tag_id =  $2
`

type RemoveTagFromBookmarkParams struct {
	BookmarkID uuid.UUID
	TagID      uuid.UUID
}

func (q *Queries) RemoveTagFromBookmark(ctx context.Context, arg RemoveTagFromBookmarkParams) error {
	_, err := q.db.ExecContext(ctx, removeTagFromBookmark, arg.BookmarkID, arg.TagID)
	return err
}

const updateBookmark = `-- name: UpdateBookmark :one
UPDATE bookmarks
SET
  url = $2,
  title = $3,
  description = $4,
  updated_at = $5
WHERE
  id = $1
RETURNING id, url, title, description, created_at, updated_at
`

type UpdateBookmarkParams struct {
	ID          uuid.UUID
	Url         string
	Title       string
	Description string
	UpdatedAt   time.Time
}

func (q *Queries) UpdateBookmark(ctx context.Context, arg UpdateBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRowContext(ctx, updateBookmark,
		arg.ID,
		arg.Url,
		arg.Title,
		arg.Description,
		arg.UpdatedAt,
	)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTag = `-- name: UpdateTag :one
UPDATE tags
SET
  name = $2,
  color = $3,
  updated_at = $4
WHERE
  id = $1
RETURNING id, name, color, created_at, updated_at
`

type UpdateTagParams struct {
	ID        uuid.UUID
	Name      string
	Color     string
	UpdatedAt time.Time
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, updateTag,
		arg.ID,
		arg.Name,
		arg.Color,
		arg.UpdatedAt,
	)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Color,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
