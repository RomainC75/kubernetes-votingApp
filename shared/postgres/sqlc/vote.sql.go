package db

import (
	"context"
)

const createVote = `-- name: CreateVote :one
INSERT INTO votes (
    content, created_at, updated_at
) VALUES (
    $1, NOW(), NOW()
)
RETURNING id, content, created_at, updated_at
`

func (q *Queries) CreateVote(ctx context.Context, content string) (Vote, error) {
	row := q.db.QueryRowContext(ctx, createVote, content)
	var i Vote
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listVotes = `-- name: ListVotes :many
SELECT id, content, created_at, updated_at FROM votes
ORDER BY created_at
`

func (q *Queries) ListVotes(ctx context.Context) ([]Vote, error) {
	rows, err := q.db.QueryContext(ctx, listVotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vote{}
	for rows.Next() {
		var i Vote
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
