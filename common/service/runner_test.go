package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRunnerInit(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	tr := TestRunner{}

	err = tr.Init(c, l, s)
	if assert.NoError(t, err, "Runner Init returns without error") {
		// test init override with failure
		tr.InitFunc = func(c Configurer, l Logger, s Storer) error {
			return errors.New("Init failed")
		}
		err = tr.Init(c, l, s)
		assert.Error(t, err, "Runner Init returns with error")
	}
}

func TestRunnerRouter(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	tr := TestRunner{}

	err = tr.Init(c, l, s)
	if assert.NoError(t, err, "Runner Init returns without error") {

		// test default routes
		router, err := tr.DefaultRouter()
		if assert.NoError(t, err, "DefaultRouter returns without error") {
			if assert.NotNil(t, router, "DefaultRouter returns a router") {

				// test default configured routes
				handle, params, redirect := router.Lookup(http.MethodGet, "/")
				if assert.NotNil(t, handle, "Handle is not nil") {
					t.Logf("Default route /")
					t.Logf("Have handler >%#v<", handle)
					t.Logf("Have params >%v<", params)
					t.Logf("Have redirect >%t<", redirect)
				}
			}
		}

		// test custom routes
		tr.RouterFunc = func(router *httprouter.Router) error {
			h, err := tr.DefaultMiddleware(tr.HandlerFunc)
			if err != nil {
				return err
			}
			router.GET("/custom", h)
			return nil
		}

		router, err = tr.DefaultRouter()
		if assert.NoError(t, err, "DefaultRouter returns without error") {
			if assert.NotNil(t, router, "DefaultRouter returns a router") {

				// test custom configured routes
				handle, params, redirect := router.Lookup(http.MethodGet, "/custom")
				if assert.NotNil(t, handle, "Handle is not nil") {
					t.Logf("Custom route /custom")
					t.Logf("Have handler >%#v<", handle)
					t.Logf("Have params >%v<", params)
					t.Logf("Have redirect >%t<", redirect)
				}
			}
		}

		// test custom routes error
		tr.RouterFunc = func(router *httprouter.Router) error {
			return errors.New("Failed router")
		}

		router, err = tr.DefaultRouter()
		if assert.Error(t, err, "DefaultRouter returns with error") {
			assert.Nil(t, router, "DefaultRouter returns nil")
		}
	}
}

func TestRunnerMiddleware(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	tr := TestRunner{}

	err = tr.Init(c, l, s)
	if assert.NoError(t, err, "Runner Init returns without error") {

		// test default middleware
		handle, err := tr.DefaultMiddleware(tr.HandlerFunc)
		if assert.NoError(t, err, "DefaultMiddleware returns without error") {
			if assert.NotNil(t, handle, "DefaultMiddleware returns a handle") {
				t.Logf("Have handle >%#v<", handle)
			}
		}

		// test custom middleware
		tr.MiddlewareFunc = func(h httprouter.Handle) (httprouter.Handle, error) {
			return nil, errors.New("Failed middleware")
		}

		handle, err = tr.DefaultMiddleware(tr.HandlerFunc)
		if assert.Error(t, err, "DefaultMiddleware returns with error") {
			assert.Nil(t, handle, "DefaultMiddleware returns nil")
		}
	}
}
