package main

import (
	"context"
	"log"

	"mail-sender/internal"
	"mail-sender/internal/config"
)

func main() {
	ctx := context.Background()

	var (
		dbConfig   *config.DatabaseConfig
		smtpConfig *config.SmtpConfig
	)

	if err := config.Read(dbConfig, "./database_config.json"); err != nil {
		log.Fatal(err)
	}

	if err := config.Read(smtpConfig, "./smtp_config.json"); err != nil {
		log.Fatal(err)
	}

	if err := internal.StartService(ctx, dbConfig, smtpConfig); err != nil {
		log.Fatal(err)
	}
}
