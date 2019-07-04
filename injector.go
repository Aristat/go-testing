//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEvent(phrase1 string) (Event, func(), error) {
	panic(wire.Build(NewEvent, NewGreeter, NewMessage))
}
