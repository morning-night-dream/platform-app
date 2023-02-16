package controller

import (
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

var _ openapi.ServerInterface = (*Controller)(nil)

type Controller struct {
	client *Client
}

func New(
	client *Client,
) *Controller {
	return &Controller{
		client: client,
	}
}
