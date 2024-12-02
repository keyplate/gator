-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE user_id = $1;

-- name: GetFeeds :many
SELECT f.name, f.url, u.name as username 
FROM feeds f 
INNER JOIN users u
ON (f.user_id = u.id);

-- name: GetFeedByUrl :one
SELECT id, name, url
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET updated_at = $1, last_fetched_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at DESC NULLS FIRST
LIMIT 1;
