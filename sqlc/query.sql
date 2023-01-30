-- name: GetBookmark :one
SELECT
  b.url AS url,
  b.created_at AS bookmark_created_at,
  t.name AS tag_name,
  c.comment AS comment,
  c.updated_at AS comment_updated_at
FROM
  bookmarks AS b
  LEFT JOIN bookmark_has_tags AS bht ON b.id = bht.bookmark_id
  LEFT JOIN tags AS t ON bht.bookmark_id = t.id
  LEFT JOIN bookmark_has_comments AS bhc ON b.id = bhc.bookmark_id
  LEFT JOIN comments AS c ON bhc.bookmark_id = c.id
WHERE
  bookmarks.id = $1;

-- name: ListBookmarks :many
SELECT
  b.url AS url,
  b.created_at AS bookmark_created_at,
  t.name AS tag_name,
  c.comment AS comment,
  c.updated_at AS comment_updated_at
FROM
  bookmarks AS b
  LEFT JOIN bookmark_has_tags AS bht ON b.id = bht.bookmark_id
  LEFT JOIN tags AS t ON bht.bookmark_id = t.id
  LEFT JOIN bookmark_has_comments AS bhc ON b.id = bhc.bookmark_id
  LEFT JOIN comments AS c ON bhc.bookmark_id = c.id;

-- name: CreateBookmark :exec

-- name: UpdateBookmark :exec

-- name: DeleteBookmark :exec

-- name: CreateComment :exec

-- name: UpdateComment :exec

-- name: DeleteComment :exec
