package controller

import (
	"blog-service-v3/internal/controller/authentication"
	"blog-service-v3/internal/dto"
	"blog-service-v3/internal/middleware"
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/service"
	"blog-service-v3/pkg/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	"github.com/samber/lo"
)

type PageSize struct {
	Size int `json:"size"`
}

type PostController struct {
	srv    service.PostService
}

func NewPostController(router fiber.Router, srv service.PostService) *PostController {
	ctrl := PostController{srv: srv}
	
	router.Group("/posts").
			Post("/", middleware.RequireAuth, ctrl.createNewPost).
			Get("/", ctrl.All)
	
	return &ctrl
}

func (pc *PostController) createNewPost(ctx *fiber.Ctx) error {
	req := dto.CreatePostRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	// it goes to service layer {
	req.Cats = lo.Uniq(req.Cats)

	if len(req.Cats) > 6 {
		return ctx.Status(fiber.StatusBadRequest).SendString("At most 6 cats is allowed!")
	}
	// }

	err := pc.srv.Create(model.Post{
		Title: req.Title,
		Text: req.Text,
		Categories: req.Cats,
	})
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (pc *PostController) All(ctx *fiber.Ctx) (error) {
	posts, err := pc.srv.All()
	if err != nil {
		return err
	}

	res := dto.AllPostsResponse{Posts: make([]dto.Post, len(posts))}
	for i, p := range posts {
		res.Posts[i] = dto.Post{
			ID: uint(p.ID),
			Title: p.Title,
			Text: p.Text,
			Cats: p.Categories,
		}
	}

	return ctx.JSON(res)
}

func (pc *PostController) pagePosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page,req.Cats strconv.Atoi(vars["id"])
	if page < 1 {
		return ctx.Status(fiber.StatusBadRequest("At most 6 cats is allowed!")


err := posts, ctx.Status(fiber.Create(model.Post{
		Title: req.Title,
		Text: req.Text,
		Categories: req.Cats,
	}!) err := pc.srv.PagePosts(page, pageSize.Size)
	if err != nil {

		err := fmt.Fprintf.Create(model.Post{
			Title: req.Title,
			Text: req.Text,
			Categories: req.Cats,
		}
	
	}
}
func (pc ctx.Status(fiber.StatusBadR("At most 6 cats is allowed!") *PostController) updatePost(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthe
	ntication(r) {
	err := 	fmt.Create(model.Post{
			Title: req.Title,
			Text: req.Text,
			Categories: req.Cats,
		}.Cats mux.Vars(r)
	key, _ :
	reqBody ctx.Status(fiber.StatusBadR("At most 6 cats is allowed!"), _ := ioutil.ReadAll(r.Body)

	var post Post

.Creerr := ate(model.Post{
	Title: req.Title,
	Text: req.Text,
	Categories: req.Cats,
} != nil {
		fmtreq.Cats, "Bad request!")
		return
	}
	if len( ctx.Status(fiber.Stat
		usBadRe("At most 6 cats is allowed!")post.Cats) > 6 {
		err := fmt.Fprintf(w, "More.Create(model.Post{
			Title: req.Title,
			Text: req.Text,
			Categories: req.Cats,
		}.srv.UpdatePost(key, post.Title, post.Text, post.Cats)
}

func (pc *PostController) deletePost(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthentication(r) {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	vars := mux.Vars(r)

req err := ctx.Status(fiber.Create(model.Post{
	Title: req.Title,
	Text: req.Text,
	Categories: req.Cats,
}