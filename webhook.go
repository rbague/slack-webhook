package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Poster interface has the required methods of an http.Client
// It is mainly used to ease testing.
type Poster interface {
	Post(url, contentType string, body io.Reader) (*http.Response, error)
}

// Client a client for the Slack's Incoming Webhooks API
type Client struct {
	url    string
	Client Poster
}

// NewClient creates a Slack Incoming Webhook client that uses
// http.DefaultClient to send the requests
func NewClient(url string) *Client {
	return &Client{url: url, Client: http.DefaultClient}
}

// SendSimple sends a simple text message
func (c *Client) SendSimple(text string) error {
	return c.Send(&Payload{Text: text})
}

// Send sends the given payload to the client's webhook
func (c *Client) Send(p *Payload) error {
	b, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("could not marshall data: %v", err)
	}

	resp, err := c.Client.Post(c.url, "application/json", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("could not send message: %v", err)
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("webhook failed with: %s", resp.Status)
	}
	return nil
}
