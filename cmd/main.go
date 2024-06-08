package main

import (
	"github.com/elyarsadig/smart-home-iot/config"
	"github.com/elyarsadig/smart-home-iot/internal/infrastructure/db/sqlite"
	httpPackage "github.com/elyarsadig/smart-home-iot/internal/infrastructure/http"
)

func main() {
	cfg := config.LoadConfig()
	
	db := sqlite.InitDB(&cfg.Database)

	httpPackage.SetupAndServe(cfg, db)
}
