package memdb

import (
	"bank/db"
	"context"
	"errors"

	"github.com/google/uuid"
)

type Database struct {
	users map[string]*db.User
}

func NewDatabase() *Database {
	return &Database{
		users: make(map[string]*db.User),
	}
}

var ErrNotImplemented = errors.New("not implemnted")

func (d *Database) User(ctx context.Context, id string) (*db.User, error) {
	u, ok := d.users[id]
	if !ok {
		return nil, db.ErrNotFound
	}
	return u, nil
}

/*
{"Name":"x","Email":"Y"}
*/
func (d *Database) CreateUser(ctx context.Context, u *db.User) (*db.User, error) {
	u.ID = uuid.New().String()
	d.users[u.ID] = u
	return u, nil
}

/*
{"Name":"x1","Email":"Y1"}
*/
func (d *Database) UpdateUser(ctx context.Context, u *db.User) error {
	_, ok := d.users[u.ID]
	if !ok {
		return db.ErrNotFound
	}
	d.users[u.ID] = u
	return nil
}

func (d *Database) DeleteUser(ctx context.Context, id string) error {
	_, ok := d.users[id]
	if !ok {
		return db.ErrNotFound
	}
	delete(d.users, id)
	return nil
}

func (d *Database) ListUsers(ctx context.Context) ([]*db.User, error) {
	users := make([]*db.User, 0, len(d.users))
	for _, user := range d.users {
		users = append(users, user)
	}
	return users, nil
}
