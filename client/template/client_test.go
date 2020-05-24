package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
)

// NewDefaultDependencies -
func NewDefaultDependencies() (configurer.Configurer, logger.Logger, error) {

	// configurer
	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, err
	}

	configVars := []string{
		// general
		"APP_HOST",
		// logger
		"APP_LOG_LEVEL",
	}
	for _, key := range configVars {
		err = c.Add(key, false)
		if err != nil {
			return nil, nil, err
		}
	}

	// logger
	l, err := log.NewLogger(c)
	if err != nil {
		return nil, nil, err
	}

	return c, l, nil
}

func TestGetTemplate(t *testing.T) {

	c, l, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	handlerFunc := func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}

	tests := []struct {
		name       string
		ID         string
		serverFunc func(rw http.ResponseWriter, req *http.Request)
		expectErr  bool
	}{
		{
			name:       "Get resource OK",
			ID:         "bceac4ab-738a-4e62-a040-835e6fab331f",
			serverFunc: handlerFunc,
			expectErr:  false,
		},
	}

	for _, tc := range tests {

		t.Logf("Running test >%s<", tc.name)

		func() {
			// Test HTTP server
			server := httptest.NewServer(http.HandlerFunc(tc.serverFunc))
			defer server.Close()

			// Set environment
			c.Set("APP_HOST", server.URL)

			// Client
			cl, err := NewClient(c, l)
			require.NoError(t, err, "NewClient returns without error")
			require.NotNil(t, cl, "NewClient returns a client")

			// set max retries to speed tests up
			cl.MaxRetries = 2

			resp, err := cl.GetTemplate(tc.ID)
			if tc.expectErr == true {
				require.Error(t, err, "RetryRequest returns with error")
				return
			}
			require.NoError(t, err, "RetryRequest returns without error")
			require.NotNil(t, resp, "RetryRequest returns a response")
		}()
	}
}
