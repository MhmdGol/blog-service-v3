package main

import (
	"blog-service-v3/internal/config"
	"blog-service-v3/internal/controller"
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
	initLogger()
	defer logger.Sync()

	conf, err := config.Load()
	if err != nil {
		return err
	}

	db, err := store.New(conf.Database)
	if err != nil {
		return err
	}

	err = dbmodel.Init(db)
	if err != nil {
		return err
	}

	router := fiber.New()

	router.Listen(fmt.Sprintf(":%s", conf.HttpPort))

	postRepo := sql.NewPostRopo(db, logger)
	postSrv := service.NewPostService(postRepo, logger)
	_ = controller.NewPostController(router, postSrv)

	catRepo := sql.NewCategoryRepo(db, logger)
	catSrv := service.NewCategoryService(catRepo, logger)
	_ = controller.NewCategoryController(router, catSrv)

	_ = controller.NewAuthController(router, ([]byte)(conf.SecretKey))

	return nil
}

func initLogger() {
	cfg := zap.NewDevelopmentConfig()

	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}
