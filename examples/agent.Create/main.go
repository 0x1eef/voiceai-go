package main

import (
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/agent"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	client.Agent().Create(
		agent.WithName("Trebor"),
		agent.WithPrompt("You are Trebor, a concise support assistant."),
	)
}
