package infra

import (
	"context"

	"minibank/domain"
	"minibank/logger"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, mongoConf *domain.MongoConfig) (*mongo.Client, error) {
	opts := options.ClientOptions{}
	client, err := mongo.Connect(ctx, opts.ApplyURI(mongoConf.MongoURL))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}

func Disconnect(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	log := logger.NewLogger(zerolog.InfoLevel)
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Err(err).Msg("error when disconnecting from MongoDB")
			panic(err)
		}
	}()
}

func InsertMany[T any](client *mongo.Client, ctx context.Context, dbName, collName string, data []T) {
	log := logger.NewLogger(zerolog.InfoLevel)
	col := client.Database(dbName).Collection(collName)
	docs := make([]interface{}, len(data))
	for i, d := range data {
		docs[i] = d
	}
	result, err := col.InsertMany(ctx, docs)
	if err != nil {
		log.Err(err).Msg("error when inserting documents into MongoDB")
	}
	log.Info().Msgf("Inserted %d documents into %s.%s\n", len(result.InsertedIDs), dbName, collName)
}
