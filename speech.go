package voiceai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Speech struct {
	client *Client
}

type SpeechPayload struct {
	Text     string  `json:"text"`
	VoiceID  *string `json:"voice_id,omitempty"`
	Format   *string `json:"audio_format,omitempty"`
	Temp     *string `json:"temperature,omitempty"`
	Model    *string `json:"model,omitempty"`
	Language *string `json:"language,omitempty"`
}

func (c *Client) Speech() *Speech {
	return &Speech{client: c}
}

func (s *Speech) Create(options ...func(*SpeechPayload)) (*http.Response, error) {
	p := &SpeechPayload{}
	for _, set := range options {
		set(p)
	}
	if p.Text == "" {
		return nil, errors.New("text is required")
	}
	if b, err := json.Marshal(p); err != nil {
		return nil, err
	} else {
		return s.client.post("/api/v1/tts/speech", nil, bytes.NewReader(b))
	}
}
