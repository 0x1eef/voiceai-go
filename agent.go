package voiceai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Agent struct {
	AgentPayload
	client *Client
}

type AgentPayload struct {
	AgentID    string             `json:"agent_id,omitempty"`
	Name       string             `json:"name"`
	Config     AgentConfigPayload `json:"config"`
	CreatedAt  string             `json:"created_at,omitempty"`
	UpdatedAt  string             `json:"updated_at,omitempty"`
	Status     string             `json:"status,omitempty"`
	StatusCode int                `json:"status_code,omitempty"`
	KBID       int                `json:"kb_id,omitempty"`
	Ctx        *context.Context   `json:"-"`
}

type AgentConfigPayload struct {
	Prompt                       string           `json:"prompt,omitempty"`
	Greeting                     string           `json:"greeting,omitempty"`
	LLMTemperature               float64          `json:"llm_temperature,omitempty"`
	LLMModel                     string           `json:"llm_model,omitempty"`
	TTSMinSentenceLen            float64          `json:"tts_min_sentence_len,omitempty"`
	TTSParams                    map[string]any   `json:"tts_params,omitempty"`
	MinSilenceDuration           float64          `json:"min_silence_duration,omitempty"`
	MinSpeechDuration            float64          `json:"min_speech_duration,omitempty"`
	UserSilenceTimeout           float64          `json:"user_silence_timeout,omitempty"`
	MaxCallDurationSeconds       float64          `json:"max_call_duration_seconds,omitempty"`
	AllowInterruptions           bool             `json:"allow_interruptions,omitempty"`
	AllowInterruptionsOnGreeting bool             `json:"allow_interruptions_on_greeting,omitempty"`
	MinInterruptionWords         float64          `json:"min_interruption_words,omitempty"`
	AutoNoiseReduction           bool             `json:"auto_noise_reduction,omitempty"`
	AllowAgentToEndCall          bool             `json:"allow_agent_to_end_call,omitempty"`
	AllowAgentToSkipTurn         bool             `json:"allow_agent_to_skip_turn,omitempty"`
	MinEndpointingDelay          float64          `json:"min_endpointing_delay,omitempty"`
	MaxEndpointingDelay          float64          `json:"max_endpointing_delay,omitempty"`
	VADActivationThreshold       float64          `json:"vad_activation_threshold,omitempty"`
	PhoneNumber                  string           `json:"phone_number,omitempty"`
	Webhooks                     map[string]any   `json:"webhooks,omitempty"`
	MCPServers                   []map[string]any `json:"mcp_servers,omitempty"`
}

func (c *Client) Agent() *Agent {
	return &Agent{client: c}
}

func (a *Agent) All(options ...func(*AgentPayload)) ([]Agent, error) {
	p := &AgentPayload{}
	for _, apply := range options {
		apply(p)
	}
	var agents []Agent
	res, err := a.client.get(p.Ctx, "/api/v1/agent", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&agents)
	if err != nil {
		return nil, err
	}
	for i := range agents {
		agents[i].client = a.client
	}
	return agents, nil
}

func (a *Agent) Create(options ...func(*AgentPayload)) (*Agent, error) {
	p := &AgentPayload{}
	for _, apply := range options {
		apply(p)
	}
	if p.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	b, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}
	res, err := a.client.post(p.Ctx, "/api/v1/agent", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return decodeAgent(res, a.client)
}

func (a *Agent) Deploy(ctx *context.Context) error {
	path := fmt.Sprintf("/api/v1/agent/%s/deploy", a.AgentID)
	_, err := a.client.post(ctx, path, nil, nil)
	return err
}

func (a *Agent) Disable(ctx *context.Context) error {
	path := fmt.Sprintf("/api/v1/agent/%s/disable", a.AgentID)
	_, err := a.client.post(ctx, path, nil, nil)
	return err
}

func (a *Agent) Pause(ctx *context.Context) error {
	path := fmt.Sprintf("/api/v1/agent/%s/pause", a.AgentID)
	_, err := a.client.post(ctx, path, nil, nil)
	return err
}

func decodeAgent(res *http.Response, client *Client) (*Agent, error) {
	var agent Agent
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&agent)
	if err != nil {
		return nil, err
	}
	agent.client = client
	return &agent, nil
}
