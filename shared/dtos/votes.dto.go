package dtos

type VoteDto struct {
	Content string `json:"content" validate:"required"`
}
