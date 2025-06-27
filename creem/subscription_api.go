package creem

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// GetSubscription 获取订阅详情
// 文档：https://docs.creem.io/api-reference/subscription#get-subscription
func (c *Client) GetSubscription(ctx context.Context, subscriptionID string) (rsp *SubscriptionDetailResponse, err error) {
	if subscriptionID == "" {
		return nil, MissSubscriptionIdErr
	}

	path := fmt.Sprintf(subscriptionDetail, subscriptionID)
	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &SubscriptionDetailResponse{BaseResponse: BaseResponse{Code: Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// UpdateSubscription 更新订阅
// 文档：https://docs.creem.io/api-reference/subscription#update-subscription
func (c *Client) UpdateSubscription(ctx context.Context, subscriptionID string, req *SubscriptionUpdateRequest) (rsp *SubscriptionUpdateResponse, err error) {
	if subscriptionID == "" {
		return nil, MissSubscriptionIdErr
	}
	if req == nil {
		return nil, errors.New("request is nil")
	}

	path := fmt.Sprintf(subscriptionUpdate, subscriptionID)
	res, bs, err := c.doCreemPost(ctx, req, path)
	if err != nil {
		return nil, err
	}

	rsp = &SubscriptionUpdateResponse{BaseResponse: BaseResponse{Code: Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// UpgradeSubscription 升级订阅
// 文档：https://docs.creem.io/api-reference/subscription#upgrade-subscription
func (c *Client) UpgradeSubscription(ctx context.Context, subscriptionID string, req *SubscriptionUpgradeRequest) (rsp *SubscriptionUpgradeResponse, err error) {
	if subscriptionID == "" {
		return nil, MissSubscriptionIdErr
	}
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.NewProductID == "" {
		return nil, MissProductIdErr
	}

	path := fmt.Sprintf(subscriptionUpgrade, subscriptionID)
	res, bs, err := c.doCreemPost(ctx, req, path)
	if err != nil {
		return nil, err
	}

	rsp = &SubscriptionUpgradeResponse{BaseResponse: BaseResponse{Code: Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// CancelSubscription 取消订阅
// 文档：https://docs.creem.io/api-reference/subscription#cancel-subscription
func (c *Client) CancelSubscription(ctx context.Context, subscriptionID string) (rsp *SubscriptionCancelResponse, err error) {
	if subscriptionID == "" {
		return nil, MissSubscriptionIdErr
	}

	path := fmt.Sprintf(subscriptionCancel, subscriptionID)
	res, bs, err := c.doCreemPost(ctx, nil, path)
	if err != nil {
		return nil, err
	}

	rsp = &SubscriptionCancelResponse{BaseResponse: BaseResponse{Code: Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}
