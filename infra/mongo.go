package infra 

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(url string) (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		cancel()
		log.Fatal("Connect error: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		cancel()
		log.Fatal("Ping error: ", err)
	}
	log.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

func DisconnectDB(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal("Disconnect error: ", err)
			panic(err)
		}
	}()
}
func InsertMany[T any](client *mongo.Client, ctx context.Context, dbName, collName string, data[] T) {
	col := client.Database(dbName).Collection(collName)
	docs := make([]interface{}, len(data))
	for i, d := range data {
		docs[i] = d
	}
	result, err := col.InsertMany(ctx, docs)
	if err != nil {
		log.Fatal("InsertMany error: ", err)
	}
	log.Printf("Inserted %d documents into %s.%s\n", len(result.InsertedIDs), dbName, collName)


}
