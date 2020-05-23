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
		// logger
		"APP_LOG_LEVEL",
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
	}
	for _, key := range configVars {
		err = c.Add(key, true)
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

func TestClient(t *testing.T) {

	c, l, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	cl, err := NewClient(c, l)
	require.NoError(t, err, "NewClient returns without error")
	require.NotNil(t, cl, "NewClient returns a client")

	// set max retries to speed tests up
	cl.MaxRetries = 2

	type RequestData struct {
		Request
		Name         string
		Strength     int
		Intelligence int
		Dexterity    int
	}

	tests := []struct {
		name                 string
		config               RequestConfig
		params               map[string]string
		requestData          *RequestData
		expectURL            string
		serverFunc           func(rw http.ResponseWriter, req *http.Request)
		expectErr            bool
		expectResponseStatus int
	}{
		{
			name: "Get resource with expected path",
			config: RequestConfig{
				Method: http.MethodGet,
				Path:   "/api/mages/:mage_id",
			},
			params: map[string]string{
				"mage_id": "52fdfc07-2182-454f-963f-5f0f9a621d72",
			},
			serverFunc: func(rw http.ResponseWriter, req *http.Request) {

				if req.URL.String() != "/api/mages/52fdfc07-2182-454f-963f-5f0f9a621d72" {
					rw.WriteHeader(http.StatusBadRequest)
					return
				}

				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(`OK`))
			},
			expectErr:            false,
			expectResponseStatus: http.StatusOK,
		},
		{
			name: "Get resource with unexpected path",
			config: RequestConfig{
				Method: http.MethodGet,
				Path:   "/api/kobolds/:kobold_id",
			},
			params: map[string]string{
				"owlbear_id": "52fdfc07-2182-454f-963f-5f0f9a621d72",
			},
			serverFunc: func(rw http.ResponseWriter, req *http.Request) {

				if req.URL.String() != "/api/kobolds/52fdfc07-2182-454f-963f-5f0f9a621d72" {
					rw.WriteHeader(http.StatusBadRequest)
					return
				}

				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(`OK`))
			},
			expectErr:            true,
			expectResponseStatus: http.StatusBadRequest,
		},
		{
			name: "Post resource with expected path",
			config: RequestConfig{
				Method: http.MethodPost,
				Path:   "/api/orcs/:orc_id",
			},
			params: map[string]string{
				"orc_id": "52fdfc07-2182-454f-963f-5f0f9a621d72",
			},
			requestData: &RequestData{
				Name:     "Brain Basher",
				Strength: 10,
			},
			serverFunc: func(rw http.ResponseWriter, req *http.Request) {

				if req.URL.String() != "/api/orcs/52fdfc07-2182-454f-963f-5f0f9a621d72" {
					rw.WriteHeader(http.StatusBadRequest)
					return
				}

				requestData := RequestData{}
				err := json.NewDecoder(req.Body).Decode(&requestData)
				if err != nil {
					rw.WriteHeader(http.StatusBadRequest)
					return
				}

				if requestData.Name != "Brain Basher" {
					rw.WriteHeader(http.StatusBadRequest)
					return
				}
				if requestData.Strength != 10 {
					rw.WriteHeader(http.StatusBadRequest)
					return
				}

				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(`OK`))
			},
			expectErr:            false,
			expectResponseStatus: http.StatusOK,
		},
	}

	for _, tc := range tests {

		t.Logf("Running test >%s<", tc.name)

		func() {
			// Test HTTP server
			server := httptest.NewServer(http.HandlerFunc(tc.serverFunc))
			defer server.Close()

			cfg := tc.config
			cfg.Host = server.URL

			resp, err := cl.RetryRequest(cfg, tc.params, tc.requestData)
			if tc.expectErr == true {
				require.Error(t, err, "RetryRequest returns with error")
				return
			}
			require.NoError(t, err, "RetryRequest returns without error")
			require.NotNil(t, resp, "RetryRequest returns a response")
			require.Equal(t, resp.StatusCode, tc.expectResponseStatus)
		}()
	}
}
