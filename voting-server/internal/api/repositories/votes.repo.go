package repositories

import (
	"context"
	"shared/dtos"
	db "shared/postgres/sqlc"
)

type VotesRepo struct {
	Store *db.Store
}

func NewVotesRepo() *VotesRepo {
	return &VotesRepo{
		Store: db.GetConnection(),
	}
}

func (vr *VotesRepo) SetNewVote(voteDto dtos.VoteDto) (db.Vote, error) {
	ctx := context.Background()
	return (*vr.Store).CreateVote(ctx, voteDto.Content)
}
