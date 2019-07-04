package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

type Message string

type Greeter struct {
	Grumpy  bool
	Message Message
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func main() {
	fmt.Printf("Hi, %s \n", "Sasha")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(main.NewExecutableSchema(Config{Resolvers: &main.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	//e, _, err := InitializeEvent("Test1")
	//
	//if err != nil {
	//	fmt.Printf("failed to create event: %s\n", err)
	//	os.Exit(2)
	//}
	//e.Start()
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewMessage(phrase string) (Message, func()) {
	cleanup := func() {
		fmt.Println("Funcccc")
	}

	return Message(phrase), cleanup
}

func NewGreeter(m Message) Greeter {
	var grumpy bool

	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}

	return Greeter{Message: m, Grumpy: grumpy}
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}
