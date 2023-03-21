package external

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/pkg/trace"
)

func NewRequestWithTID[T any](ctx context.Context, msg *T) *connect.Request[T] {
	req := connect.NewRequest(msg)

	req.Header().Set("tid", trace.GetTIDCtx(ctx))

	return req
}
