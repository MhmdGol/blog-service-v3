package controller

import (
	"blog-service-v3/internal/controller/authentication"
	"blog-service-v3/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	h.router.HandleFunc("/post/read", h.AllPosts).Methods("GET")
	h.router.HandleFunc("/post/read/{id}", h.PagePosts).Methods("GET")
	h.router.HandleFunc("/post/update/{id}", h.updatePost).Methods("PUT")
	h.router.HandleFunc("/post/delete/{id}", h.deletePost).Methods("DELETE")
}

func (h *categoryApi) SetupCategoryRoutes() {
	h.router.HandleFunc("/category/create", h.createNewCategory).Methods("POST")
	h.router.HandleFunc("/category/read", h.AllCategories).Methods("GET")
	h.router.HandleFunc("/category/update/{id}", h.updateCategory).Methods("PUT")
	h.router.HandleFunc("/category/delete/{id}", h.deleteCategory).Methods("DELETE")
}

func (h *loginApi) SetupLoginRoutes() {
	h.router.HandleFunc("/login", authentication.LoginHandler).Methods("POST")
}

func (h *postApi) Start() {
	// the port must be replaced with config
	log.Fatal(http.ListenAndServe(":8080", h.router))
}

func (h *categoryApi) Start() {
	// the port must be replaced with config
	log.Fatal(http.ListenAndServe(":8080", h.router))
}

func (h *loginApi) Start() {
	// the port must be replaced with config
	log.Fatal(http.ListenAndServe(":8080", h.router))
}
