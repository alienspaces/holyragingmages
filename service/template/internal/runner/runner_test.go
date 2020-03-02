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

func TestGetTemplatesHandler(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	rnr := NewRunner()

	err = rnr.Init(c, l, s)
	require.NoError(t, err, "Init returns without error")

	// handler
	h, _ := rnr.DefaultMiddleware(rnr.HandlerConfig[0].Path, rnr.HandlerConfig[0].HandlerFunc)

	// router
	rtr := httprouter.New()
	rtr.GET(rnr.HandlerConfig[0].Path, h)

	// request
	r, err := http.NewRequest(rnr.HandlerConfig[0].Method, rnr.HandlerConfig[0].Path, nil)
	require.NoError(t, err, "NewRequest returns without error")

	// recorder
	w := httptest.NewRecorder()

	// serve
	rtr.ServeHTTP(w, r)

	// test status
	require.Equal(t, http.StatusOK, w.Code, "Create response status code is OK")
}

func TestPostTemplatesHandler(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	rnr := NewRunner()

	err = rnr.Init(c, l, s)
	require.NoError(t, err, "Init returns without error")

	// handler
	h, _ := rnr.DefaultMiddleware(rnr.HandlerConfig[2].Path, rnr.HandlerConfig[2].HandlerFunc)

	// router
	rtr := httprouter.New()
	rtr.POST(rnr.HandlerConfig[2].Path, h)

	// request
	rd := Request{
		Data: Data{
			ID: "e3a9e0f8-ce9c-477b-8b93-cf4da03af4c9",
		},
	}
	jd, err := json.Marshal(rd)
	require.NoError(t, err, "Marshal returns without error")

	req, err := http.NewRequest(rnr.HandlerConfig[2].Method, rnr.HandlerConfig[2].Path, bytes.NewBuffer(jd))
	require.NoError(t, err, "NewRequest returns without error")

	// recorder
	rec := httptest.NewRecorder()

	// serve
	rtr.ServeHTTP(rec, req)

	// test status
	require.Equal(t, http.StatusOK, rec.Code, "Create response status code is OK")

	res := Response{}
	err = json.NewDecoder(rec.Body).Decode(&res)

	require.NoError(t, err, "Decode returns without error")
	require.NotNil(t, res.Data, "Data is not nil")
	require.NotEmpty(t, res.Data.ID, "ID is not empty")
}
