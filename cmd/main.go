package main

import (
	"blog-service-v3/internal/config"
	"blog-service-v3/internal/controller"
	syslogger "blog-service-v3/internal/logger"
	"blog-service-v3/internal/repository/sql"
	"blog-service-v3/internal/repository/sql/dbmodel"
	service "blog-service-v3/internal/service/impl"
	"blog-service-v3/internal/store"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	syslogger.InitLogger(logger)
	defer logger.Sync()

	conf, err := config.Load()
	if err != nil {
		logger.Info("Config didnt load")
		return err
	}

	db, err := store.New(conf.Database)
	if err != nil {
		logger.Info("Store didn't connect")
		return err
	}

	err = dbmodel.Init(db)
	if err != nil {
		logger.Info("Database couldn't initialize")
		return err
	}

	app := fiber.New()
	app.Listen(fmt.Sprintf(":%s", conf.HttpPort))
	logger.Info("App is listening", zap.String("port", conf.HttpPort))

	router := app.Group("/api/v1")

	postRepo := sql.NewPostRopo(db, logger)
	postSrv := service.NewPostService(postRepo, logger)
	controller.NewPostController(router, postSrv)
	logger.Info("Post layers created")

	catRepo := sql.NewCategoryRepo(db, logger)
	catSrv := service.NewCategoryService(catRepo, logger)
	controller.NewCategoryController(router, catSrv)
	logger.Info("Category layers created")

	controller.NewAuthController(router, ([]byte)(conf.SecretKey))
	logger.Info("Auth layers created")

	return nil
}
