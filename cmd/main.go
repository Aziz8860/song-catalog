package main

import (
	"log"

	"github.com/Aziz8860/song-catalog/internal/configs"
	membershipsHandler "github.com/Aziz8860/song-catalog/internal/handler/memberships"
	"github.com/Aziz8860/song-catalog/internal/models/memberships"
	membershipsRepo "github.com/Aziz8860/song-catalog/internal/repository/memberships"
	membershipsSvc "github.com/Aziz8860/song-catalog/internal/service/memberships"
	"github.com/Aziz8860/song-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/", // for local configs file path
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize configs: %v", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v", err)
	}
	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
