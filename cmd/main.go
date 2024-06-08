package main

import (
	"github.com/elyarsadig/smart-home-iot/config"
	"github.com/elyarsadig/smart-home-iot/internal/domain/repository"
	"github.com/elyarsadig/smart-home-iot/internal/domain/service"
	"github.com/elyarsadig/smart-home-iot/internal/infrastructure/db/sqlite"
	httpPackage "github.com/elyarsadig/smart-home-iot/internal/infrastructure/http"
	"github.com/elyarsadig/smart-home-iot/internal/infrastructure/http/handler"
)

func main() {
	cfg := config.LoadConfig()
	
	db := sqlite.InitDB(&cfg.Database)

	roomRepo := repository.NewRoom(db)
	deviceRepo := repository.NewDevice(db)

	roomService := service.NewRoom(roomRepo)
	deviceService := service.NewDevice(deviceRepo)

	roomHandler := handler.NewRoom(roomService)
	deviceHandler := handler.NewDevice(deviceService)

	router := httpPackage.NewRouter(roomHandler, deviceHandler)

	httpPackage.SetupAndServe(cfg, db, router)
}
