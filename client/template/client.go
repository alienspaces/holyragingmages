package client

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/alienspaces/holyragingmages/common/client"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
)

// Client -
type Client struct {
	client.Client
}

// Request -
type Request struct {
	client.Request
	Data TemplateData `json:"data"`
}

// Response -
type Response struct {
	client.Response
	Data []TemplateData `json:"data"`
}

// TemplateData -
type TemplateData struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// NewClient -
func NewClient(c configurer.Configurer, l logger.Logger) (*Client, error) {

	cl := Client{
		client.Client{
			Config: c,
			Log:    l,
		},
	}

	// Base path for all requests
	cl.Path = "/templates"

	err := cl.Init()
	if err != nil {
		return nil, err
	}

	return &cl, nil
}

// GetTemplate -
func (c *Client) GetTemplate(templateID string) (*Response, error) {

	resp, err := c.RetryRequest(
		client.RequestConfig{
			Method: http.MethodGet,
			Path:   "/api/templates",
		},
		map[string]string{
			":id": templateID,
		},
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	respData := Response{}
	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	return &respData, err
}

// GetTemplates -
func (c *Client) GetTemplates(params map[string]string) (*Response, error) {

	resp, err := c.RetryRequest(
		client.RequestConfig{
			Method: http.MethodGet,
			Path:   "/api/templates",
		},
		params,
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	respData := Response{}
	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	return &respData, err
}

// CreateTemplate -
func (c *Client) CreateTemplate(reqData *Request) (*Response, error) {

	// id
	params := map[string]string{}
	if reqData.Data.ID != "" {
		params["id"] = reqData.Data.ID
	}

	resp, err := c.RetryRequest(
		client.RequestConfig{
			Method: http.MethodPost,
			Path:   "/api/templates",
		},
		params,
		reqData,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	respData := Response{}
	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	return &respData, err
}

// UpdateTemplate -
func (c *Client) UpdateTemplate(reqData *Request) (*Response, error) {

	resp, err := c.RetryRequest(
		client.RequestConfig{
			Method: http.MethodPut,
			Path:   "/api/templates",
		},
		map[string]string{
			"id": reqData.Data.ID,
		},
		reqData,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	respData := Response{}
	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	return &respData, err
}

// DeleteTemplate -
func (c *Client) DeleteTemplate(templateID string) (*Response, error) {

	resp, err := c.RetryRequest(
		client.RequestConfig{
			Method: http.MethodDelete,
			Path:   "/api/templates/:template_id",
		},
		map[string]string{
			"id": templateID,
		},
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	respData := Response{}
	err = c.DecodeData(resp.Body, &respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	return &respData, err
}
