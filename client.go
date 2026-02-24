package voiceai

type Client struct {
	host  string
	token string
}

func NewClient(options ...func(*Client)) (*Client, error) {
	c := &Client{
		host: "dev.voice.ai",
	}
	for _, set := range options {
		set(c)
	}
	return c, nil
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) SetHost(host string) {
	c.host = host
}
