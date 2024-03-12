package main

import (
	"context"
	"flag"
	"github.com/chrissexton/tootcast/config"
	"github.com/chrissexton/tootcast/sinks"
	"github.com/chrissexton/tootcast/sources"
	"log/slog"
)

var cfgFile = flag.String("config", "Configuration file", "config.pkl")

func main() {
	flag.Parse()

	slog.SetLogLoggerLevel(slog.LevelDebug)

	cfg, err := config.LoadFromPath(context.Background(), *cfgFile)
	if err != nil {
		panic(err)
	}

	sink := sinks.All[cfg.Sink.GetType()](cfg.Sink)
	source := sources.All[cfg.Source.GetType()](cfg.Source)
	sink.AddSource(source.MsgChan())
	sink.Serve()
}
