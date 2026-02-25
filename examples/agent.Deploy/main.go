package main

import (
	"fmt"
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
	agent, err := client.Agent().Create(
		agent.WithName("Trebor"),
		agent.WithPrompt("Trebor is a helpful assistant"),
	)
	if err != nil {
		panic(err)
	}
	if err := agent.Deploy(); err != nil {
		panic(err)
	} else {
		fmt.Printf("Agent deployed\n")
	}
}
