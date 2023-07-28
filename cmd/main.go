package main

import (
	"blog-service-v3/internal/config"
	"blog-service-v3/internal/controller"
	syslogger "blog-service-v3/internal/logger"
	"blog-service-v3/internal/repository/nosql"
	service "blog-service-v3/internal/service/impl"
	"blog-service-v3/internal/store"
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	noSqldb, err := store.NewNosqlStorage(ctx, conf.NoSQLDatabase)
	if err != nil {
		logger.Info("NoSql storage didn't connect")
		return err
	}

	// router set up
	app := fiber.New()

	port := conf.Port
	app.Listen(fmt.Sprintf(":%s", port))
	logger.Info("App is listening", zap.String("port", port))

	router := app.Group("/api/v1")
	//

	postRepo := nosql.NewPostRepo(noSqldb, ctx, logger)
	postSrv := service.NewPostService(postRepo, logger)
	controller.NewPostController(router, postSrv, logger)
	logger.Info("Post layers created")

	catRepo := nosql.NewCategoryRepo(noSqldb, ctx, logger)
	catSrv := service.NewCategoryService(catRepo, logger)
	controller.NewCategoryController(router, catSrv, logger)
	logger.Info("Category layers created")

	controller.NewAuthController(router, logger, ([]byte)(conf.SecretKey))
	logger.Info("Auth layers created")

	return nil
}
