package db

import (
	"context"
)

type Querier interface {
	CreateVote(ctx context.Context, content string) (Vote, error)
	ListVotes(ctx context.Context) ([]Vote, error)
}

var _ Querier = (*Queries)(nil)
