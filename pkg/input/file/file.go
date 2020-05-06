package file

import (
	"fmt"
	"sync"
	"time"

	"github.com/shima-park/mercury/pkg/channel"
	"github.com/shima-park/mercury/pkg/input"
)

func init() {
	err := input.Register("file", NewInput)
	if err != nil {
		panic(err)
	}
}

type Input struct {
	ch     chan<- *channel.Message
	config interface{}
	done   chan struct{}
}

func NewInput(config interface{}, subscriber channel.Subscriber, context *input.Context) (input.Input, error) {
	ch, err := subscriber.Subscribe()
	if err != nil {
		return nil, err
	}

	return &Input{
		config: config,
		ch:     ch,
		done:   make(chan struct{}),
	}, nil
}

func (i *Input) Run() {
	var wg sync.WaitGroup

	send := func(msg *channel.Message) {
		wg.Add(1)
		i.ch <- msg

		select {
		case <-time.After(time.Second * 30): // TODO
			return
		default:
			wg.Wait()
		}
	}

	for {
		select {
		case <-i.done:
			return
		default:
		}

		var isAck bool
		ack := func(ack bool) {
			isAck = ack
			wg.Done()
		}

		msg := &channel.Message{
			Timestamp: time.Now(),
			Message:   []byte("test message"),
			Ack:       ack,
		}

		var retires = 3
		for !isAck && retires > 0 {
			fmt.Println("==========", isAck, retires)
			send(msg)
			retires--
		}
	}
}

func (i *Input) Stop() {
	close(i.done)
}

func (i *Input) Wait() {

}
