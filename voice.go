package voiceai

import (
	"bytes"
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
	ID         string `json:"voice_id"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Visibility string `json:"voice_visibility"`
	client     *Client
}

type VoicePayload struct {
	Path       string `json:"-"`
	ID         string `json:"-"`
	Name       string `json:"name,omitempty"`
	Visibility string `json:"voice_visibility,omitempty"`
	Language   string `json:"language,omitempty"`
}

func (c *Client) Voice() *Voice {
	return &Voice{client: c}
}

func (v *Voice) Delete(options ...func(*VoicePayload)) (*http.Response, error) {
	p := &VoicePayload{}
	for _, set := range options {
		set(p)
	}
	if p.ID == "" {
		return nil, fmt.Errorf("an ID is required")
	}
	return v.client.delete(fmt.Sprintf("/api/v1/tts/voice/%s", p.ID), nil)
}

func (v *Voice) Clone(options ...func(*VoicePayload)) (*http.Response, error) {
	headers := make(map[string]string)
	p := &VoicePayload{}
	for _, set := range options {
		set(p)
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
		return v.client.post("/api/v1/tts/clone-voice", headers, body)
	}
}

func (v *Voice) All() ([]Voice, error) {
	var voices []Voice
	res, err := v.client.get("/api/v1/tts/voices", nil)
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

func (v *Voice) Get(options ...func(*VoicePayload)) (*Voice, error) {
	var voice Voice
	p := &VoicePayload{}
	for _, set := range options {
		set(p)
	}
	if p.ID == "" {
		return nil, fmt.Errorf("an ID is required")
	}
	res, err := v.client.get(fmt.Sprintf("/api/v1/tts/voice/%s", p.ID), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&voice)
	if err != nil {
		return nil, err
	}
	return &voice, nil
}

func (v *Voice) Update(options ...func(*VoicePayload)) (*http.Response, error) {
	p := &VoicePayload{}
	for _, set := range options {
		set(p)
	}
	if p.ID == "" {
		return nil, fmt.Errorf("an ID is required")
	}
	b, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}
	return v.client.patch(fmt.Sprintf("/api/v1/tts/voice/%s", p.ID), nil, bytes.NewReader(b))
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
