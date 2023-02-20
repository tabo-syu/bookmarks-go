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

-- name: FindBookmarksByTag :many
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
  b.created_at DESC;

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

-- name: FindTagsByBookmark :many
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
  b.created_at DESC;

-- name: CreateTag :one
INSERT INTO tags
  (name, color)
VALUES
  ($1, $2)
RETURNING *;

-- name: AddTagToBookmark :exec
INSERT INTO bookmark_has_tags
  (bookmark_id, tag_id)
VALUES
  ($1, $2);

-- name: RemoveTagFromBookmark :exec
DELETE FROM 
  bookmark_has_tags
WHERE
  bookmark_id = $1 AND tag_id =  $2;

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

-- name: GetComment :one
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
  created_at DESC;

-- name: ListComments :many
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
  created_at DESC;

-- name: CreateComment :one
INSERT INTO comments
  (bookmark_id, body)
VALUES
  ($1, $2)
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM
  comments
WHERE
  id = $1;
