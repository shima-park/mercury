package channel

import "time"

type Subscriber interface {
	Subscribe() (chan<- *Message, error)
}

type Publisher interface {
	Publish() (<-chan *Message, error)
}

type Message struct {
	Timestamp time.Time
	Message   []byte
	Ack       func(isAck bool)
}

type Connector struct {
	ch chan *Message
}

func NewConnector() *Connector {
	return &Connector{
		ch: make(chan *Message),
	}
}

func (c *Connector) Subscribe() (chan<- *Message, error) {
	return c.ch, nil
}

func (c *Connector) Publish() (<-chan *Message, error) {
	return c.ch, nil
}
