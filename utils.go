package cdekcalculator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DoRequest[T any](req *http.Request) (*T, error) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error http.Client: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error io.ReadAll: %w", err)
	}

	errResp := &ErrorResp{}
	if err := json.Unmarshal(body, errResp); err != nil {
		return nil, fmt.Errorf("error json.Unmarshal: %w", err)
	} else if err == nil && len(errResp.Requests) > 0 {
		return nil, fmt.Errorf("%v", *errResp)
	}

	var cdekResp T
	if err := json.Unmarshal(body, &cdekResp); err != nil {
		return nil, fmt.Errorf("error json.Unmarshal: %w", err)
	}
	return &cdekResp, nil
}
