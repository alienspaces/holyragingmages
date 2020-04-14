package runner

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
	"gitlab.com/alienspaces/holyragingmages/common/prepare"
	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/store"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/harness"
)

// NewDefaultDependencies -
func NewDefaultDependencies() (configurer.Configurer, logger.Logger, storer.Storer, preparer.Preparer, error) {

	// configurer
	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, nil, nil, err
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
			return nil, nil, nil, nil, err
		}
	}

	// logger
	l, err := log.NewLogger(c)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// storer
	s, err := store.NewStore(c, l)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// preparer
	p, err := prepare.NewPrepare(l)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = s.Init()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return c, l, s, p, nil
}

func NewTestHarness() (*harness.Testing, error) {

	// harness
	config := harness.DataConfig{}

	h, err := harness.NewTesting(config)
	if err != nil {
		return nil, err
	}

	// harness commit data
	h.CommitData = true

	return h, nil
}

func TestNewRunner(t *testing.T) {

	c, l, s, p, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	r := NewRunner()

	err = r.Init(c, l, s, p)
	require.NoError(t, err, "Init returns without error")
}

func TestTemplateHandler(t *testing.T) {

	// test dependencies
	c, l, s, p, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	// test harness
	th, err := NewTestHarness()
	require.NoError(t, err, "New test data returns without error")

	type TestCase struct {
		name          string
		config        func(rnr *Runner) service.HandlerConfig
		requestParams func(data *harness.Data) map[string]string
		requestData   func(data *harness.Data) *Request
		responseCode  int
	}

	tests := []TestCase{
		{
			name: "GET - Get existing resource",
			config: func(rnr *Runner) service.HandlerConfig {
				return rnr.HandlerConfig[1]
			},
			requestParams: func(data *harness.Data) map[string]string {
				params := map[string]string{
					":id": data.TemplateRecs[0].ID,
				}
				return params
			},
			requestData: func(data *harness.Data) *Request {
				return nil
			},
			responseCode: http.StatusOK,
		},
		{
			name: "GET - Get missing resource",
			config: func(rnr *Runner) service.HandlerConfig {
				return rnr.HandlerConfig[1]
			},
			requestParams: func(data *harness.Data) map[string]string {
				params := map[string]string{
					":id": "17c19414-2d15-4d20-8fc3-36fc10341dc8",
				}
				return params
			},
			requestData: func(data *harness.Data) *Request {
				return nil
			},
			responseCode: http.StatusNotFound,
		},
		{
			name: "POST - Create basic resource",
			config: func(rnr *Runner) service.HandlerConfig {
				return rnr.HandlerConfig[2]
			},
			requestData: func(data *harness.Data) *Request {
				req := Request{
					Data: Data{
						ID: "e3a9e0f8-ce9c-477b-8b93-cf4da03af4c9",
					},
				}
				return &req
			},
			responseCode: http.StatusOK,
		},
		{
			name: "PUT - Create basic resource",
			config: func(rnr *Runner) service.HandlerConfig {
				return rnr.HandlerConfig[3]
			},
			requestData: func(data *harness.Data) *Request {
				req := Request{
					Data: Data{
						ID: data.TemplateRecs[0].ID,
					},
				}
				return &req
			},
			responseCode: http.StatusOK,
		},
		{
			name: "PUT - Missing data",
			config: func(rnr *Runner) service.HandlerConfig {
				return rnr.HandlerConfig[3]
			},
			requestData: func(data *harness.Data) *Request {
				return nil
			},
			responseCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {

		t.Logf("Running test >%s<", tc.name)

		func() {
			rnr := NewRunner()

			err = rnr.Init(c, l, s, p)
			require.NoError(t, err, "Runner init returns without error")

			err = th.Setup()
			require.NoError(t, err, "Test data setup returns without error")
			defer func() {
				err = th.Teardown()
				require.NoError(t, err, "Test data teardown returns without error")
			}()

			// config
			cfg := tc.config(rnr)

			// handler
			h, _ := rnr.DefaultMiddleware(cfg.Path, cfg.HandlerFunc)

			// router
			rtr := httprouter.New()

			switch cfg.Method {
			case http.MethodGet:
				rtr.GET(cfg.Path, h)
			case http.MethodPost:
				rtr.POST(cfg.Path, h)
			case http.MethodPut:
				rtr.PUT(cfg.Path, h)
			case http.MethodDelete:
				rtr.DELETE(cfg.Path, h)
			default:
				//
			}

			// request params
			params := map[string]string{}
			if tc.requestParams != nil {
				params = tc.requestParams(th.Data)
			}

			requestPath := cfg.Path
			for paramKey, paramValue := range params {
				requestPath = strings.Replace(requestPath, paramKey, paramValue, 1)
			}

			// request data
			data := tc.requestData(th.Data)

			var req *http.Request

			if data != nil {
				jsonData, err := json.Marshal(data)
				require.NoError(t, err, "Marshal returns without error")

				req, err = http.NewRequest(cfg.Method, requestPath, bytes.NewBuffer(jsonData))
				require.NoError(t, err, "NewRequest returns without error")
			} else {
				req, err = http.NewRequest(cfg.Method, requestPath, nil)
				require.NoError(t, err, "NewRequest returns without error")
			}

			// recorder
			rec := httptest.NewRecorder()

			// serve
			rtr.ServeHTTP(rec, req)

			// test status
			require.Equal(t, tc.responseCode, rec.Code, "Response code equals expected")

			res := Response{}
			err = json.NewDecoder(rec.Body).Decode(&res)
			require.NoError(t, err, "Decode returns without error")

			// test data
			if tc.responseCode == http.StatusOK {
				require.NotEmpty(t, res.Data, "Data is not empty")
				require.NotEmpty(t, res.Data[0].ID, "ID is not empty")

				// TODO: test for a real date
				require.NotEmpty(t, res.Data[0].CreatedAt, "CreatedAt is not empty")
			}
		}()
	}
}
