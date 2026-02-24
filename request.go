package voiceai

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) get(path string, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("https://%s%s", c.host, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	applyHeaders(req, headers)
	return request(req)
}

func (c *Client) post(path string, headers map[string]string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("https://%s%s", c.host, path)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")
	applyHeaders(req, headers)
	return request(req)
}

func (c *Client) delete(path string, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("https://%s%s", c.host, path)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	applyHeaders(req, headers)
	return request(req)
}

func (c *Client) patch(path string, headers map[string]string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("https://%s%s", c.host, path)
	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	applyHeaders(req, headers)
	return request(req)
}

func applyHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func request(req *http.Request) (*http.Response, error) {
	if res, err := http.DefaultClient.Do(req); err != nil {
		return res, err
	} else if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return res, fmt.Errorf("bad status: %d body: %s", res.StatusCode, string(body))
	} else {
		return res, nil
	}
}
