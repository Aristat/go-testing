// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from injector.go:

func InitializeEvent(phrase1 string) (Event, func(), error) {
	message, cleanup := NewMessage(phrase1)
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		cleanup()
		return Event{}, nil, err
	}
	return event, func() {
		cleanup()
	}, nil
}