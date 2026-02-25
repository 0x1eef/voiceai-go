package voiceai

import (
	"bytes"
	"encoding/json"
	"fmt"
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
}

type AgentConfigPayload struct {
	Prompt                       string           `json:"prompt,omitempty"`
	Greeting                     string           `json:"greeting,omitempty"`
	LLMTemperature               float64          `json:"llm_temperature,omitempty"`
	LLMModel                     string           `json:"llm_model,omitempty"`
	TTSMinSentenceLen            int              `json:"tts_min_sentence_len,omitempty"`
	TTSParams                    map[string]any   `json:"tts_params,omitempty"`
	MinSilenceDuration           float64          `json:"min_silence_duration,omitempty"`
	MinSpeechDuration            float64          `json:"min_speech_duration,omitempty"`
	UserSilenceTimeout           int              `json:"user_silence_timeout,omitempty"`
	MaxCallDurationSeconds       int              `json:"max_call_duration_seconds,omitempty"`
	AllowInterruptions           bool             `json:"allow_interruptions,omitempty"`
	AllowInterruptionsOnGreeting bool             `json:"allow_interruptions_on_greeting,omitempty"`
	MinInterruptionWords         int              `json:"min_interruption_words,omitempty"`
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

func (a *Agent) Create(options ...func(*AgentPayload)) (*Agent, error) {
	var agent Agent
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
	res, err := a.client.post("/api/v1/agent", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&agent)
	return &agent, nil
}

func (c *Client) Agent() *Agent {
	return &Agent{client: c}
}
