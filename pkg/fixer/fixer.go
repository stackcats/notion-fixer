package fixer

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type GetRatesResponse struct {
	Success bool               `json:"success"`
	Rates   map[string]float64 `json:"rates"`
}

const uri = "http://data.fixer.io/api/latest"

func GetRates(ctx context.Context, accessKey, base string) (map[string]float64, error) {
	client := resty.New()

	req, err := client.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"access_key": accessKey,
			"base":       base,
		}).
		SetResult(&GetRatesResponse{}).
		Get(uri)

	if err != nil {
		return nil, err
	}

	res := req.Result().(*GetRatesResponse)

	if res.Success {
		return res.Rates, nil
	}

	return nil, fmt.Errorf("fetch fixer error %#v", res)
}
