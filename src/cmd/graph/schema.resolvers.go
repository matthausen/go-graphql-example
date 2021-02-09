package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matthausen/gql-example/graph/generated"
	"github.com/matthausen/gql-example/graph/model"
)

func (r *myMutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	r.User.Create(user.Name, true)
	panic(fmt.Errorf("CreateUser method not implemented"))
}

func (r *myMutationResolver) UpdateUser(ctx context.Context, id string, changes map[string]interface{}) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) MyMutation() generated.MyMutationResolver { return &myMutationResolver{r} }

func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myMutationResolver struct{ *Resolver }
type myQueryResolver struct{ *Resolver }
