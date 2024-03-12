config: config/config.pkl
	pkl-gen-go config/config.pkl --base-path github.com/chrissexton/tootcast

tootcast: cmd/tootcast/main.go
	go build cmd/tootcast

all: config tootcast
.PHONY: config
