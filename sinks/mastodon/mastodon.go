package mastodon

import (
	"fmt"
	"github.com/chrissexton/tootcast/config"
	"github.com/chrissexton/tootcast/msg"
	"github.com/chrissexton/tootcast/util"
	"github.com/mattn/go-mastodon"
)

type Mastodon struct {
	sources []chan msg.Message
	cfg     config.Mastodon
	client  *mastodon.Client
}

func New(cfg any) util.Sink {
	m := &Mastodon{
		sources: make([]chan msg.Message, 1),
		cfg:     cfg.(config.Mastodon),
	}
	return m
}

func (m *Mastodon) message(in msg.Message) {
	fmt.Println(in)
	Toot(in.Body, MastodonCredentials{
		serverDomain: m.cfg.GetServer(),
		accessToken:  m.cfg.GetAccessToken(),
	})
}

func (m *Mastodon) Serve() {
	msgs := util.Merge(m.sources...)
	for {
		select {
		case msg := <-msgs:
			m.message(msg)
		}
	}
}

func (m *Mastodon) AddSource(src chan msg.Message) {
	m.sources = append(m.sources, src)
}
