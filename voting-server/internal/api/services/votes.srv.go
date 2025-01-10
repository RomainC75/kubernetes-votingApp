package services

type VoteSrv struct {
}

func NewVoteSrv() *VoteSrv {
	return &VoteSrv{}
}

func (s *VoteSrv) SetVote(voteContent string) string {
	return voteContent
}
