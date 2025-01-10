package services

type VoteSrvI interface {
	SetVote(voteContent string) string
}
