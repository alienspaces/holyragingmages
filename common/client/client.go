package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
)

const (
	maxRetries int = 5
	// AuthTypeBearer will use the AuthToken as Bearer data
	AuthTypeBearer string = "JWT"
	// AuthTypeBasic will use AuthUser and AuthPass as credentials
	AuthTypeBasic string = "Basic"
)

// Client -
type Client struct {
	Config configurer.Configurer
	Log    logger.Logger

	// RequestLogFunc will be called with the request URL, resulting request data and response data
	// to be used by client consumers wanting to store requests and responses for debugging etc
	RequestLogFunc func(url, requestData, responseData string)

	MaxRetries int
	// Path is the base path for all requests
	Path string
	// Host is the host for all requests
	Host string
	// AuthType is the method of authorization to use
	AuthType string
	// AuthToken is the authorization "token" to use in the case of authorization type "JWT"
	AuthToken string
	// AuthUser is the authorization "user" to use in the case of authorization type "Basic"
	AuthUser string
	// AuthPass is the authorization "password" to use in the case of authorization type "Basic"
	AuthPass string
}

// Request -
type Request struct {
	Pagination *RequestPagination `json:"pagination,omitempty"`
}

// RequestPagination -
type RequestPagination struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

// Response -
type Response struct {
	Error      *ResponseError      `json:"error,omitempty"`
	Pagination *ResponsePagination `json:"pagination,omitempty"`
}

// ResponseError -
type ResponseError struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

// ResponsePagination -
type ResponsePagination struct {
	Number int `json:"page_number"`
	Size   int `json:"page_size"`
	Count  int `json:"page_count"`
}

// NewClient -
func NewClient(c configurer.Configurer, l logger.Logger) (*Client, error) {

	cl := Client{
		Config: c,
		Log:    l,
	}

	return &cl, nil
}

