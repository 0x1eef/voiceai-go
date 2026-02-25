package speech

import "github.com/0x1eef/voiceai"

func WithContext(ctx context.Context) func(*voiceai.SpeechPayload) {
	return func(p *voiceai.SpeechPayload) {
		p.Ctx = &ctx
	}
}

func WithText(text string) func(*voiceai.SpeechPayload) {
	return func(p *voiceai.SpeechPayload) {
		p.Text = text
	}
}

func WithVoiceID(voiceID string) func(*voiceai.SpeechPayload) {
	return func(p *voiceai.SpeechPayload) {
		p.VoiceID = &voiceID
	}
}

func WithFormat(format string) func(*voiceai.SpeechPayload) {
	return func(p *voiceai.SpeechPayload) {
		p.Format = &format
	}
}

func WithTemp(temp string) func(*voiceai.SpeechPayload) {
	return func(p *voiceai.SpeechPayload) {
		p.Temp = &temp
	}
}

func WithModel(model string) func(*voiceai.SpeechPayload) {
	return func(p *voiceai.SpeechPayload) {
		p.Model = &model
	}
}

func WithLanguage(language string) func(*voiceai.SpeechPayload) {
	return func(p *voiceai.SpeechPayload) {
		p.Language = &language
	}
}
