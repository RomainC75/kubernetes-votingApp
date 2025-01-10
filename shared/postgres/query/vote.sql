-- name: ListVotes :many
SELECT * FROM votes
ORDER BY created_at;

-- name: CreateVote :one
INSERT INTO votes (
    content, created_at, updated_at
) VALUES (
    $1, NOW(), NOW()
)
RETURNING *;
