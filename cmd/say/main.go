package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/speech"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		log.Fatalf("%s", err)
	}
	stream, err := client.Speech().Stream(
		speech.WithText(strings.Join(os.Args[1:], " ")),
		speech.WithFormat("mp3"),
	)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer stream.Close()
	cmd := exec.Command("mpv", "--no-video", "-")
	cmd.Stdin = stream
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s", err)
	}
}

func init() {
	log.SetFlags(0)
}
