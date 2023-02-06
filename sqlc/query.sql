-- name: ListBookmarks :many
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
  created_at DESC;

-- name: FindBookmarks :many
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
  id = ANY(@ids::UUID[])
ORDER BY
  created_at DESC;

-- name: FindBookmarksByTags :many
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
  bht.tag_id = ANY(@ids::UUID[])
ORDER BY
  b.created_at DESC;

-- name: FindCommentsByBookmark :many
SELECT
  id,
  bookmark_id,
  comment,
  created_at,
  updated_at
FROM
  comments
WHERE
  bookmark_id = ANY(@ids::UUID[])
ORDER BY
  created_at DESC;

-- name: CreateBookmark :one
INSERT INTO bookmarks
  (url, title, description)
VALUES
  ($1, $2, $3)
RETURNING *;

-- name: UpdateBookmark :exec

-- name: DeleteBookmark :exec

-- name: CreateComment :exec

-- name: UpdateComment :exec

-- name: DeleteComment :exec
