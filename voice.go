package voiceai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Voice struct {
	ID         string `json:"voice_id,omitempty"`
	Status     string `json:"status,omitempty"`
	Name       string `json:"name,omitempty"`
	Visibility string `json:"voice_visibility,omitempty"`
	client     *Client
}

type VoicePayload struct {
	Path       string           `json:"-"`
	ID         string           `json:"-"`
	Name       string           `json:"name,omitempty"`
	Visibility string           `json:"voice_visibility,omitempty"`
	Language   string           `json:"language,omitempty"`
	Ctx        *context.Context `json:"-"`
}

func (v *Voice) All(options ...func(*VoicePayload)) ([]Voice, error) {
	p := &VoicePayload{}
	for _, set := range options {
		set(p)
	}
	var voices []Voice
	res, err := v.client.get(p.Ctx, "/api/v1/tts/voices", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&voices)
	if err != nil {
		return nil, err
	}
	return voices, nil
}

func (v *Voice) Clone(options ...func(*VoicePayload)) (*Voice, error) {
	headers := make(map[string]string)
	p := &VoicePayload{}
	for _, apply := range options {
		apply(p)
	}
	if p.Path == "" {
		return nil, errors.New("a path is required")
	}
	f, err := os.Open(p.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if w, body, err := newMultiPart(p, f); err != nil {
		return nil, err
	} else {
		headers["Content-Type"] = w.FormDataContentType()
		return decodeVoice(
			v.client.post(p.Ctx, "/api/v1/tts/clone-voice", headers, body),
		)
	}
}

func (v *Voice) Delete(options ...func(*VoicePayload)) (*Voice, error) {
	p := &VoicePayload{}
	for _, apply := range options {
		apply(p)
	}
	if p.ID == "" {
		return nil, fmt.Errorf("an ID is required")
	}
	return decodeVoice(
		v.client.delete(p.Ctx, fmt.Sprintf("/api/v1/tts/voice/%s", p.ID), nil),
	)
}

func (v *Voice) Get(options ...func(*VoicePayload)) (*Voice, error) {
	p := &VoicePayload{}
	for _, apply := range options {
		apply(p)
	}
	if p.ID == "" {
		return nil, fmt.Errorf("an ID is required")
	}
	return decodeVoice(
		v.client.get(p.Ctx, fmt.Sprintf("/api/v1/tts/voice/%s", p.ID), nil),
	)
}

func (v *Voice) Update(options ...func(*VoicePayload)) (*Voice, error) {
	p := &VoicePayload{}
	for _, apply := range options {
		apply(p)
	}
	if p.ID == "" {
		return nil, fmt.Errorf("an ID is required")
	}
	b, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}
	return decodeVoice(
		v.client.patch(p.Ctx, fmt.Sprintf("/api/v1/tts/voice/%s", p.ID), nil, bytes.NewReader(b)),
	)
}

func (c *Client) Voice() *Voice {
	return &Voice{client: c}
}

func newMultiPart(p *VoicePayload, f *os.File) (*multipart.Writer, *bytes.Buffer, error) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	defer w.Close()
	if err := writeField(w, "name", p.Name); err != nil {
		return nil, nil, err
	}
	if err := writeField(w, "voice_visibility", p.Visibility); err != nil {
		return nil, nil, err
	}
	if err := writeField(w, "language", p.Language); err != nil {
		return nil, nil, err
	}
	if part, err := w.CreateFormFile("file", filepath.Base(p.Path)); err != nil {
		return nil, nil, err
	} else {
		_, err := io.Copy(part, f)
		return w, &body, err
	}
}

func writeField(w *multipart.Writer, key string, value string) error {
	if value == "" {
		return nil
	}
	return w.WriteField(key, value)
}

func decodeVoice(res *http.Response, err error) (*Voice, error) {
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var voice Voice
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&voice)
	if err != nil {
		return nil, err
	}
	return &voice, nil
}
