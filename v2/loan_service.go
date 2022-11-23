package binance

import (
	"context"
	"net/http"
)

type LoadOrdersService struct {
	c *Client
}

func (s *LoadOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *LoadOrdersResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/ongoing/orders",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(LoadOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type LoadOrdersResponse struct {
	Total int `json:"total"`
	Rows  []struct {
		LoanCoin         string `json:"loanCoin"`
		CollateralCoin   string `json:"collateralCoin"`
		TotalDebt        string `json:"totalDebt"`
		ResidualInterest string `json:"residualInterest"`
		CurrentLTV       string `json:"currentLTV"`
		CollateralAmount string `json:"collateralAmount"`
		ExpirationTime   string `json:"expirationTime"`
		OrderID          string `json:"orderId"`
	} `json:"rows"`
}
