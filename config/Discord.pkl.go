// Code generated from Pkl module `tootcast.Config`. DO NOT EDIT.
package config

type Discord interface {
	Source

	GetBotToken() string

	GetGuildID() string

	GetChannelIDs() []string
}

var _ Discord = (*DiscordImpl)(nil)

type DiscordImpl struct {
	Type string `pkl:"type"`

	BotToken string `pkl:"botToken"`

	GuildID string `pkl:"guildID"`

	ChannelIDs []string `pkl:"channelIDs"`
}

func (rcv *DiscordImpl) GetType() string {
	return rcv.Type
}

func (rcv *DiscordImpl) GetBotToken() string {
	return rcv.BotToken
}

func (rcv *DiscordImpl) GetGuildID() string {
	return rcv.GuildID
}

func (rcv *DiscordImpl) GetChannelIDs() []string {
	return rcv.ChannelIDs
}
