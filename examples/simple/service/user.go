package service

import (
	"encoding/json"
	"strconv"

	frame "github.com/galento/goframe"
	"github.com/galento/goframe/examples/simple/pkg/store"
	"github.com/galentuo/goframe"
	"github.com/gorilla/mux"
)

type UserService struct {
	*goframe.HTTPServer
	store UserInterface
}

func NewUserService() *UserService {
	srv := UserService{
		HTTPServer: goframe.NewHTTPService("user"),
		store:      store.NewUserStore(),
	}
	srv.Route("/{userID:[0-9]+}", "GET", srv.GetUser)
	srv.Route("/{userID:[0-9]+}", "PUT", srv.PutUser)
	return &srv
}

func (us *UserService) GetUser(c frame.APIContext) error {
	vars := mux.Vars(c.Request())
	_userID := vars["userID"]
	userID, _ := strconv.ParseInt(_userID, 10, 64)
	user, err := us.store.Get(&c, userID)
	if err != nil {
		c.Response().JSON(404, frame.ERROR, nil, "user not found")
		return nil
	}
	c.Response().JSON(200, frame.OK, user, "")
	return nil
}

func (us *UserService) PutUser(c frame.APIContext) error {
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
		c.Response().JSON(500, frame.ERROR, nil, "invalid request")
		return nil
	}
	err := us.store.Insert(&c, store.User{
		ID:      userID,
		Name:    u.Name,
		Email:   u.Email,
		Enabled: true,
	})
	if err != nil {
		c.Response().JSON(500, frame.ERROR, nil, "failed to insert user")
		return nil
	}
	c.Response().JSON(201, frame.OK, nil, "inserted user")
	return nil
}
