package settings

import "github.com/0x1eef/voiceai"

func WithHost(host string) func(*voiceai.Client) {
	return func(c *voiceai.Client) {
		c.SetHost(host)
	}
}

func WithToken(token string) func(*voiceai.Client) {
	return func(c *voiceai.Client) {
		c.SetToken(token)
	}
}
