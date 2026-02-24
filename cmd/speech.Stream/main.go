package main

import (
	"io"
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/speech"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	stream, err := client.Speech().Stream(
		speech.WithText("Hello! My name is Trebor"),
		speech.WithVoiceID("f9e6a5eb-a7fd-4525-9e92-75125249c933"),
		speech.WithFormat("mp3"),
	)
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	out, _ := os.Create("share/outputs/trebor.mp3")
	defer out.Close()
	io.Copy(out, stream)
}
