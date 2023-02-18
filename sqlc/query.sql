-- name: GetBookmark :one
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
  created_at DESC;

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

-- name: DeleteBookmark :exec
DELETE FROM
  bookmarks
WHERE
  id = $1;

-- name: UpdateBookmark :one
UPDATE bookmarks
SET
  url = $2,
  title = $3,
  description = $4,
  updated_at = $5
WHERE
  id = $1
RETURNING *;

-- name: GetTag :one
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
  created_at DESC;

-- name: ListTags :many
SELECT
  id,
  name,
  color,
  created_at,
  updated_at
FROM
  tags
ORDER BY
  created_at DESC;

-- name: CreateTag :one
INSERT INTO tags
  (name, color)
VALUES
  ($1, $2)
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM
  tags
WHERE
  id = $1;

-- name: UpdateTag :one
UPDATE tags
SET
  name = $2,
  color = $3,
  updated_at = $4
WHERE
  id = $1
RETURNING *;
