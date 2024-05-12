-- name: CreateFeedFollow :one
INSERT INTO feed_follow (id, created_at, updated_at, feed_id, user_id)
VALUES ($1,$2,$3,$4,$5)

    RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow
WHERE id = $1;


-- name: GetFeedFollow :many
SELECT * from feed_follow
WHERE user_id = $1;