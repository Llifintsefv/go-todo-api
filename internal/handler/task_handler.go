package handler

import (
	"go-todo-api/internal/models"
	"go-todo-api/internal/service"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler interface {
	CreateTask(c *fiber.Ctx) error
	GetTasks(c *fiber.Ctx) error
	UpdateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
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

func (h *taskHandler) UpdateTask(c *fiber.Ctx) error {
	ctx := c.Context()
	taskId, err := c.ParamsInt("taskId")
	if err != nil {
		h.logger.Error("Invalid task ID", slog.Any("error", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}
	updateRequest := models.Task{}
	if err := c.BodyParser(&updateRequest); err != nil {
		h.logger.Error("Error parsing request body", slog.Any("error", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request body"})
	}
	updateRequest.ID = taskId
	task, err := h.service.UpdateTask(ctx, updateRequest)
	if err == models.ErrTaskNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	} else if err != nil {
		h.logger.Error("Error updating task", slog.Any("error", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(task)
}

func (h *taskHandler) DeleteTask(c *fiber.Ctx) error {
	ctx := c.Context()
	taskId, err := c.ParamsInt("taskId")
	if err != nil {
		h.logger.Error("Invalid task ID", slog.Any("error", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}
	if err := h.service.DeleteTask(ctx, taskId); err != nil {
		if err == models.ErrTaskNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		h.logger.Error("Error deleting task", slog.Any("error", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Task deleted successfully"})
}

func isValidStatus(status string) bool {
	switch status {
	case models.StatusNew, models.StatusInProgress, models.StatusDone:
		return true
	default:
		return false
	}
}
