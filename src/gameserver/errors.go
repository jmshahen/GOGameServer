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
