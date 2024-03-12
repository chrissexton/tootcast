package util

import (
	"github.com/chrissexton/tootcast/msg"
)

type Sink interface {
	Serve()
	AddSource(chan msg.Message)
}
type MkSink func(any) Sink

type MkSrc func(any) Source
type Source interface {
	MsgChan() chan msg.Message
	Serve()
}
