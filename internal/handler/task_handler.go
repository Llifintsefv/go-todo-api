package handler

import (
	"go-todo-api/internal/models"
	"go-todo-api/internal/service"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler interface {
	CreateTask(c *fiber.Ctx) error
}

type taskHandler struct {
	service service.TaskService
	logger  *slog.Logger
}

func NewTaskHandler(service service.TaskService, logger *slog.Logger) TaskHandler {
	return &taskHandler{
		service: service,
		logger:  logger,
	}
}

func (h *taskHandler) CreateTask(c *fiber.Ctx) error {
	CreateRequest := models.Task{}
	ctx := c.Context()
	err := c.BodyParser(&CreateRequest)
	if err != nil {
		h.logger.Error("Error parsing request body", slog.Any("error", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request body"})
	}

	if CreateRequest.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing title"})
	}

	TaskResponse, err := h.service.CreateTask(ctx, CreateRequest)
	if err != nil {
		h.logger.Error("Error", slog.Any("error", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(fiber.StatusCreated).JSON(TaskResponse)

}
