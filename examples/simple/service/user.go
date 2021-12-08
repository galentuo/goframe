package service

import (
	"encoding/json"
	"strconv"

	frame "github.com/galentuo/goframe"
	"github.com/galentuo/goframe/examples/simple/pkg/store"
	"github.com/gorilla/mux"
)

type UserService struct {
	*frame.HTTPServer
	store UserInterface
}

func NewUserService() *UserService {
	srv := UserService{
		HTTPServer: frame.NewHTTPServer("user"),
		store:      store.NewUserStore(),
	}
	srv.Route("/{userID:[0-9]+}", "GET", srv.GetUser)
	srv.Route("/{userID:[0-9]+}", "PUT", srv.PutUser)
	return &srv
}

func (us *UserService) GetUser(c frame.ServerContext) error {
	vars := mux.Vars(c.Request())
	_userID := vars["userID"]
	userID, _ := strconv.ParseInt(_userID, 10, 64)
	user, err := us.store.Get(c, userID)
	if err != nil {
		return c.Response().ErrorJSON(frame.NewInternalError(404, "user-001", "user not found"))
	}
	return c.Response().SuccessJSON(200, user, "")
}

func (us *UserService) PutUser(c frame.ServerContext) error {
	vars := mux.Vars(c.Request())
	_userID := vars["userID"]
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
