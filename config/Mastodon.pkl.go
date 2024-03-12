// Code generated from Pkl module `tootcast.Config`. DO NOT EDIT.
package config

type Mastodon interface {
	Sink

	GetServer() string

	GetClientID() string

	GetClientSecret() string

	GetAccessToken() string

	GetUserName() string

	GetPassword() string
}

var _ Mastodon = (*MastodonImpl)(nil)

type MastodonImpl struct {
	Type string `pkl:"type"`

	Server string `pkl:"server"`

	ClientID string `pkl:"clientID"`

	ClientSecret string `pkl:"clientSecret"`

	AccessToken string `pkl:"accessToken"`

	UserName string `pkl:"userName"`

	Password string `pkl:"password"`
}

func (rcv *MastodonImpl) GetType() string {
	return rcv.Type
}

func (rcv *MastodonImpl) GetServer() string {
	return rcv.Server
}

func (rcv *MastodonImpl) GetClientID() string {
	return rcv.ClientID
}

func (rcv *MastodonImpl) GetClientSecret() string {
	return rcv.ClientSecret
}

func (rcv *MastodonImpl) GetAccessToken() string {
	return rcv.AccessToken
}

func (rcv *MastodonImpl) GetUserName() string {
	return rcv.UserName
}

func (rcv *MastodonImpl) GetPassword() string {
	return rcv.Password
}
