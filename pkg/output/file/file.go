package file

import (
	"fmt"

	"github.com/shima-park/mercury/pkg/channel"
	"github.com/shima-park/mercury/pkg/output"
)

func init() {
	err := output.Register("file", NewOutput)
	if err != nil {
		panic(err)
	}
}

type Output struct {
	ch     <-chan *channel.Message
	config interface{}
	done   chan struct{}
}

func NewOutput(config interface{}, publisher channel.Publisher, context *output.Context) (output.Output, error) {
	ch, err := publisher.Publish()
	if err != nil {
		return nil, err
	}

	return &Output{
		ch:     ch,
		config: config,
		done:   make(chan struct{}),
	}, nil
}

func (o *Output) Run() {
	for msg := range o.ch {
		select {
		case <-o.done:
			return
		default:
		}

		//time.Sleep(time.Second)
		fmt.Println("Timestamp:", msg.Timestamp, "Message:", string(msg.Message))
		msg.Ack(true)
	}
}

func (o *Output) Stop() {
	close(o.done)
}

func (o *Output) Wait() {

}
