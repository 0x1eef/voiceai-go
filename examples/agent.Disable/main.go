package main

import (
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	agents, err := client.Agent().All()
	if err != nil {
		panic(err)
	}
	for _, agent := range agents {
		err := agent.Pause()
		if err != nil {
			continue
		}
		err = agent.Disable()
		if err != nil {
			panic(err)
		}
		println("disabled agent: ", agent.AgentID)
	}
}
