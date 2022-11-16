package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/snykk/fiber-mongo-crud/datatransfer"
	"github.com/snykk/fiber-mongo-crud/service"
	"github.com/snykk/fiber-mongo-crud/utils"
)

type TodoController struct {
	TodoService service.TodoService
}

func NewTodoController(todoService *service.TodoService) TodoController {
	return TodoController{TodoService: *todoService}
}

func (controller *TodoController) Route(app *fiber.App) {
	app.Get("/api/todos", controller.List)
	app.Get("/api/todos/:id", controller.GetById)
	app.Post("/api/todos", controller.Create)
	app.Put("/api/todos/:id", controller.UpdateTodo)
	app.Delete("/api/todos", controller.DeleteAll)
	app.Delete("/api/todos/:id", controller.DeleteById)
}

func (controller *TodoController) Create(c *fiber.Ctx) error {
	var request datatransfer.TodoRequest
	err := c.BodyParser(&request)
	if err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if isValid, err := utils.IsRequestValid(request); !isValid {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())

	}

	if err := utils.IsPriorityValid(request.Priority); err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	request.Id = uuid.New().String()
	response := controller.TodoService.Create(request)
	return NewSuccessResponse(c, "todos data created successfully", response)
}

func (controller *TodoController) List(c *fiber.Ctx) error {
	responses, err := controller.TodoService.List()
	if err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return NewSuccessResponse(c, "todos data fetched successfully", responses)
}

func (controller *TodoController) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	responses, err := controller.TodoService.GetById(id)

	if err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return NewSuccessResponse(c, "todos data fetched successfully", responses)
}

func (controller *TodoController) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var request datatransfer.TodoUpdateRequest
	err := c.BodyParser(&request)
	if err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := utils.IsPriorityValid(request.Priority); err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	response, err := controller.TodoService.UpdateTodo(id, request)
	if err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return NewSuccessResponse(c, fmt.Sprintf("todos data with id %s updated successfully", response.Id), response)
}

func (controller *TodoController) DeleteAll(c *fiber.Ctx) error {
	err := controller.TodoService.DeleteAll()
	if err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return NewSuccessResponse(c, "todos cleared successfully", nil)
}

func (controller *TodoController) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	err := controller.TodoService.DeleteById(id)
	if err != nil {
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return NewSuccessResponse(c, fmt.Sprintf("todos with id %s deleted successfully", id), nil)
}
