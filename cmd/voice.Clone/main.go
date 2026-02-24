package main

import (
	"fmt"
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
	res, err := client.Voice().Clone(
		voice.WithName("Trebor"),
		voice.WithPath("share/inputs/trebor.mp3"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", res)
}
