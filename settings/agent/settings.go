package agent

import "github.com/0x1eef/voiceai"

func WithName(name string) func(*voiceai.AgentPayload) {
	return func(a *voiceai.AgentPayload) {
		a.Name = name
	}
}

func WithPrompt(prompt string) func(*voiceai.AgentPayload) {
	return func(a *voiceai.AgentPayload) {
		a.Config.Prompt = prompt
	}
}

func WithGreeting(greeting string) func(*voiceai.AgentPayload) {
	return func(a *voiceai.AgentPayload) {
		a.Config.Greeting = greeting
	}
}

func WithPhoneNumber(phone string) func(*voiceai.AgentPayload) {
	return func(a *voiceai.AgentPayload) {
		a.Config.PhoneNumber = phone
	}
}
