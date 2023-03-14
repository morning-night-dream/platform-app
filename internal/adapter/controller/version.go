package controller

import (
	"context"

	"github.com/bufbuild/connect-go"
	versionv1 "github.com/morning-night-dream/platform-app/pkg/connect/version/v1"
)

type Version struct {
	version string
}

func NewVersion(
	version string,
) *Version {
	return &Version{
		version: version,
	}
}

func (h *Version) Confirm(
	ctx context.Context,
	req *connect.Request[versionv1.ConfirmRequest],
) (*connect.Response[versionv1.ConfirmResponse], error) {
	return connect.NewResponse(&versionv1.ConfirmResponse{Version: h.version}), nil
}
