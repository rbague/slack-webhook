[![GoDoc](https://godoc.org/github.com/rbague/slack-webhook?status.svg)](https://godoc.org/github.com/rbague/slack-webhook)
[![Go Report Card](https://goreportcard.com/badge/github.com/rbague/slack-webhook)](https://goreportcard.com/report/github.com/rbague/slack-webhook)
[![Build Status](https://travis-ci.org/rbague/slack-webhook.svg)](https://travis-ci.org/rbague/slack-webhook)

# slack-webhook
A minimal client for Slack's [Incoming Webhooks](https://api.slack.com/incoming-webhooks) API.

By default, uses http.DefaultClient to post the payload. This can be overridden by changing the Client variable in the Client struct

[embedmd]:# (webhook.go /type Client.*/ /}/)
```go
type Client struct {
	url    string
	Client Poster
}
```

### Usage:
[embedmd]:# (example/main.go)
```go
package main

import (
	"log"

	"github.com/rbague/slack-webhook"
)

func main() {
	client := webhook.NewClient("WEBHOOK_URL_HERE")

	err := client.SendSimple("Hello, World!")
	if err != nil {
		log.Fatalf("could not send simple message: %v", err)
	}
}
```

### Configuration
To get your webhook URL:
 1. go to [Incoming Webhooks](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks) Slack app
 2. choose your team and add a configuration
 3. choose the default channel and add integration

Further configuration can be added in the integration settings. 
The default configuration can also be overridden in the Payload struct

[embedmd]:# (payload.go /type Payload.*/ /}/)
```go
type Payload struct {
	Text string `json:"text"`

	// The channel where to send the payload to, or the configured channel
	// Can be both a channel '#other-channel' or a direct message '@username'
	Channel     string        `json:"channel,omitempty"`
	UserName    string        `json:"username,omitempty"`
	IconURL     string        `json:"icon_url,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"` // :ghost:
	UnfurlLinks bool          `json:"unfurl_links,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`

	// Markdown used to disable markdown formatting on the text field
	Markdown *bool `json:"mrkdwn,omitempty"`
}
```
