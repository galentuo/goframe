package service

import (
	"encoding/json"
	"strconv"

	frame "github.com/galentuo/goframe"
	"github.com/galentuo/goframe/examples/simple/pkg/store"
)

type UserService struct {
	Service frame.HTTPService
	store   UserInterface
}

var (
	ErrUserNotFound = frame.NewInternalError(404, "user-001", "user not found")
)

func NewUserService() *UserService {
	srv := frame.NewHTTPService("user")
	srv.Use(m1)
	us := UserService{
		Service: srv,
		store:   store.NewUserStore(),
	}
	g := srv.NewGroup()
	g.Use(m2)
	srv.Route("/{userID:[0-9]+}", "GET", us.GetUser)
	g.Route("/{userID:[0-9]+}", "PUT", us.PutUser)
	return &us
}

func (us *UserService) GetUser(c frame.ServerContext) error {
	_userID := c.Param("userID")
	userID, _ := strconv.ParseInt(_userID, 10, 64)
	user, err := us.store.Get(c, userID)
	if err != nil {
		return c.Response().ErrorJSON(ErrUserNotFound.MoreDetailed("userID:", _userID))
	}
	return c.Response().SuccessJSON(200, user, "")
}

func (us *UserService) PutUser(c frame.ServerContext) error {
	_userID := c.Param("userID")
	userID, _ := strconv.ParseInt(_userID, 10, 64)
	type userReq struct {
		Name  *string `json:"name"`
		Email *string `json:"email"`
	}
	var u userReq
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&u); err != nil {
		return c.Response().ErrorJSON(
			frame.NewInternalError(500, "simple-001", "invalid request"))
	}
	err := us.store.Insert(c, store.User{
		ID:      userID,
		Name:    u.Name,
		Email:   u.Email,
		Enabled: true,
	})
	if err != nil {
		return c.Response().ErrorJSON(
			frame.NewInternalError(500, "user-002", "failed to insert user"))
	}
	return c.Response().SuccessJSON(201, nil, "inserted user")
}
