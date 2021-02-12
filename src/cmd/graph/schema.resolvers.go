package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/matthausen/gql-example/cmd/graph/generated"
	"github.com/matthausen/gql-example/cmd/graph/model"
	"github.com/matthausen/gql-example/pkg/user"
)

func (r *myMutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	id, err := r.User.Create(user.Name, *user.IsPremium)
	if err != nil {
		return nil, err
	} else {
		log.Printf("To Do saved with identifier: %s", *id)
		return &model.User{
			ID:        *id,
			Name:      user.Name,
			IsPremium: *user.IsPremium,
		}, nil
	}

}

func (r *myMutationResolver) UpdateUser(ctx context.Context, id string, updatedUser *model.UserInput) (*model.User, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid identifier: %s", id)
	}
	err = r.User.Update(id, updatedUser.Name, *updatedUser.IsPremium)
	if err != nil {
		return nil, err
	} else {
		log.Printf("To Do with identifier: %s updated", id)
		return &model.User{
			ID:        id,
			Name:      updatedUser.Name,
			IsPremium: *updatedUser.IsPremium,
		}, nil
	}

}

func (r *myQueryResolver) User(ctx context.Context, id string) (*model.User, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid identifier: %s", id)
	}
	item, err := r.User.Get(id)
	if err != nil {
		return nil, err
	} else {
		return &model.User{
			ID:        item.Id,
			Name:      item.Name,
			IsPremium: item.IsPremium,
		}, nil
	}

}

func (r *myQueryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var items []*model.User
	var savedItems []user.UserItem
	savedItems, err := r.User.List()
	if err != nil {
		return nil, err
	}
	for i, savedItem := range savedItems {
		var item model.User
		savedItem = savedItems[i]
		item.ID = savedItem.Id
		item.Name = savedItem.Name
		item.IsPremium = savedItem.IsPremium
		items = append(items, &item)
	}
	return items, nil

}

// MyMutation returns generated.MyMutationResolver implementation.
func (r *Resolver) MyMutation() generated.MyMutationResolver { return &myMutationResolver{r} }

// MyQuery returns generated.MyQueryResolver implementation.
func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myMutationResolver struct{ *Resolver }
type myQueryResolver struct{ *Resolver }
