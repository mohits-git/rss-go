-- name: CreatePost :one
INSERT INTO posts (id, title, url, description, published_at, feed_id)
VALUES (
  $1, 
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetPostsForUser :one
SELECT * FROM posts
ORDER BY created_at DESC
LIMIT $1;
