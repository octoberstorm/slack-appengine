package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

type WebHook struct {
	hookURL string
}

type WebHookPostPayload struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"username,omitempty"`
	IconUrl     string        `json:"icon_url,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"`
	UnfurlLinks bool          `json:"unfurl_links,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

func NewWebHook(hookURL string) *WebHook {
	return &WebHook{hookURL}
}

func request(req *http.Request) ([]byte, error) {
	cl := urlfetch.Client(sl.ctx)
	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (hk *WebHook) PostMessage(payload *WebHookPostPayload) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", hk.hookURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return request(req)
}
