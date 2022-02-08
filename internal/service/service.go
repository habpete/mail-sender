package service

import (
	"context"
	"log"
	"mail-sender/internal/sender"
	"mail-sender/internal/storage"
	"time"
)

type Service struct {
	storage *storage.Storage

	sender *sender.Sender

	workInterval time.Duration
}

func New(dbStorage *storage.Storage, mailSender *sender.Sender) *Service {
	return &Service{
		storage: dbStorage,
		sender:  mailSender,
	}
}

func (i *Service) StartReader(ctx context.Context) {
	ticker := time.NewTicker(i.workInterval)
	for {
		select {
		case <-ctx.Done():
			log.Print("shutdown reader")
			return
		case <-ticker.C:
			//
		}
	}
}

func (i *Service) StartSender(ctx context.Context) {
	ticker := time.NewTicker(i.workInterval)
	for {
		select {
		case <-ctx.Done():
			log.Print("shutdown sender")
			return
		case <-ticker.C:
			//
		}
	}
}
