package services

import (
	"fmt"
	"shared/dtos"
	db "shared/postgres/sqlc"
	"voting-server/internal/api/repositories"
)

type VoteSrv struct {
	voteRepo repositories.VoteRepoI
}

func NewVoteSrv() *VoteSrv {
	return &VoteSrv{
		voteRepo: repositories.NewVotesRepo(),
	}
}

func (voteSrv *VoteSrv) SetVote(voteDto dtos.VoteDto) (db.Vote, error) {
	fmt.Println("---->", voteSrv.voteRepo)
	return voteSrv.voteRepo.SetNewVote(voteDto)
}
