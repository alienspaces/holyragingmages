package client

import (
	"bytes"
	"encoding/json"
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
	Config configurer.Configurer
	Log    logger.Logger

	// NOTE: Maybe doesn't make sense to attach request config
	RequestConfig []RequestConfig
}

// RequestConfig -
type RequestConfig struct {
	Method string
	Path   string
	Host   string
}

// NewClient -
func NewClient(c configurer.Configurer, l logger.Logger) (*Client, error) {

	cl := Client{
		Config: c,
		Log:    l,
	}

	return &cl, nil
}

// RetryRequest -
func (c *Client) RetryRequest(rc RequestConfig, params map[string]string, r *Request) (*http.Response, error) {

	var err error
	var resp *http.Response

	retries := 0

RETRY:
	for {
		retries++

		resp, err = c.Request(rc, params, r)
		if err != nil {
			c.Log.Warn("Client request failed retries >%d< >%v<", retries, err)
			if retries == maxRetries {
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
func (c *Client) Request(rc RequestConfig, params map[string]string, r *Request) (*http.Response, error) {

	c.Log.Context("function", "Request")
	defer func() {
		c.Log.Context("function", "")
	}()

	var err error

	// Request URL
	url := rc.Host + rc.Path

	// Replace placeholder parameters
	for param, value := range params {
		url = strings.Replace(url, param, value, 1)
	}

	// data
	data, err := json.Marshal(r)
	if err != nil {
		c.Log.Warn("Failed marshalling request data >%v<", err)
		return nil, err
	}

	c.Log.Info("Client request URL >%s< Data >%v<", url, data)

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
		defer resp.Body.Close()
	case http.MethodPost, http.MethodPut:
		// Post / Put
		c.Log.Info("Method %s", rc.Method)
		req, err = http.NewRequest(rc.Method, url, bytes.NewBuffer(data))
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}
		resp, err = client.Do(req)
		if err != nil {
			c.Log.Warn("Failed client request >%v<", err)
			return nil, err
		}
		defer resp.Body.Close()
	case http.MethodDelete:
		// Delete
		c.Log.Info("Method Delete")
	default:
		// boom
	}

	c.Log.Info("Client response status >%s<", resp.Status)

	return resp, nil
}
