// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/thoohv5/template/internal/data"
	app2 "github.com/thoohv5/template/internal/pkg/app"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/internal/pkg/log"
	"github.com/thoohv5/template/internal/router"
	"github.com/thoohv5/template/internal/server"
	"github.com/thoohv5/template/internal/service"
	"github.com/thoohv5/template/pkg/app"
)

// Injectors from wire.go:

// initApp init application.
func initApp() (app.IApp, func(), error) {
	iConfig := config.New()
	logger := log.New(iConfig)
	iData, cleanup, err := data.New(iConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	iService := service.New(iConfig, logger, iData)
	iServer := server.New(iConfig, logger, iService)
	registerRouter := router.New(iConfig, iServer)
	iApp := app2.New(iConfig, registerRouter)
	return iApp, func() {
		cleanup()
	}, nil
}