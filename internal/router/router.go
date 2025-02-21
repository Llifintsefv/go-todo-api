package router

import (
	"go-todo-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(taskHandler handler.TaskHandler) *fiber.App {
	app := fiber.New()

	app.Post("/tasks", taskHandler.CreateTask)
	app.Get("/tasks", taskHandler.GetTasks)
	app.Get("/tasks/:taskId", taskHandler.UpdateTask)
	app.Delete("/tasks/:taskId", taskHandler.DeleteTask)
	return app
}
