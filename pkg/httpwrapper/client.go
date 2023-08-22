package httpwrapper

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HTTPClient struct {
	BaseURL        string
	Client         *http.Client
	DefaultHeaders map[string]string
	Token          string
}

func NewClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (c *HTTPClient) NewRequest(method, path string, headers map[string]string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL + path)

	if err != nil {
		return nil, fmt.Errorf("failed to parse Full Endpoint url: %w", err)
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to make NewRequest: %w", err)
	}

	for k, v := range c.DefaultHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func (c *HTTPClient) SetToken(token string) {
	c.Token = token
}

func (c *HTTPClient) GetToken() string {
	return c.Token
}

func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-AUTH-TOKEN", c.Token)
	return c.Client.Do(req)
}
