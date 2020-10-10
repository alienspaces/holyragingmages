package server

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// Authen middleware tests
func TestAuthenMiddlewareConfig(t *testing.T) {

	// handlerFunc := func(w http.ResponseWriter, r *http.Request, pathParams httprouter.Params, queryParams map[string]interface{}, l logger.Logger, m modeller.Modeller) {
	// 	l.Info("Here I am lord!")
	// 	return
	// }

	// Test configuration
	type TestCase struct {
		name               string
		runner             func() TestRunner
		expectCachedConfig bool
	}

	tests := []TestCase{
		{
			name: "Without authentication",
			runner: func() TestRunner {
				r := TestRunner{}
				r.HandlerConfig = []HandlerConfig{
					{
						Method:           http.MethodPost,
						Path:             "/test",
						MiddlewareConfig: MiddlewareConfig{},
					},
				}

				return r
			},
			expectCachedConfig: false,
		},
		{
			name: "With authentication",
			runner: func() TestRunner {
				r := TestRunner{}
				r.HandlerConfig = []HandlerConfig{
					{
						Method: http.MethodPost,
						Path:   "/test",
						MiddlewareConfig: MiddlewareConfig{
							AuthTypes: []string{
								AuthTypeJWT,
							},
						},
					},
				}

				return r
			},
			expectCachedConfig: true,
		},
	}

	for _, tc := range tests {

		t.Logf("Running test >%s<", tc.name)

		c, l, s, err := NewDefaultDependencies()
		require.NoError(t, err, "NewDefaultDependencies returns without error")

		tr := tc.runner()

		err = tr.Init(c, l, s)
		require.NoError(t, err, "Runner Init returns without error")

		// Clear authen cache
		authenCache = nil

		nextHandlerFunc, err := tr.Authen("/test", nil)
		require.NoError(t, err, "Authen returns without error")
		require.NotNil(t, nextHandlerFunc, "Authen return the next handler function")

		if tc.expectCachedConfig == true {
			cached, ok := authenCache["/test"]
			require.True(t, ok, "Request path found in authen cache")
			t.Logf("Cached %v", cached)
		}
	}
}
