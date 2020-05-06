package input

import (
	"fmt"

	"github.com/shima-park/mercury/pkg/channel"
	"github.com/shima-park/mercury/pkg/logger"
)

type Context struct {
	processDone chan struct{}
}

type Factory = func(config interface{}, connector channel.Subscriber, context *Context) (Input, error)

var registry = make(map[string]Factory)

func Register(name string, factory Factory) error {
	logger.Info("Registering input factory")
	if name == "" {
		return fmt.Errorf("Error registering input: name cannot be empty")
	}
	if factory == nil {
		return fmt.Errorf("Error registering input '%v': factory cannot be empty", name)
	}
	if _, exists := registry[name]; exists {
		return fmt.Errorf("Error registering input '%v': already registered", name)
	}

	registry[name] = factory
	logger.Info("Successfully registered input")

	return nil
}

func GetFactory(name string) (Factory, error) {
	if _, exists := registry[name]; !exists {
		return nil, fmt.Errorf("Error creating input. No such input type exist: '%v'", name)
	}
	return registry[name], nil
}
