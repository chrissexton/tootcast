package sinks

import (
	"github.com/chrissexton/tootcast/sinks/console"
	"github.com/chrissexton/tootcast/sinks/mastodon"
	"github.com/chrissexton/tootcast/util"
)

var All = map[string]util.MkSink{
	"mastodon": mastodon.New,
	"console":  console.New,
}
