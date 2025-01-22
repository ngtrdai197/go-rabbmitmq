package userprofile

import "context"

type Handler struct {
	// defined business => UserProfileBusiness
}

func NewHandler() *Handler {
	return &Handler{}
}

func (p *Handler) Handle(ctx context.Context, topic string, message []byte) {
	// Unmarshal the message to a struct
}
