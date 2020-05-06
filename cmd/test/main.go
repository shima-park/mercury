package main

import (
	"sync"

	"github.com/shima-park/mercury/pkg/channel"
	_ "github.com/shima-park/mercury/pkg/include"
	"github.com/shima-park/mercury/pkg/input"
	"github.com/shima-park/mercury/pkg/output"
)

func main() {
	c := channel.NewConnector()

	inputF, err := input.GetFactory("file")
	if err != nil {
		panic(err)
	}

	outputF, err := output.GetFactory("file")

	input, err := inputF(nil, c, nil)
	if err != nil {
		panic(err)
	}

	output, err := outputF(nil, c, nil)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go input.Run()
	go output.Run()

	wg.Wait()

}
