package repository

import (
	"context"
	"log"
	"time"

	"github.com/iBoBoTi/graphql-go-server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	Database = "graphql"
	VideoCollection = "videos"
)

type VideoRepository interface{
	Save(ctx context.Context,video *model.Video) error
	FindAll(ctx context.Context) ([]*model.Video, error)
}

type database struct{
	client *mongo.Client
}

func New() VideoRepository{
	clientOpt := options.Client().ApplyURI("mongodb://localhost:27017/graphql")
	clientOpt = clientOpt.SetMaxPoolSize(50)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) 
	defer cancel()

	dbClient, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
        log.Fatalf("error pinging database: %v", err)
    }
	log.Println("connected to database")

	return &database{
		client: dbClient,
	}
}
func (db *database) Save(ctx context.Context,video *model.Video) error{
	_, err := db.client.Database(Database).Collection(VideoCollection).InsertOne(ctx,video)
	if err != nil {
		return err
	}
	return nil
}

func (db *database) FindAll(ctx context.Context) ([]*model.Video, error) {

	cursor, err := db.client.Database(Database).Collection(VideoCollection).Find(ctx,bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []*model.Video
	for cursor.Next(ctx) {
		var v *model.Video
		err := cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	} 

	return result, nil
}