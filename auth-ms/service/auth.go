package service

import (
	"fmt"
)

const(
	greeting = "Hello my dear user"
)

type auth struct {}

func NewAuthenticator() Authenticator{
	return &auth{}
}

func (a *auth) Greet(name string) string{
	return fmt.Sprintf("%s %s", greeting, name)
}