package controller

import (
	"blog-service-v3/internal/dto"
	"blog-service-v3/internal/middleware"
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CategoryController struct {
	srv    service.CategoryService
	logger *zap.Logger
	vld    *validator.Validate
}

func NewCategoryController(router fiber.Router, srv service.CategoryService, logger *zap.Logger) *CategoryController {
	ctrl := CategoryController{
		srv:    srv,
		logger: logger,
		vld:    validator.New(),
	}

	router.Group("/category").
		Post("/", middleware.RequireAuth, ctrl.CreateNewCategory).
		Get("/", ctrl.All).
		Put("/:id", middleware.RequireAuth, ctrl.UpdateByID).
		Delete("/:id", middleware.RequireAuth, ctrl.DeleteByID)

	return &ctrl
}

func (cc *CategoryController) CreateNewCategory(ctx *fiber.Ctx) error {
	cc.logger.Info("controller, category, CreateNewCategory")
	req := dto.Category{}
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := cc.vld.Struct(req)
	if err != nil {
		return err
	}

	err = cc.srv.Create(model.Category{
		Name: req.Name,
	})
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (cc *CategoryController) All(ctx *fiber.Ctx) error {
	cc.logger.Info("controller, category, All")
	categories, err := cc.srv.All()
	if err != nil {
		return err
	}

	res := make([]dto.Category, len(categories))
	for i, c := range categories {
		res[i] = dto.Category{
			ID:   string(c.ID),
			Name: c.Name,
		}
	}

	return ctx.JSON(res)
}

func (cc *CategoryController) UpdateByID(ctx *fiber.Ctx) error {
	cc.logger.Info("controller, category, UpdateByID")
	req := dto.Category{}
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := cc.vld.Struct(req)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	err = cc.srv.UpdateByID(model.Category{
		ID:   (model.ID)(id),
		Name: req.Name,
	})
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (cc *CategoryController) DeleteByID(ctx *fiber.Ctx) error {
	cc.logger.Info("controller, category, DeleteByID")
	id, _ := strconv.Atoi(ctx.Params("id"))
	idToDelete := (model.ID)(id)

	err := cc.srv.DeleteByID(idToDelete)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}
