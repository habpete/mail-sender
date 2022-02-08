package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"mail-sender/internal/config"
	"mail-sender/internal/sender"
	"mail-sender/internal/service"
	"mail-sender/internal/storage"
)

func StartService(ctx context.Context,
	dbConfig *config.DatabaseConfig,
	smtpConfig *config.SmtpConfig) error {
	if dbConfig == nil || smtpConfig == nil {
		return errors.New("bad config")
	}

	strg := storage.New(dbConfig)
	sndr, err := sender.New(smtpConfig)
	if err != nil {
		return fmt.Errorf("start failed: %w", err)
	}
	defer sndr.Shutdown()

	cCtx, cancel := context.WithCancel(ctx)
	srv := service.New(strg, sndr)
	go func() {
		srv.StartReader(cCtx)
	}()
	go func() {
		srv.StartSender(cCtx)
	}()

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	log.Print("awaiting signal")
	sig := <-sigs
	log.Printf("receive signal %s", sig.String())

	cancel()

	return nil
}
