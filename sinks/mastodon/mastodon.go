package mastodon

import (
	"context"
	"fmt"
	"github.com/chrissexton/tootcast/config"
	"github.com/chrissexton/tootcast/msg"
	"github.com/chrissexton/tootcast/util"
	"github.com/mattn/go-mastodon"
	"log"
	"log/slog"
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
	m.client = mastodon.NewClient(&mastodon.Config{
		Server:       m.cfg.GetServer(),
		ClientID:     m.cfg.GetClientID(),
		ClientSecret: m.cfg.GetClientSecret(),
		AccessToken:  m.cfg.GetAccessToken(),
	})
	err := m.client.Authenticate(context.Background(),
		m.cfg.GetUserName(), m.cfg.GetPassword())
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return m
}

func (m *Mastodon) message(in msg.Message) {
	fmt.Println(in)
	toot := mastodon.Toot{
		Status:     in.Body,
		Sensitive:  false,
		Visibility: "unlisted",
	}
	_, err := m.client.PostStatus(context.Background(), &toot)
	if err != nil {
		slog.Error("Error tooting", "err", err)
		return
	}
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
