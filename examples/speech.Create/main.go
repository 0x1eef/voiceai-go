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
	audio, err := client.Speech().Create(
		speech.WithText("Hello! My name is Trebor"),
		speech.WithVoiceID("trebors_voice_id"),
		speech.WithFormat("mp3"),
	)
	if err != nil {
		panic(err)
	}
	defer audio.Close()
	out, _ := os.Create("share/outputs/trebor.mp3")
	defer out.Close()
	io.Copy(out, audio)
}
