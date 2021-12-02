package service

import (
	"github.com/galentuo/goframe"
	"github.com/galentuo/goframe/examples/simple/pkg/store"
)

type UserInterface interface {
	Get(goframe.Context, int64) (*store.User, error)
	Insert(goframe.Context, store.User) error
}
