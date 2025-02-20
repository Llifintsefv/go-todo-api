package router

import (
	"go-todo-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(taskHandler handler.TaskHandler) *fiber.App {
	app := fiber.New()

	app.Post("/tasks",taskHandler.CreateTask)

	return app
}
