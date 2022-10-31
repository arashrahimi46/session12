package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"testMongo/internal/model"
	"testMongo/pkg"
	"time"
)

const (
	UsersCollectionName = "users"
)

type mongoRepository struct {
	config   *pkg.Config
	mongoCli *mongo.Client
}

func NewMongoRepository(config *pkg.Config, mongoCli *mongo.Client) model.Database {
	return &mongoRepository{
		config:   config,
		mongoCli: mongoCli,
	}
}

func (m *mongoRepository) InsertUser(user model.User) (lastInsertedId string, err error) {
	collection := m.
		mongoCli.
		Database(m.config.MongoConfig.DatabaseName).
		Collection(UsersCollectionName)
	user.Id = time.Now().String()
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(string), nil

}

func (m mongoRepository) DeleteUser() {
	//TODO implement me
	panic("implement me")
}

func (m mongoRepository) GetUser() {
	//TODO implement me
	panic("implement me")
}
