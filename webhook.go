package slack

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

type WebHook struct {
	hookURL string
	ctx     context.Context
}

type WebHookPostPayload struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"username,omitempty"`
	IconUrl     string        `json:"icon_url,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"`
	UnfurlLinks bool          `json:"unfurl_links,omitempty"`
	LinkNames   int           `json:"link_names,omitempty"`
	Parse       string        `json:"parse,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

func NewWebHook(hookURL string, ctx context.Context) *WebHook {
	return &WebHook{hookURL: hookURL, ctx: ctx}
}

func (hk *WebHook) request(req *http.Request) ([]byte, error) {
	cl := urlfetch.Client(hk.ctx)
	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (hk *WebHook) PostMessage(payload *WebHookPostPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", hk.hookURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	_, err = hk.request(req)

	return err
}
