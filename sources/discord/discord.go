package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chrissexton/tootcast/config"
	"github.com/chrissexton/tootcast/msg"
	"github.com/chrissexton/tootcast/util"
	"log/slog"
)

type Connection struct {
	client  *discordgo.Session
	token   string
	guildID string

	channelIDs []string
	msgChan    chan msg.Message
}

func New(_cfg any) util.Source {
	cfg := _cfg.(config.Discord)
	c, err := discordgo.New(cfg.GetBotToken())
	if err != nil {
		return nil
	}
	conn := &Connection{
		client:     c,
		token:      cfg.GetBotToken(),
		guildID:    cfg.GetGuildID(),
		msgChan:    make(chan msg.Message, 1),
		channelIDs: cfg.GetChannelIDs(),
	}
	go conn.Serve()
	return conn
}

func (c *Connection) MsgChan() chan msg.Message {
	return c.msgChan
}

func folowedChannel(ch string, followed []string) bool {
	for _, id := range followed {
		if ch == id {
			return true
		}
	}
	return false
}

func (c *Connection) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !folowedChannel(m.ChannelID, c.channelIDs) {
		return
	}
	c.msgChan <- msg.Message{
		Author: m.Member.Nick,
		Body:   m.Content,
	}
}

func (c *Connection) Serve() {
	if err := c.setup(); err != nil {
		slog.Error("Error connecting: %w", err)
		return
	}
}

func (c *Connection) setup() error {
	c.client.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuilds |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsGuildEmojis |
			discordgo.IntentsGuildMessageReactions)
	err := c.client.Open()
	if err != nil {
		return err
	}
	c.client.AddHandler(c.messageCreate)
	return nil
}

func (c *Connection) Profile(id string) (*discordgo.Member, error) {
	mem, err := c.client.GuildMember(c.guildID, id)
	if err != nil {
		slog.Error("Error getting user", "err", err)
		return nil, err
	}
	return mem, nil
}
