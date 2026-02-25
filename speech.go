package voiceai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
)

type Speech struct {
	client *Client
}

type SpeechPayload struct {
	Text     string           `json:"text"`
	VoiceID  *string          `json:"voice_id,omitempty"`
	Format   *string          `json:"audio_format,omitempty"`
	Temp     *string          `json:"temperature,omitempty"`
	Model    *string          `json:"model,omitempty"`
	Language *string          `json:"language,omitempty"`
	Ctx      *context.Context `json:"-"`
}

func (c *Client) Speech() *Speech {
	return &Speech{client: c}
}

func (s *Speech) Create(options ...func(*SpeechPayload)) (io.ReadCloser, error) {
	p := &SpeechPayload{}
	for _, apply := range options {
		apply(p)
	}
	if p.Text == "" {
		return nil, errors.New("text is required")
	}
	if b, err := json.Marshal(p); err != nil {
		return nil, err
	} else {
		r := bytes.NewReader(b)
		res, err := s.client.post(p.Ctx, "/api/v1/tts/speech", nil, r)
		if err != nil {
			return nil, err
		}
		return res.Body, nil
	}
}

func (s *Speech) Stream(options ...func(*SpeechPayload)) (io.ReadCloser, error) {
	p := &SpeechPayload{}
	for _, apply := range options {
		apply(p)
	}
	if p.Text == "" {
		return nil, errors.New("text is required")
	}
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	res, err := s.client.post(p.Ctx, "/api/v1/tts/speech/stream", nil, r)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
