package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
)

const (
	maxRetries int = 5
)

// Client -
type Client struct {
	Config     configurer.Configurer
	Log        logger.Logger
	MaxRetries int
	// Path is the base path for all requests
	Path string
	// Host is the host for all requests
	Host string

	// NOTE: Maybe doesn't make sense to attach request config
	RequestConfig []RequestConfig
}

// RequestConfig -
type RequestConfig struct {
	Method string
	Path   string
}

// Request -
type Request struct {
	Pagination RequestPagination `json:"pagination"`
}

// RequestPagination -
type RequestPagination struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

// Response -
type Response struct {
	Error      ResponseError      `json:"error"`
	Pagination ResponsePagination `json:"pagination"`
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

	err := cl.Init()
	if err != nil {
		return nil, err
	}

	return &cl, nil
}

// Init - override to perform custom initialization
func (c *Client) Init() error {

	c.Log.Info("** Initialise **")

	if c.Config == nil {
		msg := "Configurer undefined, cannot init client"
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	// host
	if c.Host == "" {
		c.Host = c.Config.Get("APP_HOST")
		c.Log.Info("Host undefined, sourced host >%s< from config", c.Host)
	}
	if c.Host == "" {
		msg := "Host undefined, cannot init client"
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	return nil
}

// GetResource -
func (c *Client) GetResource(path, templateID string, respData interface{}) error {

	c.Log.Context("function", "GetResource")
	defer func() {
		c.Log.Context("function", "")
	}()

	c.Log.Info("path >%s< templateID >%s< respData >%v<", path, templateID, respData)

	if templateID == "" {
		msg := fmt.Sprintf("Template ID is empty >%s<, cannot get template", templateID)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	resp, err := c.RetryRequest(
		RequestConfig{
			Method: http.MethodGet,
			Path:   path,
		},
		map[string]string{
			"id": templateID,
		},
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	return err
}

// GetResources -
func (c *Client) GetResources(path string, params map[string]string, respData interface{}) error {

	c.Log.Context("function", "GetResources")
	defer func() {
		c.Log.Context("function", "")
	}()

	c.Log.Info("path >%s< params >%#v< respData >%#v<", path, params, respData)

	resp, err := c.RetryRequest(
		RequestConfig{
			Method: http.MethodGet,
			Path:   path,
		},
		params,
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	return nil
}

// CreateResource -
func (c *Client) CreateResource(path, resourceID string, reqData interface{}, respData interface{}) error {

	c.Log.Context("function", "CreateResource")
	defer func() {
		c.Log.Context("function", "")
	}()

	c.Log.Info("path >%s< resourceID >%s< reqData >%#v< respData >%#v<", path, resourceID, reqData, respData)

	if reqData == nil {
		msg := fmt.Sprintf("Request data is nil >%v<, cannot create template", reqData)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	// id
	params := map[string]string{}
	if resourceID != "" {
		params["id"] = resourceID
	}

	resp, err := c.RetryRequest(
		RequestConfig{
			Method: http.MethodPost,
			Path:   path,
		},
		params,
		reqData,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	return err
}

// UpdateResource -
func (c *Client) UpdateResource(path, resourceID string, reqData interface{}, respData interface{}) error {

	c.Log.Context("function", "UpdateTemplate")
	defer func() {
		c.Log.Context("function", "")
	}()

	c.Log.Info("path >%s< resourceID >%s< reqData >%#v< respData >%#v<", path, resourceID, reqData, respData)

	if reqData == nil {
		msg := fmt.Sprintf("Request data is nil >%v<, cannot update template", reqData)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	// id
	params := map[string]string{}
	if resourceID != "" {
		params["id"] = resourceID
	}

	resp, err := c.RetryRequest(
		RequestConfig{
			Method: http.MethodPut,
			Path:   path,
		},
		params,
		reqData,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	return err
}

// DeleteResource -
func (c *Client) DeleteResource(path, resourceID string, respData interface{}) error {

	c.Log.Context("function", "DeleteResource")
	defer func() {
		c.Log.Context("function", "")
	}()

	if resourceID == "" {
		msg := fmt.Sprintf("Template ID is empty >%s<, cannot delete template", resourceID)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	resp, err := c.RetryRequest(
		RequestConfig{
			Method: http.MethodDelete,
			Path:   path,
		},
		map[string]string{
			"id": resourceID,
		},
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	return nil
}

// RetryRequest -
func (c *Client) RetryRequest(rc RequestConfig, params map[string]string, data interface{}) (*http.Response, error) {

	var err error
	var resp *http.Response

	// default maximum retries
	if c.MaxRetries == 0 {
		c.MaxRetries = maxRetries
	}
	retries := 0

RETRY:
	for {
		retries++

		resp, err = c.Request(rc, params, data)
		if err != nil {
			c.Log.Warn("Client request failed retries >%d< >%v<", retries, err)
			if retries == c.MaxRetries {
				c.Log.Warn("Client request exceeded max retries, giving up now")
				break
			}
			time.Sleep(time.Duration(retries) * time.Second)
			continue RETRY
		}
		break
	}

	return resp, err
}

// Request -
func (c *Client) Request(rc RequestConfig, params map[string]string, data interface{}) (*http.Response, error) {

	c.Log.Context("function", "Request")
	defer func() {
		c.Log.Context("function", "")
	}()

	var err error

	// Replace placeholder parameters and add query parameters
	url, err := c.BuildURL(rc.Method, rc.Path, params)
	if err != nil {
		c.Log.Warn("Failed building URL >%v<", err)
		return nil, err
	}

	// data
	dataBytes, err := c.EncodeData(data)
	if err != nil {
		c.Log.Warn("Failed marshalling request data >%v<", err)
		return nil, err
	}

	c.Log.Info("Client request URL >%s< Data >%v<", url, dataBytes)

	var resp *http.Response
	var req *http.Request

	client := &http.Client{}

	switch rc.Method {
	case http.MethodGet:
		// Get
		c.Log.Info("Method %s", rc.Method)
		req, err = http.NewRequest(rc.Method, url, nil)
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}
		resp, err = client.Do(req)
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}
	case http.MethodPost, http.MethodPut:
		// Post / Put
		c.Log.Info("Method %s", rc.Method)
		req, err = http.NewRequest(rc.Method, url, bytes.NewBuffer(dataBytes))
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}
		resp, err = client.Do(req)
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}
	case http.MethodDelete:
		// Delete
		c.Log.Info("Method Delete")
	default:
		// boom
	}

	c.Log.Info("Client response status >%s<", resp.Status)

	// Check response code
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		err = fmt.Errorf("Response status %d", resp.StatusCode)
	}

	return resp, err
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

// BuildURL replaces placeholder parameters and adds query parameters
func (c *Client) BuildURL(method, url string, params map[string]string) (string, error) {

	// Request URL
	url = c.Host + c.Path + url

	// Add resource identifier to URL when detected
	switch method {
	case http.MethodGet, http.MethodPost:
		if _, ok := params["id"]; ok {
			url = url + "/:id"
		}
		if _, ok := params[":id"]; ok {
			url = url + "/:id"
		}
	case http.MethodPut, http.MethodDelete:
		if _, ok := params["id"]; !ok {
			if _, ok := params[":id"]; !ok {
				msg := "Params must contain :id for method Put"
				c.Log.Warn(msg)
				return url, fmt.Errorf(msg)
			}
		}
		url = url + "/:id"
	default:
		// no-op
	}

	// Replace placeholders and add query parameters
	queryParamCount := 0
	for param, value := range params {
		found := false
		if strings.Index(url, "/:"+param) != -1 {
			url = strings.Replace(url, "/:"+param, "/"+value, 1)
			found = true
		}
		if strings.Index(url, "/"+param) != -1 {
			url = strings.Replace(url, "/"+param, "/"+value, 1)
			found = true
		}
		if !found {
			if queryParamCount == 0 {
				url = url + "?"
			}
			param = strings.Replace(param, ":", "", 1)
			url = url + param + "=" + value
			queryParamCount++
		}
	}

	return url, nil
}
