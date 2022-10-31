package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	httpHandler "testMongo/internal/handler/http"
	"testMongo/internal/repository"
	service2 "testMongo/internal/service"
	"testMongo/pkg"
	"time"
)

func main() {
	config := pkg.NewConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoCli, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoConfig.Uri))
	if err != nil {
		log.Fatal(err)
	}
	err = mongoCli.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	log.Println("connected to mongo")
	defer func() {
		if err = mongoCli.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	databaseRepository := repository.NewMongoRepository(config, mongoCli)

	service := service2.NewService(config, databaseRepository)

	httpHandler := httpHandler.NewHttpHandler(service)
	http.HandleFunc("/", httpHandler.GetHomePage)
	http.HandleFunc("/add/user", httpHandler.AddUser)
	log.Println("starting server at" + config.HttpPort)
	err = http.ListenAndServe(config.HttpPort, nil)
	if err != nil {
		panic(err)
	}
}
