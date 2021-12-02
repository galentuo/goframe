package service

import (
	"encoding/json"
	"strconv"

	"github.com/galentuo/goframe"
	"github.com/galentuo/goframe/examples/simple/pkg/store"
	"github.com/gorilla/mux"
)

type UserService struct {
	store UserInterface
}

func NewUserService() *UserService {

	return &UserService{
		store: store.NewUserStore(),
	}
}

func (us UserService) Name() string {
	return "user-service"
}

func (us UserService) Prefix() string {
	return "/user"
}

func (us UserService) Middleware(ep goframe.JSONEndpoint) goframe.JSONEndpoint {
	return func(c goframe.APIContext) error {
		// if c.Request().Header.Get("Authorization") == "" {
		// 	c.Response().JSON(403, goframe.ERROR, nil, "You don't have permission to do this")
		// 	return nil
		// }
		return ep(c)

	}
}

func (us UserService) Endpoints() map[string]map[string]goframe.JSONEndpoint {
	return map[string]map[string]goframe.JSONEndpoint{
		"/restricted": {
			"GET": func(c goframe.APIContext) error {
				c.Response().GenericJSON("entry")
				return nil
			},
		},
		"/{userID:[0-9]+}": {
			"GET": func(c goframe.APIContext) error {
				vars := mux.Vars(c.Request())
				_userID := vars["userID"]
				userID, _ := strconv.ParseInt(_userID, 10, 64)
				user, err := us.store.Get(&c, userID)
				if err != nil {
					c.Response().JSON(404, goframe.ERROR, nil, "user not found")
					return nil
				}
				c.Response().JSON(200, goframe.OK, user, "")
				return nil
			},
			"PUT": func(c goframe.APIContext) error {
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
					c.Response().JSON(500, goframe.ERROR, nil, "invalid request")
					return nil
				}
				err := us.store.Insert(&c, store.User{
					ID:      userID,
					Name:    u.Name,
					Email:   u.Email,
					Enabled: true,
				})
				if err != nil {
					c.Response().JSON(500, goframe.ERROR, nil, "failed to insert user")
					return nil
				}
				c.Response().JSON(201, goframe.OK, nil, "inserted user")
				return nil
			},
		},
	}
}
