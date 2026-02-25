package main

import (
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/voice"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	voice, err := client.Voice().Get(
		voice.WithID("d1bf0f33-8e0e-4fbf-acf8-45c3c6262513"),
	)
	if err != nil {
		panic(err)
	}
	println(voice.Name)
}
