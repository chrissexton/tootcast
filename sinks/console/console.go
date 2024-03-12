package console

import (
	"fmt"
	"github.com/chrissexton/tootcast/msg"
	"github.com/chrissexton/tootcast/util"
)

type Console struct {
	sources []chan msg.Message
}

func New(config any) util.Sink {
	return &Console{make([]chan msg.Message, 0)}
}

func (c *Console) AddSource(src chan msg.Message) {
	c.sources = append(c.sources, src)
}

func (c *Console) Serve() {
	msgs := util.Merge(c.sources...)
	for {
		select {
		case msg := <-msgs:
			fmt.Println(msg)
		}
	}
}
