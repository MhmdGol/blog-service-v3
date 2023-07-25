package controller

import (
	"blog-service-v3/internal/dto"
	"blog-service-v3/internal/middleware"
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type PostController struct {
	srv service.PostService
}

func NewPostController(router fiber.Router, srv service.PostService) *PostController {
	ctrl := PostController{srv: srv}

	router.Group("/posts").
		Post("/", middleware.RequireAuth, ctrl.CreateNewPost).
		Get("/", ctrl.All).
		Get("/:page", ctrl.Paginated).
		Put("/", middleware.RequireAuth, ctrl.UpdateByID).
		Delete("/:id", middleware.RequireAuth, ctrl.DeleteByID)

	return &ctrl
}

func (pc *PostController) CreateNewPost(ctx *fiber.Ctx) error {
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
		Title:      req.Title,
		Text:       req.Text,
		Categories: req.Cats,
	})
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (pc *PostController) All(ctx *fiber.Ctx) error {
	posts, err := pc.srv.All()
	if err != nil {
		return err
	}

	res := dto.AllPostsResponse{Posts: make([]dto.Post, len(posts))}
	for i, p := range posts {
		res.Posts[i] = dto.Post{
			ID:    uint(p.ID),
			Title: p.Title,
			Text:  p.Text,
			Cats:  p.Categories,
		}
	}

	return ctx.JSON(res)
}

func (pc *PostController) Paginated(ctx *fiber.Ctx) error {
	size := dto.PageSize{}
	if err := ctx.BodyParser(&size); err != nil {
		return err
	}

	pgNum, _ := strconv.Atoi(ctx.Params("page"))
	posts, err := pc.srv.Paginated(pgNum, size.Size)
	if err != nil {
		return err
	}

	res := dto.AllPostsResponse{Posts: make([]dto.Post, len(posts))}
	for i, p := range posts {
		res.Posts[i] = dto.Post{
			ID:    uint(p.ID),
			Title: p.Title,
			Text:  p.Text,
			Cats:  p.Categories,
		}
	}

	return ctx.JSON(res)
}

func (pc *PostController) UpdateByID(ctx *fiber.Ctx) error {
	req := dto.Post{}
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	// it goes to service layer {
	req.Cats = lo.Uniq(req.Cats)

	if len(req.Cats) > 6 {
		return ctx.Status(fiber.StatusBadRequest).SendString("At most 6 cats is allowed!")
	}
	// }

	err := pc.srv.UpdateByID(model.Post{
		ID:         model.ID(req.ID),
		Title:      req.Title,
		Text:       req.Text,
		Categories: req.Cats,
	})

	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (pc *PostController) DeleteByID(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	idToDelete := (model.ID)(id)

	err := pc.srv.DeleteByID(idToDelete)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}
