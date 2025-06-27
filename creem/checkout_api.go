package creem

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gocreem/pkg/xhttp"
)

// CreateCheckoutSession 创建结账会话
// 文档：https://docs.creem.io/api-reference/checkout#create-checkout-session
func (c *Client) CreateCheckoutSession(ctx context.Context, req *CheckoutSessionCreateRequest) (rsp *CheckoutSessionResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.ProductID == "" {
		return nil, MissProductIdErr
	}
	if req.ReturnURL == "" {
		return nil, MissReturnUrlErr
	}
	if req.CancelURL == "" {
		return nil, MissCancelUrlErr
	}
	if req.SuccessURL == "" {
		return nil, MissSuccessUrlErr
	}

	res, bs, err := c.doCreemPost(ctx, req, checkoutSessionCreate)
	if err != nil {
		return nil, err
	}

	rsp = &CheckoutSessionResponse{BaseResponse: BaseResponse{Code: Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", xhttp.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusCreated {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// GetCheckoutSession 获取结账会话详情
// 文档：https://docs.creem.io/api-reference/checkout#get-checkout-session
func (c *Client) GetCheckoutSession(ctx context.Context, sessionID string) (rsp *CheckoutSessionResponse, err error) {
	if sessionID == "" {
		return nil, MissCheckoutSessionIdErr
	}

	path := fmt.Sprintf(checkoutSessionDetail, sessionID)
	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &CheckoutSessionResponse{BaseResponse: BaseResponse{Code: Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", xhttp.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}
