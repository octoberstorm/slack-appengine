package slack

import (
	"golang.org/x/net/context"
	"net/url"
)

type Slack struct {
	token string          // token
	ctx   context.Context // app engine context
}

// Create a slack client with an API token & an app engine context
func New(token string, ctx context.Context) *Slack {
	return &Slack{
		token: token,
		ctx:   ctx,
	}
}

func (sl *Slack) urlValues() *url.Values {
	uv := url.Values{}
	uv.Add("token", sl.token)
	return &uv
}
