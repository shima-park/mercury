package output

import (
	"fmt"

	"github.com/shima-park/mercury/pkg/channel"
	"github.com/shima-park/mercury/pkg/logger"
)

type Context struct {
	processDone chan struct{}
}

type Factory = func(config interface{}, connector channel.Publisher, context *Context) (Output, error)

var registry = make(map[string]Factory)

func Register(name string, factory Factory) error {
	logger.Info("Registering output factory")
	if name == "" {
		return fmt.Errorf("Error registering output: name cannot be empty")
	}
	if factory == nil {
		return fmt.Errorf("Error registering output '%v': factory cannot be empty", name)
	}
	if _, exists := registry[name]; exists {
		return fmt.Errorf("Error registering output '%v': already registered", name)
	}

	registry[name] = factory
	logger.Info("Successfully registered output")

	return nil
}

func GetFactory(name string) (Factory, error) {
	if _, exists := registry[name]; !exists {
		return nil, fmt.Errorf("Error creating output. No such output type exist: '%v'", name)
	}
	return registry[name], nil
}
