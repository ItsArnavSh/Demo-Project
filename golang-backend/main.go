package main

import (
	"context"
	"database/sql"
	"demo/application/api"
	"demo/config"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// Load config
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Initialize logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("cannot initialize logger:", err)
	}
	defer logger.Sync()

	gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()

	conn, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/hospital?sslmode=disable")
	if err != nil {
		logger.Fatal("Could not connect: ", zap.Error(err))
	}
	// Initialize server
	server := api.NewServer(ctx, logger, conn)

	logger.Info("Starting server", zap.String("address", conf.ServerAddress))

	// Start server
	if err := server.Start(conf.ServerAddress); err != nil {
		logger.Fatal("cannot start server", zap.Error(err))
	}
}
