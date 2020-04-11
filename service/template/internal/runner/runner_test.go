package runner

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/store"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
)

// NewDefaultDependencies -
func NewDefaultDependencies() (configurer.Configurer, logger.Logger, storer.Storer, error) {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, nil, err
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
			return nil, nil, nil, err
		}
	}

	l, err := log.NewLogger(c)
	if err != nil {
		return nil, nil, nil, err
	}

	s, err := store.NewStore(c, l)
	if err != nil {
		return nil, nil, nil, err
	}

	err = s.Init()
	if err != nil {
		return nil, nil, nil, err
	}

	return c, l, s, nil
}

func TestNewRunner(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	r := NewRunner()

	err = r.Init(c, l, s)
	assert.NoError(t, err, "Init returns without error")
}

func TestTemplatesHandler(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	type TestCase struct {
		name         string
		config       func(rnr *Runner) service.HandlerConfig
		requestData  func() *Request
		responseCode int
	}

	tests := []TestCase{
		{
			name: "GET - Get basic resource",
			config: func(rnr *Runner) service.HandlerConfig {
				return rnr.HandlerConfig[0]
			},
			requestData: func() *Request {
				return nil
			},
			responseCode: http.StatusOK,
		},
		{
			name: "POST - Create basic resource",
			config: func(rnr *Runner) service.HandlerConfig {
				return rnr.HandlerConfig[2]
			},
			requestData: func() *Request {
				req := Request{
					Data: Data{
						ID: "e3a9e0f8-ce9c-477b-8b93-cf4da03af4c9",
					},
				}
				return &req
			},
			responseCode: http.StatusOK,
		},
		// TODO: add PUT and DELETE tests
	}

	for _, tc := range tests {

		t.Logf("Running test >%s<", tc.name)

		rnr := NewRunner()

		err = rnr.Init(c, l, s)
		require.NoError(t, err, "Init returns without error")

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

		// request
		rd := tc.requestData()

		var req *http.Request

		if rd != nil {
			jd, err := json.Marshal(rd)
			require.NoError(t, err, "Marshal returns without error")

			req, err = http.NewRequest(cfg.Method, cfg.Path, bytes.NewBuffer(jd))
			require.NoError(t, err, "NewRequest returns without error")
		} else {
			req, err = http.NewRequest(cfg.Method, cfg.Path, nil)
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
		require.NotNil(t, res.Data, "Data is not nil")
		require.NotEmpty(t, res.Data.ID, "ID is not empty")

		// TODO: test for a real date
		require.NotEmpty(t, res.Data.CreatedAt, "CreatedAt is not empty")
	}
}
