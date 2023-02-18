package health_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/e2e/helper"
	healthv1 "github.com/morning-night-dream/platform-app/pkg/connect/proto/health/v1"
)

func TestE2EHealthCheck(t *testing.T) {
	t.Parallel()

	url := helper.GetCoreEndpoint(t)

	t.Run("ヘルスチェックが成功する", func(t *testing.T) {
		t.Parallel()

		client := helper.NewConnectClient(t, http.DefaultClient, url)

		req := &healthv1.CheckRequest{}

		_, err := client.Health.Check(context.Background(), connect.NewRequest(req))
		if err != nil {
			t.Errorf("faile to health check: %s", err)
		}
	})
}
