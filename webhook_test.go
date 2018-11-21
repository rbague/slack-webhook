package webhook_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rbague/slack-webhook"
)

type fakePoster struct{}

func (fakePoster) Post(_, _ string, body io.Reader) (*http.Response, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("could not read body: %v", err)
	}

	resp := &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(b))}
	resp.StatusCode = 200
	resp.Status = "200: OK"

	if string(b) == "null" { // nil payload
		resp.StatusCode = 400
		resp.Status = "400: Bad Request"
		return resp, nil
	}

	var p webhook.Payload
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, fmt.Errorf("could not marshall data: %v", err)
	}

	if len(p.Attachments) == 0 {
		if len(p.Text) == 0 {
			resp.StatusCode = 400
			resp.Status = "400: Bad Request"
		}
		return resp, nil
	}

	return resp, nil
}

func client() *webhook.Client {
	c := webhook.NewClient("")
	c.Client = new(fakePoster)
	return c
}

func TestSendSimple(t *testing.T) {
	tt := []struct {
		text string
		fail bool
	}{
		{"hello world", false},
		{"", true},
		{"\u2800", false},
	}

	client := client()
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%d_%v", i, tc.fail), func(t *testing.T) {
			err := client.SendSimple(tc.text)
			if tc.fail && err == nil {
				t.Fatal("should have failed but succeeded")
			}
			if !tc.fail && err != nil {
				t.Fatalf("should have not failed got: %v", err)
			}
		})
	}
}

func TestSend(t *testing.T) {
	p := &webhook.Payload{
		Attachments: []*webhook.Attachment{
			&webhook.Attachment{
				Text: "Hello, World!",
			},
		},
	}

	tt := []struct {
		payload *webhook.Payload
		fail    bool
	}{
		{nil, true},
		{&webhook.Payload{Attachments: []*webhook.Attachment{&webhook.Attachment{}}}, false},
		{p, false},
	}

	client := client()
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%d_%v", i, tc.fail), func(t *testing.T) {
			err := client.Send(tc.payload)
			if tc.fail && err == nil {
				t.Fatal("should have failed but succeeded")
			}
			if !tc.fail && err != nil {
				t.Fatalf("should have not failed got: %v", err)
			}
		})
	}
}