// Init - override to perform custom initialization
func (c *Client) Init() error {

	c.Log.Debug("** Initialise **")

	if c.Config == nil {
		msg := "Configurer undefined, cannot init client"
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	// Host
	if c.Host == "" {
		msg := "Host undefined, cannot init client"
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	// AuthTypeBearer
	if c.AuthType == AuthTypeBearer {
		if c.AuthToken == "" {
			msg := "AuthType is AuthTypeBearer and AuthToken is undefined, cannot init client"
			c.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
	}

	// AuthTypeBasic
	if c.AuthType == AuthTypeBasic {
		if c.AuthUser == "" {
			msg := "AuthType is AuthTypeBasic and AuthUser is undefined, cannot init client"
			c.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
		if c.AuthPass == "" {
			msg := "AuthType is AuthTypeBasic and AuthUser is undefined, cannot init client"
			c.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
	}

	return nil
}

// RetryRequest -
func (c *Client) RetryRequest(method, path string, params map[string]string, reqData interface{}, respData interface{}) error {

	var err error

	// Initialise client
	err = c.Init()
	if err != nil {
		c.Log.Warn("Failed initialization >%v<", err)
		return err
	}

	var resp *http.Response

	// Replace placeholder parameters and add query parameters
	url, err := c.BuildURL(method, path, params)
	if err != nil {
		c.Log.Warn("Failed building URL >%v<", err)
		return err
	}

	// data
	data, err := c.EncodeData(reqData)
	if err != nil {
		c.Log.Warn("Failed marshalling request data >%v<", err)
		return err
	}

	// default maximum retries
	if c.MaxRetries == 0 {
		c.MaxRetries = maxRetries
	}
	retries := 0

RETRY:
	for {
		retries++

		resp, err = c.Request(method, url, data)
		if err != nil {
			c.Log.Warn("Client request failed retries >%d< >%v<", retries, err)
			if retries == c.MaxRetries {
				c.Log.Warn("Client request exceeded max retries, giving up now")
				return err
			}
			time.Sleep(time.Duration(retries) * time.Second)
			continue RETRY
		}
		break
	}

	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	return err
}

// Request -
func (c *Client) Request(method, url string, data []byte) (*http.Response, error) {

	c.Log.Context("function", "Request")
	defer func() {
		c.Log.Context("function", "")
	}()

	var err error

	c.Log.Debug("Client request URL >%s< data length >%d<", url, len(data))

	var resp *http.Response
	var req *http.Request

	// Request + Response logging
	var requestDump []byte
	var responseDump []byte

	client := &http.Client{}

	switch method {
	case http.MethodGet:

		// Get
		c.Log.Debug("Method %s", method)

		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}

		err := c.SetAuthHeaders(req)
		if err != nil {
			c.Log.Warn("Failed setting request auth headers >%v<", err)
			return nil, err
		}

		if c.RequestLogFunc != nil {
			requestDump, err = httputil.DumpRequest(req, true)
			if err != nil {
				c.Log.Warn("Failed request dump >%v<", err)
				return nil, err
			}
		}

		resp, err = client.Do(req)
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}

		if c.RequestLogFunc != nil {
			responseDump, err := httputil.DumpResponse(resp, true)
			if err != nil {
				c.Log.Warn("Failed response dump >%v<", err)
				return nil, err
			}
			c.RequestLogFunc(url, string(requestDump), string(responseDump))
		}

	case http.MethodPost, http.MethodPut:

		// Post / Put
		c.Log.Debug("Method %s", method)

		req, err = http.NewRequest(method, url, bytes.NewBuffer(data))
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}

		err := c.SetAuthHeaders(req)
		if err != nil {
			c.Log.Warn("Failed setting request auth headers >%v<", err)
			return nil, err
		}

		req.Header.Add("Content-Type", "application/json")

		if c.RequestLogFunc != nil {
			requestDump, err = httputil.DumpRequest(req, true)
			if err != nil {
				c.Log.Warn("Failed request dump >%v<", err)
				return nil, err
			}
		}

		resp, err = client.Do(req)
		if err != nil {
			c.Log.Warn("Failed client request >%#v< >%v<", resp, err)
			return nil, err
		}

		if c.RequestLogFunc != nil {
			responseDump, err = httputil.DumpResponse(resp, true)
			if err != nil {
				c.Log.Warn("Failed response dump >%v<", err)
				return nil, err
			}
			c.RequestLogFunc(url, string(requestDump), string(responseDump))
		}

	default:
		// boom
		msg := fmt.Sprintf("Method >%s< currently unsupported!", method)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	c.Log.Debug("Client response status >%s<", resp.Status)

	// Check response code
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		var responseData interface{}
		err = c.DecodeData(resp.Body, &responseData)
		err = fmt.Errorf("Response status >%d< body >%v<", resp.StatusCode, responseData)
	}

	return resp, err
}

// SetAuthHeaders sets authentication headers on an request object based
// on client authentication configuration
func (c *Client) SetAuthHeaders(req *http.Request) error {

	c.Log.Debug("Setting request authentication headers")

	// Auth type bearer with token
	if c.AuthType == AuthTypeBearer {
		var bearer = "Bearer " + c.AuthToken
		req.Header.Add("Authorization", bearer)
		return nil
	}

	// Auth type basic with user and pass
	if c.AuthType == AuthTypeBasic {
		req.SetBasicAuth(c.AuthUser, c.AuthPass)
		return nil
	}

	return nil
}

// BuildURL replaces placeholder parameters and adds query parameters
// The parameter "id" or ":id" has special behaviour. When provided the
// returned URL will have "/:id" appended and replaced with whatever
// the parameter value for "id" or ":id" was.
func (c *Client) BuildURL(method, requestURL string, params map[string]string) (string, error) {

	// Request URL
	requestURL = c.Host + c.Path + requestURL

	// Add resource identifier to URL when detected
	switch method {
	case http.MethodGet, http.MethodPost:
		if _, ok := params["id"]; ok {
			requestURL = requestURL + "/:id"
		}
		if _, ok := params[":id"]; ok {
			requestURL = requestURL + "/:id"
		}
	case http.MethodPut, http.MethodDelete:
		if _, ok := params["id"]; !ok {
			if _, ok := params[":id"]; !ok {
				msg := "Params must contain :id for method Put"
				c.Log.Warn(msg)
				return requestURL, fmt.Errorf(msg)
			}
		}
		requestURL = requestURL + "/:id"
	default:
		// no-op
	}

	// Replace placeholders and add query parameters
	paramString := ""
	for param, value := range params {

		// do not allow empty param values
		if value == "" {
			return requestURL, fmt.Errorf("Param >%s< has empty value", param)
		}

		found := false
		if strings.Index(requestURL, "/:"+param) != -1 {
			requestURL = strings.Replace(requestURL, "/:"+param, "/"+value, 1)
			found = true
		}
		if strings.Index(requestURL, "/"+param) != -1 {
			requestURL = strings.Replace(requestURL, "/"+param, "/"+value, 1)
			found = true
		}
		if !found {
			param = strings.Replace(param, ":", "", 1)
			if paramString != "" {
				paramString = paramString + "&"
			}
			paramString = paramString + param + "=" + url.QueryEscape(value)
		}
	}

	if paramString != "" {
		requestURL = requestURL + "?" + paramString
	}

	// do not allow missing parameters
	if strings.Index(requestURL, "/:") != -1 {
		return requestURL, fmt.Errorf("URL >%s< still contains placeholders", requestURL)
	}

	return requestURL, nil
}

// RegisterRequestLogFunc -
func (c *Client) RegisterRequestLogFunc(logFunc func(url, request, response string)) {
	c.RequestLogFunc = logFunc
}

// EncodeData is a convenience function that encodes struct data into bytes
func (c *Client) EncodeData(data interface{}) ([]byte, error) {

	c.Log.Context("function", "EncodeData")
	defer func() {
		c.Log.Context("function", "")
	}()

	dataBytes, err := json.Marshal(data)
	if err != nil {
		c.Log.Warn("Failed encoding data >%v<", err)
		return nil, err
	}
	return dataBytes, nil
}

// DecodeData is a convenience function that decodes bytes into struct data
func (c *Client) DecodeData(rc io.ReadCloser, data interface{}) error {

	c.Log.Context("function", "DecodeData")
	defer func() {
		c.Log.Context("function", "")
	}()

	// close before returning
	defer rc.Close()

	err := json.NewDecoder(rc).Decode(&data)
	if err != nil && err.Error() != "EOF" {
		c.Log.Warn("Failed decoding data >%v<", err)
		return err
	}
	return nil
}
