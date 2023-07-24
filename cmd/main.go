package main

import (
	"blog-service-v3/internal/configuration"
	"blog-service-v3/internal/controller"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/service"
	"blog-service-v3/internal/store"
)

func main() {
	//Setup viper
	configuration.SetConfigs()

	db := store.New()

	postRepo := repository.PostNew(db)
	postSrv := service.PostNew(postRepo)
	postCtrl := controller.PostNew(postSrv)

	postCtrl.Start()

	catRepo := repository.CategoryNew(db)
	catSrv := service.CategoryNew(catRepo)
	catCtrl := controller.CategoryNew(catSrv)

	catCtrl.Start()
}
