package main

import (
	// "minibank/infra"
	// "minibank/seeder"
	// "os"
	"context"
	"minibank/domain"
	"minibank/infra"
	"minibank/logger"

	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(zerolog.InfoLevel)
	log.Info().Msg("Starting the application...")
	mongoConf := domain.LoadMongoConfig()
	client, err := infra.Connect(ctx, mongoConf)
	if err != nil {
		log.Err(err).Msg("Failed to connect to MongoDB")
		return
	}
	db := client.Database("minibank")
	log.Info().Msg(db.Name())
	defer infra.Disconnect(client, ctx, func() {})
}
