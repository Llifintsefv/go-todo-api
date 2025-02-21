package handler

import (
	"go-todo-api/internal/models"
	"go-todo-api/internal/service"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler interface {
	CreateTask(c *fiber.Ctx) error
	GetTasks(c *fiber.Ctx) error
	GetTask(c *fiber.Ctx) error
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
	createRequest := models.Task{}
	ctx := c.Context()
	err := c.BodyParser(&createRequest)
	if err != nil {
		h.logger.Error("Error parsing request body", slog.Any("error", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request body"})
	}

	if createRequest.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	}

	taskResponse, err := h.service.CreateTask(ctx, createRequest)
	if err != nil {
		h.logger.Error("Error", slog.Any("error", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(fiber.StatusCreated).JSON(taskResponse)

}

func (h *taskHandler) GetTasks(c *fiber.Ctx) error {
	ctx := c.Context()

	var tasks []models.Task
	tasks, err := h.service.GetTasks(ctx)
	if err != nil {
		h.logger.Error("Error", slog.Any("error", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}

func (h *taskHandler) GetTask(c *fiber.Ctx) error {
	ctx := c.Context()

	taskIdStr := c.Params("taskId")

	if taskIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "taskId is required"})
	}

	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	task, err := h.service.GetTask(ctx, taskId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	return c.Status(fiber.StatusOK).JSON(task)

}
