package gameserver

import (
	"fmt"
	"time"
)

type ServerError struct {
	When time.Time
	What string
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func NewServerError(s string) error {
	return &ServerError{
		time.Now(),
		s,
	}
}

type NetworkError struct {
	When time.Time
	What string
	Msg  string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("at %v, %s\n\"%s\"", e.When, e.What, e.Msg)
}

func NewNetworkError(what string, msg string) error {
	return &NetworkError{
		time.Now(),
		what,
		msg,
	}
}
