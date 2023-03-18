package cdekcalculator

import (
	"context"
)

const URLCalculator = "/calculator/tarifflist"

func (c *CDEKClient) Calculate(ctx context.Context, addrFrom, addrTo Location, size Size) ([]PriceSending, error) {
	cdekReq := &CDEKRequest{
		FromLocation: addrFrom,
		ToLocation:   addrTo,
		Packages:     []Size{size},
	}
	var zero []PriceSending

	req, err := c.createRequest(ctx, cdekReq, URLCalculator)
	if err != nil {
		return zero, err
	}

	cdekResp, err := DoRequest[CDEKResponse](req)
	if err != nil {
		return zero, err
	}

	return cdekResp.PriceSending, nil
}
