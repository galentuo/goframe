package service

import (
	frame "github.com/galentuo/goframe"
	"github.com/galentuo/goframe/examples/simple/pkg/store"
)

type UserInterface interface {
	Get(frame.Context, int64) (*store.User, error)
	Insert(frame.Context, store.User) error
}
