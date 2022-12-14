package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/uraldav/gqlgen-todos/graph/generated"
	"github.com/uraldav/gqlgen-todos/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Title:     input.Title,
		Text:      input.Text,
		ID:        fmt.Sprintf("T%d", rand.Int()),
		User:      &model.User{ID: input.UserID, Name: "user " + input.UserID},
		UserID:    input.UserID,
		CreatedAt: time.Now(),
	}
	r.todos = append(r.todos, todo)

	return todo, nil
}

// SwitchDone is the resolver for the switchDone field.
func (r *mutationResolver) SwitchDone(ctx context.Context, id string) (*model.Todo, error) {
	for i := range r.todos {
		if r.todos[i].ID == id {
			r.todos[i].Done = !r.todos[i].Done
			return r.todos[i], nil
		}
	}

	return nil, gqlerror.Errorf("not found")
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
