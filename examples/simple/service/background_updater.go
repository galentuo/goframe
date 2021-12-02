package service

import "fmt"

type UserUpdater struct {
}

func NewUserUpdater() *UserUpdater {
	return &UserUpdater{}
}

func (u UserUpdater) Name() string {
	return "user-updater"
}

func (u UserUpdater) Run() error {
	fmt.Println("Starting user updater")
	return nil
}
