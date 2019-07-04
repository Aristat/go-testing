package resolver

import (
	"context"

	"github.com/Aristat/go-testing"
	graphql1 "github.com/Aristat/go-testing/generated/graphql"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() graphql1.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graphql1.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() graphql1.TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input graphql1.NewTodo) (*main.Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*main.Todo, error) {
	panic("not implemented")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *main.Todo) (*graphql1.User, error) {
	panic("not implemented")
}
