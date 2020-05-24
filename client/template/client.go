package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/alienspaces/holyragingmages/common/client"
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

// GetTemplate -
func (c *Client) GetTemplate(templateID string) (*Response, error) {

	resp, err := c.RetryRequest(
		client.RequestConfig{
			Method: http.MethodGet,
			Path:   "/api/templates/:template_id",
		},
		map[string]string{
			":template_id": templateID,
		},
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	respData := Response{}
	err = json.NewDecoder(resp.Body).Decode(&respData)
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
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	return &respData, err
}

// CreateTemplate -
func (c *Client) CreateTemplate(reqData *Request) (*Response, error) {

	resp, err := c.RetryRequest(
		client.RequestConfig{
			Method: http.MethodPost,
			Path:   "/api/templates",
		},
		nil,
		reqData,
	)
	if err != nil {
		msg := fmt.Sprintf("Failed request >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	respData := Response{}
	err = json.NewDecoder(resp.Body).Decode(&respData)
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
			Path:   "/api/templates/:template_id",
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
	err = json.NewDecoder(resp.Body).Decode(&respData)
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
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		msg := fmt.Sprintf("Failed decoding response >%v<", err)
		c.Log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	return &respData, err
}
