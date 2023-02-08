// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

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

const findBookmarks = `-- name: FindBookmarks :many
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
  id = ANY($1::UUID[])
ORDER BY
  created_at DESC
`

func (q *Queries) FindBookmarks(ctx context.Context, ids []uuid.UUID) ([]Bookmark, error) {
	rows, err := q.db.QueryContext(ctx, findBookmarks, pq.Array(ids))
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

const findBookmarksByTags = `-- name: FindBookmarksByTags :many
SELECT
  b.id,
  b.url,
  b.title,
  b.description,
  b.created_at,
  b.updated_at
FROM
  bookmark_has_tags AS bht
LEFT JOIN bookmarks AS b ON b.id = bht.bookmark_id
LEFT JOIN tags AS t ON t.id = bht.tag_id
WHERE
  bht.tag_id = ANY($1::UUID[])
ORDER BY
  b.created_at DESC
`

type FindBookmarksByTagsRow struct {
	ID          uuid.NullUUID
	Url         sql.NullString
	Title       sql.NullString
	Description sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

func (q *Queries) FindBookmarksByTags(ctx context.Context, ids []uuid.UUID) ([]FindBookmarksByTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, findBookmarksByTags, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindBookmarksByTagsRow{}
	for rows.Next() {
		var i FindBookmarksByTagsRow
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

const findCommentsByBookmark = `-- name: FindCommentsByBookmark :many
SELECT
  id,
  bookmark_id,
  comment,
  created_at,
  updated_at
FROM
  comments
WHERE
  bookmark_id = ANY($1::UUID[])
ORDER BY
  created_at DESC
`

func (q *Queries) FindCommentsByBookmark(ctx context.Context, ids []uuid.UUID) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, findCommentsByBookmark, pq.Array(ids))
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
			&i.Comment,
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
