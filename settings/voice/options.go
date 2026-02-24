package voice

import "github.com/0x1eef/voiceai"

func WithID(id string) func(*voiceai.VoicePayload) {
	return func(p *voiceai.VoicePayload) {
		p.ID = id
	}
}

func WithPath(path string) func(*voiceai.VoicePayload) {
	return func(p *voiceai.VoicePayload) {
		p.Path = path
	}
}

func WithName(name string) func(*voiceai.VoicePayload) {
	return func(p *voiceai.VoicePayload) {
		p.Name = name
	}
}

func WithVisibility(visibility string) func(*voiceai.VoicePayload) {
	return func(p *voiceai.VoicePayload) {
		p.Visibility = visibility
	}
}

func WithLanguage(language string) func(*voiceai.VoicePayload) {
	return func(p *voiceai.VoicePayload) {
		p.Language = language
	}
}
