package client

import (
	"encoding/json"
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
		respData, err := json.Marshal(&Response{})
		if err != nil {
			l.Warn("Failed encoding data >%v<", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(respData)
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
		{
			name:       "Get resource not OK",
			ID:         "",
			serverFunc: handlerFunc,
			expectErr:  true,
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
				require.Error(t, err, "GetTemplate returns with error")
				return
			}
			require.NoError(t, err, "GetTemplate returns without error")
			require.NotNil(t, resp, "GetTemplate returns a response")
		}()
	}
}

func TestGetTemplates(t *testing.T) {

	c, l, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	handlerFunc := func(rw http.ResponseWriter, req *http.Request) {
		respData, err := json.Marshal(&Response{})
		if err != nil {
			l.Warn("Failed encoding data >%v<", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(respData)
	}

	tests := []struct {
		name       string
		params     map[string]string
		serverFunc func(rw http.ResponseWriter, req *http.Request)
		expectErr  bool
	}{
		{
			name: "Get resources with ID OK",
			params: map[string]string{
				"id": "bceac4ab-738a-4e62-a040-835e6fab331f",
			},
			serverFunc: handlerFunc,
			expectErr:  false,
		},
		{
			name: "Get resources with params OK",
			params: map[string]string{
				"blah": "bceac4ab-738a-4e62-a040-835e6fab331f",
			},
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

			resp, err := cl.GetTemplates(tc.params)
			if tc.expectErr == true {
				require.Error(t, err, "GetTemplates returns with error")
				return
			}
			require.NoError(t, err, "GetTemplates returns without error")
			require.NotNil(t, resp, "GetTemplates returns a response")
		}()
	}
}
