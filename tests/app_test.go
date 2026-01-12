package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyborgvova/echoprint/app"
	"github.com/cyborgvova/echoprint/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Response struct {
	Message string `json:"message"`
}

func TestGoodRequest(t *testing.T) {
	port := 8888
	tests := []struct {
		name string
		path string
		text string
		want int
	}{
		{
			name: "first",
			path: "/",
			text: "Hello World !!!",
			want: http.StatusOK,
		}, {
			name: "second",
			path: "/ready",
			text: "ready",
			want: http.StatusOK,
		}, {
			name: "third",
			path: "/health",
			text: "health",
			want: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			cfg := &config.Config{
				Text: test.text,
				Port: port,
			}

			application := app.New(cfg)
			ts := httptest.NewServer(application.Handler())
			defer ts.Close()

			req, err := http.NewRequestWithContext(ctx, "GET", ts.URL+test.path, nil)
			require.NoError(t, err, "Failed to create request")

			resp, err := ts.Client().Do(req)
			require.NoError(t, err, "HTTP request failed")
			defer resp.Body.Close()

			assert.Equal(t, test.want, resp.StatusCode)

			bodyBytes, err := io.ReadAll(resp.Body)
			require.NoError(t, err, "Failed to read response body")

			var res Response
			err = json.Unmarshal(bodyBytes, &res)
			require.NoError(t, err, "Failed to unmarshal JSON")

			assert.Equal(t, test.text, res.Message)
		})
	}
}
