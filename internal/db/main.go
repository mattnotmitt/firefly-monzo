package db

import (
	"context"
	"log"
	"time"

	"github.com/mattnotmitt/firefly-monzo/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Manager is an object holding functions and objects related to DB operations
type Manager struct {
	client *mongo.Client
	close  *context.CancelFunc
}

// Init connects to the mongodb and returns a Manager object
func Init(cfg config.Config) Manager {
	mongoURL := cfg.MongoURL
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return Manager{client: client, close: &cancel}
}
