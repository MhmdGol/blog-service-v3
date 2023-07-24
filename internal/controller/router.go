package controller

import (
	"blog-service-v3/internal/controller/authentication"
	"blog-service-v3/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type postApi struct {
	router *mux.Router
	app    service.PostApp
}

type categoryApi struct {
	router *mux.Router
	app    service.CategoryApp
}

type loginApi struct {
	router *mux.Router
}

func PostNew(app service.PostApp) *postApi {
	a := &postApi{
		router: mux.NewRouter().StrictSlash(true),
		app:    app,
	}

	a.SetupPostRoutes()
	return a
}

func CategoryNew(app service.CategoryApp) *categoryApi {
	a := &categoryApi{
		router: mux.NewRouter().StrictSlash(true),
		app:    app,
	}

	a.SetupCategoryRoutes()
	return a
}

func LoginNew() *loginApi {
	a := &loginApi{
		router: mux.NewRouter().StrictSlash(true),
	}

	a.SetupLoginRoutes()
	return a
}

func (h *postApi) SetupPostRoutes() {
	h.router.HandleFunc("/post/create", h.createNewPost).Methods("POST")
	h.router.HandleFunc("/post/read", h.allPosts).Methods("GET")
	h.router.HandleFunc("/post/read/{id}", h.pagePosts).Methods("GET")
	h.router.HandleFunc("/post/update/{id}", h.updatePost).Methods("PUT")
	h.router.HandleFunc("/post/delete/{id}", h.deletePost).Methods("DELETE")
}

func (h *categoryApi) SetupCategoryRoutes() {
	h.router.HandleFunc("/category/create", h.createNewCategory).Methods("POST")
	h.router.HandleFunc("/category/read", h.allCategories).Methods("GET")
	h.router.HandleFunc("/category/update/{id}", h.updateCategory).Methods("PUT")
	h.router.HandleFunc("/category/delete/{id}", h.deleteCategory).Methods("DELETE")
}

func (h *loginApi) SetupLoginRoutes() {
	h.router.HandleFunc("/login", authentication.LoginHandler).Methods("POST")
}

func (h *postApi) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("postApiPort")), h.router))
}

func (h *categoryApi) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("categoryApiPort")), h.router))
}

func (h *loginApi) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("loginApiPort")), h.router))
}
