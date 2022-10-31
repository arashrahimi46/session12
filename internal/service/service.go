package service

import (
	"testMongo/internal/model"
	"testMongo/pkg"
)

type service struct {
	config             *pkg.Config
	databaseRepository model.Database
}

func NewService(config *pkg.Config, databaseRepository model.Database) model.Service {
	return &service{
		config:             config,
		databaseRepository: databaseRepository,
	}
}

func (s *service) GetHomePage() string {
	return "Hello World From Service"
}

func (s *service) AddUser(user model.User) error {
	_, err := s.databaseRepository.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}
