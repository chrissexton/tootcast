// Code generated from Pkl module `tootcast.Config`. DO NOT EDIT.
package config

type Mastodon interface {
	Sink

	GetServer() string

	GetAccessToken() string
}

var _ Mastodon = (*MastodonImpl)(nil)

type MastodonImpl struct {
	Type string `pkl:"type"`

	Server string `pkl:"server"`

	AccessToken string `pkl:"accessToken"`
}

func (rcv *MastodonImpl) GetType() string {
	return rcv.Type
}

func (rcv *MastodonImpl) GetServer() string {
	return rcv.Server
}

func (rcv *MastodonImpl) GetAccessToken() string {
	return rcv.AccessToken
}
