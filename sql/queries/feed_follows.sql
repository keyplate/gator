-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES(
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds ON (inserted_feed_follows.feed_id = feeds.id)
INNER JOIN users ON (inserted_feed_follows.user_id = users.id);

-- name: GetFeedFollowsByUser :many
SELECT feed_follows.*, f.name, u.name as username
FROM feed_follows
INNER JOIN users u on (u.id = feed_follows.user_id)
INNER JOIN feeds f on (f.id = feed_follows.feed_id)
WHERE u.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING users, feeds
WHERE feed_follows.user_id = users.id
AND feed_follows.feed_id = feeds.id
AND users.name = $1
AND feeds.url = $2;
