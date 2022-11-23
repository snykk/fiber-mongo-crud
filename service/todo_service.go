package service

import (
	"github.com/snykk/fiber-mongo-crud/datatransfer"
	"github.com/snykk/fiber-mongo-crud/entity"
	"github.com/snykk/fiber-mongo-crud/repository"
)

type TodoService interface {
	Create(request datatransfer.TodoRequest) (response datatransfer.TodoResponse, err error)
	List() (responses []datatransfer.TodoResponse, err error)
	GetById(id string) (responses datatransfer.TodoResponse, err error)
	UpdateTodo(id string, request datatransfer.TodoUpdateRequest) (response datatransfer.TodoResponse, err error)
	DeleteAll() error
	DeleteById(id string) error
}

func NewTodoService(todoRepository *repository.TodoRepository) TodoService {
	return &todoService{
		TodoRepository: *todoRepository,
	}
}

type todoService struct {
	TodoRepository repository.TodoRepository
}

func (service *todoService) Create(request datatransfer.TodoRequest) (response datatransfer.TodoResponse, err error) {
	// validation.Validate(request)

	todo := entity.Todo{
		Id:       request.Id,
		Activity: request.Activity,
		Priority: request.Priority,
	}

	err = service.TodoRepository.Insert(todo)
	if err != nil {
		return response, err
	}

	response = datatransfer.TodoResponse{
		Id:       todo.Id,
		Activity: todo.Activity,
		Priority: todo.Priority,
		IsDone:   todo.IsDone,
	}
	return response, err
}

func (service *todoService) List() (responses []datatransfer.TodoResponse, err error) {
	todos, err := service.TodoRepository.FindAll()
	if err != nil {
		return responses, err
	}

	for _, todo := range todos {
		responses = append(responses, datatransfer.TodoResponse{
			Id:       todo.Id,
			Activity: todo.Activity,
			Priority: todo.Priority,
			IsDone:   todo.IsDone,
		})
	}

	return responses, err
}

func (service *todoService) GetById(id string) (response datatransfer.TodoResponse, err error) {
	todo, err := service.TodoRepository.GetById(id)
	if err != nil {
		return response, err
	}

	response = datatransfer.TodoResponse{
		Id:       todo.Id,
		Activity: todo.Activity,
		Priority: todo.Priority,
		IsDone:   todo.IsDone,
	}

	return response, err
}

func (service *todoService) UpdateTodo(id string, request datatransfer.TodoUpdateRequest) (response datatransfer.TodoResponse, err error) {
	data := structToMap(request)

	err = service.TodoRepository.UpdateById(id, data)
	if err != nil {
		return response, err
	}

	todo, err := service.TodoRepository.GetById(id)
	if err != nil {
		return response, err
	}

	response = datatransfer.TodoResponse{
		Id:       todo.Id,
		Activity: todo.Activity,
		Priority: todo.Priority,
		IsDone:   todo.IsDone,
	}

	return response, err
}

func (service *todoService) DeleteAll() error {
	err := service.TodoRepository.DeleteAll()
	return err
}

func (service *todoService) DeleteById(id string) error {
	err := service.TodoRepository.DeleteById(id)
	return err
}

func structToMap(request datatransfer.TodoUpdateRequest) map[string]interface{} {
	newMap := make(map[string]interface{})
	if request.Activity != "" {
		newMap["activity"] = request.Activity
	}
	if request.Priority != "" {
		newMap["priority"] = request.Priority
	}
	if request.IsDone != nil {
		newMap["is_done"] = request.IsDone
	}
	return newMap
}
