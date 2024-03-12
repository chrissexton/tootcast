// Code generated from Pkl module `tootcast.Config`. DO NOT EDIT.
package config

type Console interface {
	Sink
}

var _ Console = (*ConsoleImpl)(nil)

type ConsoleImpl struct {
	Type string `pkl:"type"`
}

func (rcv *ConsoleImpl) GetType() string {
	return rcv.Type
}
