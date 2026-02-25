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
	voices, err := client.Voice().All()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%-40s %-24s %-12s %-12s\n", "ID", "NAME", "STATUS", "VISIBILITY")
	for _, voice := range voices {
		fmt.Printf(
			"%-40s %-24s %-12s %-12s\n",
			voice.ID,
			voice.Name,
			voice.Status,
			voice.Visibility,
		)
	}
}
