// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Loner1024/uniix.io/configs"
	"github.com/Loner1024/uniix.io/internal/domain/blog"
	"github.com/Loner1024/uniix.io/internal/server"
	"github.com/Loner1024/uniix.io/internal/services"
	"github.com/Loner1024/uniix.io/internal/store/firebase"
	"github.com/Loner1024/uniix.io/logger"
)

// Injectors from wire.go:

func wireApp() (*server.Server, error) {
	config, err := configs.InitConfigs()
	if err != nil {
		return nil, err
	}
	repo, err := firebase.NewStore(config)
	if err != nil {
		return nil, err
	}
	useCase := blog.NewUseCase(repo)
	service := services.NewService(useCase)
	sugaredLogger, err := logger.New(config)
	if err != nil {
		return nil, err
	}
	serverServer := server.NewServer(service, config, sugaredLogger)
	return serverServer, nil
}
