package sources

import (
	"github.com/chrissexton/tootcast/sources/discord"
	"github.com/chrissexton/tootcast/util"
)

var All = map[string]util.MkSrc{
	"discord": discord.New,
}
