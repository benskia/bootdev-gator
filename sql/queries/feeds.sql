-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: GetFeeds :many
SELECT
    name,
    url,
    user_id
FROM
    feeds;

-- name: GetFeedByURL :one
SELECT
    *
FROM
    feeds
WHERE
    url = $1;

-- name: GetFeedByID :one
SELECT
    *
FROM
    feeds
WHERE
    id = $1;

