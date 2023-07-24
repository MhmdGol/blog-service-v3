package main

import (
	"blog-service-v3/internal/config"
	"blog-service-v3/internal/controller"
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository/sql"
	"blog-service-v3/internal/service"
	"blog-service-v3/internal/store"
	"log"
	"os"
)

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func run() error {
	conf, err := config.Load()
	if err != nil {
		return err
	}

	db, err := store.New(conf.Database)
	if err != nil {
		return err
	}

	err = model.Init(db)
	if err != nil {
		return err
	}

	postRepo := sql.NewPostRopo(db)
	postSrv := service.PostNew(postRepo)
	postCtrl := controller.PostNew(postSrv)

	postCtrl.Start()

	catRepo := sql.NewCategoryRepo(db)
	catSrv := service.CategoryNew(catRepo)
	catCtrl := controller.CategoryNew(catSrv)

	catCtrl.Start()

	loginCtrl := controller.LoginNew()

	loginCtrl.Start()

	return nil
}

// log, err := logger.NewCustomLogger()
// if err != nil {
//     panic("failed to initialize logger: " + err.Error())
// }

// defer log.Close()
