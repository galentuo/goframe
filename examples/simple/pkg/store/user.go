package store

import (
	"errors"
	"fmt"
	"sync"

	"github.com/galentuo/goframe"
)

type UserStore struct {
	inner *sync.Map
}

type User struct {
	ID      int64   `json:"id"`
	Name    *string `json:"name"`
	Email   *string `json:"email"`
	Enabled bool    `json:"enabled"`
}

func NewUserStore() *UserStore {
	return &UserStore{
		inner: &sync.Map{},
	}
}

func (store UserStore) Get(c goframe.Context, id int64) (*User, error) {
	c.Logger().Debug(fmt.Sprintf("fetching user %d from map", id))
	u, ok := store.inner.Load(id)
	if !ok {
		return nil, errors.New("User not found")
	}
	usr := u.(User)
	c.Logger().Debug(fmt.Sprintf("got user %d :: %+v", id, usr))
	return &usr, nil
}

func (store *UserStore) Insert(c goframe.Context, user User) error {
	c.Logger().Debug(fmt.Sprintf("inserting user %+v", user))
	store.inner.Store(user.ID, user)
	return nil
}
