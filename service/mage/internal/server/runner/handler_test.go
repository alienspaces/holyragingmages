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

	"gitlab.com/alienspaces/holyragingmages/common/server"
	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/harness"
)

func TestMageHandler(t *testing.T) {

	// test dependencies
	c, l, s, p, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	// test harness
	th, err := NewTestHarness()
	require.NoError(t, err, "New test data returns without error")

	type TestCase struct {
		name          string
		config        func(rnr *Runner) server.HandlerConfig
		requestParams func(data *harness.Data) map[string]string
		requestData   func(data *harness.Data) *Request
		responseCode  int
		responseData  func(data *harness.Data) *Response
	}

	tests := []TestCase{
		{
			name: "GET - Get existing resource",
			config: func(rnr *Runner) server.HandlerConfig {
				return rnr.HandlerConfig[1]
			},
			requestParams: func(data *harness.Data) map[string]string {
				params := map[string]string{
					":id": data.MageRecs[0].ID,
				}
				return params
			},
			requestData: func(data *harness.Data) *Request {
				return nil
			},
			responseCode: http.StatusOK,
			responseData: func(data *harness.Data) *Response {
				res := Response{
					Data: []Data{
						{
							ID:           data.MageRecs[0].ID,
							Name:         data.MageRecs[0].Name,
							Strength:     data.MageRecs[0].Strength,
							Dexterity:    data.MageRecs[0].Dexterity,
							Intelligence: data.MageRecs[0].Intelligence,
							Experience:   data.MageRecs[0].Experience,
							Coin:         data.MageRecs[0].Coin,
						},
					},
				}
				return &res
			},
		},
		{
			name: "GET - Get missing resource",
			config: func(rnr *Runner) server.HandlerConfig {
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
			responseData: func(data *harness.Data) *Response {
				return nil
			},
		},
		{
			name: "POST - Create basic resource",
			config: func(rnr *Runner) server.HandlerConfig {
				return rnr.HandlerConfig[2]
			},
			requestData: func(data *harness.Data) *Request {
				req := Request{
					Data: Data{
						ID:   "e3a9e0f8-ce9c-477b-8b93-cf4da03af4c9",
						Name: "Legislate Law",
					},
				}
				return &req
			},
			responseCode: http.StatusOK,
			responseData: func(data *harness.Data) *Response {
				res := Response{
					Data: []Data{
						{
							ID:           "e3a9e0f8-ce9c-477b-8b93-cf4da03af4c9",
							Name:         "Legislate Law",
							Strength:     0,
							Dexterity:    0,
							Intelligence: 0,
							Experience:   0,
							Coin:         0,
						},
					},
				}
				return &res
			},
		},
		{
			name: "PUT - Update basic resource",
			config: func(rnr *Runner) server.HandlerConfig {
				return rnr.HandlerConfig[3]
			},
			requestParams: func(data *harness.Data) map[string]string {
				params := map[string]string{
					":id": data.MageRecs[0].ID,
				}
				return params
			},
			requestData: func(data *harness.Data) *Request {
				req := Request{
					Data: Data{
						ID:           data.MageRecs[0].ID,
						Name:         "Barricade Block",
						Strength:     data.MageRecs[0].Strength,
						Dexterity:    data.MageRecs[0].Dexterity,
						Intelligence: data.MageRecs[0].Intelligence,
						Experience:   data.MageRecs[0].Experience,
						Coin:         data.MageRecs[0].Coin,
					},
				}
				return &req
			},
			responseCode: http.StatusOK,
			responseData: func(data *harness.Data) *Response {
				res := Response{
					Data: []Data{
						{
							ID:           data.MageRecs[0].ID,
							Name:         "Barricade Block",
							Strength:     data.MageRecs[0].Strength,
							Dexterity:    data.MageRecs[0].Dexterity,
							Intelligence: data.MageRecs[0].Intelligence,
							Experience:   data.MageRecs[0].Experience,
							Coin:         data.MageRecs[0].Coin,
						},
					},
				}
				return &res
			},
		},
		{
			name: "PUT - Missing data",
			config: func(rnr *Runner) server.HandlerConfig {
				return rnr.HandlerConfig[3]
			},
			requestData: func(data *harness.Data) *Request {
				return nil
			},
			responseCode: http.StatusBadRequest,
			responseData: func(data *harness.Data) *Response {
				return nil
			},
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
			reqData := tc.requestData(th.Data)

			var req *http.Request

			if reqData != nil {
				jsonData, err := json.Marshal(reqData)
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

			// response data
			var resData *Response
			if tc.responseData != nil {
				resData = tc.responseData(th.Data)
			}

			// test data
			if tc.responseCode == http.StatusOK {
				require.NotEmpty(t, res.Data, "Data is not empty")

				// response data
				if resData != nil {
					require.Equal(t, resData.Data[0].ID, res.Data[0].ID, "ID equals expected")
					require.Equal(t, resData.Data[0].Name, res.Data[0].Name, "Name equals expected")
					require.Equal(t, resData.Data[0].Strength, res.Data[0].Strength, "Strength equals expected")
					require.Equal(t, resData.Data[0].Dexterity, res.Data[0].Dexterity, "Dexterity equals expected")
					require.Equal(t, resData.Data[0].Intelligence, res.Data[0].Intelligence, "Intelligence equals expected")
					require.Equal(t, resData.Data[0].Experience, res.Data[0].Experience, "Experience equals expected")
					require.Equal(t, resData.Data[0].Coin, res.Data[0].Coin, "Coin equals expected")
				}

				// record timestamps
				require.False(t, res.Data[0].CreatedAt.IsZero(), "CreatedAt is not zero")
				if cfg.Method == http.MethodPost {
					require.True(t, res.Data[0].UpdatedAt.IsZero(), "UpdatedAt is zero")
				}
				if cfg.Method == http.MethodPut {
					require.False(t, res.Data[0].UpdatedAt.IsZero(), "UpdatedAt is not zero")
				}
			}
		}()
	}
}
