package cdekcalculator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const TEST_URL = "https://api.edu.cdek.ru/v2"
const URL = "https://api.cdek.ru/v2"

type CDEKClient struct {
	token      string
	apiAddr    string
	isTestMode bool
}

func NewCDEKClient(token string, testMode bool) *CDEKClient {
	client := &CDEKClient{
		isTestMode: testMode,
	}
	client.Token(token)
	client.TestMode(testMode)
	return client
}

func (c *CDEKClient) TestMode(testMode bool) {
	if testMode {
		c.apiAddr = TEST_URL
	} else {
		c.apiAddr = URL
	}
	c.isTestMode = testMode
}

func (c *CDEKClient) IsTestMode() bool {
	return c.isTestMode
}

func (c *CDEKClient) Token(token string) {
	c.token = "Bearer " + token
}

func (c *CDEKClient) createRequest(ctx context.Context, cdekReq any, path string) (*http.Request, error) {
	newURL, err := url.JoinPath(c.apiAddr, path)
	if err != nil {
		return nil, fmt.Errorf("error url.JoinPath: %w", err)
	}

	jsonBody, err := json.Marshal(cdekReq)
	if err != nil {
		return nil, fmt.Errorf("error json.Marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, newURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error http.NewRequest: %w", err)
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.token)
	return req, nil
}
