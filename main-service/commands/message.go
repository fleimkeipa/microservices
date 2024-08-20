package commands

import "order-service/repositories/interfaces"

type MessageCommandHandlers struct {
	repo interfaces.MessageInterfaces
}

func NewMessageCommandHandlers(repo interfaces.MessageInterfaces) *MessageCommandHandlers {
	return &MessageCommandHandlers{
		repo: repo,
	}
}

func (rc *MessageCommandHandlers) Send(subj string, data string) error {
	return rc.repo.Send(subj, data)
}
