package user

import "time"

type UserItem struct {
	Id        string
	Name      string
	IsPremium bool
	CreatedOn time.Time
	UpdatedOn *time.Time
}

type User interface {
	Initialise() error
	Create(name string, isPremium bool) (*string, error)
	Update(id string, name string, isPremium bool) error
	Get(id string) (*UserItem, error)
	List() ([]UserItem, error)
}
