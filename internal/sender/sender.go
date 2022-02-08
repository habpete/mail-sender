package sender

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"mail-sender/internal/config"
)

type Sender struct {
	client *smtp.Client

	config *config.SmtpConfig
}

func New(smtpConfig *config.SmtpConfig) (*Sender, error) {
	client, err := connect(smtpConfig)
	if err != nil {
		return nil, err
	}

	return &Sender{
		client: client,
		config: smtpConfig,
	}, nil
}

func connect(smtpConfig *config.SmtpConfig) (*smtp.Client, error) {
	client, err := smtp.Dial(smtpConfig.RemoteHost)
	if err != nil {
		return nil, fmt.Errorf("connect to smtp failed: %w", err)
	}

	auth := smtp.PlainAuth("", smtpConfig.User, smtpConfig.Password, smtpConfig.RemoteHost)
	if err := client.Auth(auth); err != nil {
		return nil, fmt.Errorf("authorization to smtp failed: %w", err)
	}

	return client, nil
}

func (i *Sender) Reconnect(ctx context.Context) {
	ticker := time.NewTicker(i.config.ReconnectInterval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			client, err := connect(i.config)
			if err != nil {
				log.Printf("reconnect to %s failed", i.config.RemoteHost)
				continue
			}
			i.client = client
		}
	}
}

func (i *Sender) Publish(ctx context.Context, data interface{}) error {
	return nil
}

func (i *Sender) Shutdown() error {
	return i.client.Close()
}
