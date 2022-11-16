package repository

import (
	"github.com/snykk/fiber-mongo-crud/config"

	"github.com/snykk/fiber-mongo-crud/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	Insert(todo entity.Todo) (err error)
	FindAll() (todos []entity.Todo, err error)
	DeleteAll() (err error)
	GetById(id string) (todo entity.Todo, err error)
	UpdateById(id string, data map[string]interface{}) (err error)
	DeleteById(id string) (err error)
}

func NewTodoRepository(database *mongo.Database) TodoRepository {
	return &todoRepository{
		Collection: database.Collection("todos"),
	}
}

type todoRepository struct {
	Collection *mongo.Collection
}

func (repository *todoRepository) Insert(todo entity.Todo) (err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err = repository.Collection.InsertOne(ctx, bson.M{
		"_id":      todo.Id,
		"activity": todo.Activity,
		"priority": todo.Priority,
		"is_done":  todo.IsDone,
	})

	return
}

func (repository *todoRepository) GetById(id string) (todo entity.Todo, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	err = repository.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo)
	return todo, err
}

func (repository *todoRepository) UpdateById(id string, data map[string]interface{}) (err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err = repository.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": data})
	return
}

func (repository *todoRepository) FindAll() (todos []entity.Todo, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cur, err := repository.Collection.Find(ctx, bson.M{})
	if err != nil {
		return todos, err
	}

	var docs []bson.M
	err = cur.All(ctx, &docs)
	if err != nil {
		return todos, err
	}

	for _, doc := range docs {
		todos = append(todos, entity.Todo{
			Id:       doc["_id"].(string),
			Activity: doc["activity"].(string),
			Priority: doc["priority"].(string),
			IsDone:   doc["is_done"].(bool),
		})
	}

	return todos, err
}

func (repository *todoRepository) DeleteAll() (err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err = repository.Collection.DeleteMany(ctx, bson.M{})
	return
}

func (repository *todoRepository) DeleteById(id string) (err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err = repository.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return
}
