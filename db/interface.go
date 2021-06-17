package db

import (
	"context"
	"errors"
)

type User struct {
	ID    string
	Name  string
	Email string
}

var ErrNotFound = errors.New("not found")

type Database interface {
	User(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	UpdateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]*User, error)
}
