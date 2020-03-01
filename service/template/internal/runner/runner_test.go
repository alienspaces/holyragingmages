package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

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
	if assert.NoError(t, err, "Init returns without error") {

		// handler
		h, _ := rnr.DefaultMiddleware(rnr.HandlerConfig[0].Path, rnr.HandlerConfig[0].HandlerFunc)

		// router
		rtr := httprouter.New()
		rtr.GET(rnr.HandlerConfig[0].Path, h)

		// request
		r, _ := http.NewRequest(rnr.HandlerConfig[0].Method, rnr.HandlerConfig[0].Path, nil)

		// recorder
		w := httptest.NewRecorder()

		// serve
		rtr.ServeHTTP(w, r)

		// test status
		if assert.Equal(t, http.StatusOK, w.Code, "Create response status code is OK") {
			// TODO: check response
		}
	}
}
