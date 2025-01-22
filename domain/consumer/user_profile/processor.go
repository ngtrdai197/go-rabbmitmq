package userprofile

import "context"

const UserProfileTopic = "user_profile"

type Processor interface {
	Processor(ctx context.Context, topic string, message []byte)
}

type processor struct {
	h *Handler
}

func NewProcessor(h *Handler) Processor {
	return &processor{h: h}
}

func (p *processor) Processor(ctx context.Context, topic string, message []byte) {
	if topic != UserProfileTopic {
		return
	}

	// Do something with the message
	p.h.Handle(ctx, topic, message)
}
