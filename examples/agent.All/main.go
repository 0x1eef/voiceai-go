package main

import (
	"fmt"
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
	fmt.Printf("%-40s %-24s %-12s %-8s\n", "ID", "NAME", "STATUS", "KB_ID")
	for _, agent := range agents {
		fmt.Printf(
			"%-40s %-24s %-12s %-8d\n",
			agent.AgentID,
			agent.Name,
			agent.Status,
			agent.KBID,
		)
	}
}
