package main

import (
	"blog-service-v3/internal/config"
	"blog-service-v3/internal/controller"
	"blog-service-v3/internal/logger"
	"blog-service-v3/internal/repository/nosql"
	service "blog-service-v3/internal/service/impl"
	"blog-service-v3/internal/store"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := logger.InitLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()
	logger.Info("Logger initialized")

	conf, err := config.Load()
	if err != nil {
		logger.Info("Config didnt load")
		return err
	}

	db, err := store.NewNosqlStorage(conf.NoSQLDatabase)
	if err != nil {
		logger.Info("NoSql storage didn't connect")
		return err
	}

	// db, err := store.NewSqlStorage(conf.SQLDatabase)
	// if err != nil {
	// 	logger.Info("Sql storage didn't connect")
	// 	return err
	// }

	app := fiber.New()
	port := conf.Port
	router := app.Group("/api/v1")

	postRepo := nosql.NewPostRepo(db, logger)
	postSrv := service.NewPostService(postRepo, logger)
	controller.NewPostController(router, postSrv, logger)
	logger.Info("Post layers created")

	catRepo := nosql.NewCategoryRepo(db, logger)
	catSrv := service.NewCategoryService(catRepo, logger)
	controller.NewCategoryController(router, catSrv, logger)
	logger.Info("Category layers created")

	controller.NewAuthController(router, logger, ([]byte)(conf.SecretKey))
	logger.Info("Auth layers created")

	app.Listen(fmt.Sprintf(":%s", port))
	logger.Info("App is listening", zap.String("port", port))

	return nil
}
