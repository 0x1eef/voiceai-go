package agent

import "github.com/0x1eef/voiceai"

func WithName(name string) func(*voiceai.AgentPayload) {
	return func(a *voiceai.AgentPayload) {
		a.Name = name
	}
}

func WithPrompt(prompt string) func(*voiceai.AgentPayload) {
	return func(c *voiceai.AgentPayload) {
		c.Config.Prompt = prompt
	}
}
