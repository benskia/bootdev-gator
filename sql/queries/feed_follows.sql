-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    *, (
        SELECT
            name
        FROM
            users
        WHERE
            id = $4), (
        SELECT
            name
        FROM
            feeds
        WHERE
            id = $5);

