package creem

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// CreateDiscountCode 创建优惠码
// 文档：https://docs.creem.io/api-reference/discount-code#create-discount-code
func (c *Client) CreateDiscountCode(ctx context.Context, req *DiscountCodeCreateRequest) (rsp *DiscountCodeCreateResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.Code == "" {
		return nil, errors.New("discount code is required")
	}
	if req.Type == "" {
		return nil, errors.New("discount type is required")
	}
	if req.Value <= 0 {
		return nil, errors.New("discount value must be greater than 0")
	}

	res, bs, err := c.doCreemPost(ctx, req, discountCodeCreate)
	if err != nil {
		return nil, err
	}

	rsp = &DiscountCodeCreateResponse{BaseResponse: BaseResponse{Code: Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusCreated {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// GetDiscountCode 获取优惠码详情
// 文档：https://docs.creem.io/api-reference/discount-code#get-discount-code
func (c *Client) GetDiscountCode(ctx context.Context, discountCodeID string) (rsp *DiscountCodeDetailResponse, err error) {
	if discountCodeID == "" {
		return nil, errors.New("discount code id is required")
	}

	path := fmt.Sprintf(discountCodeDetail, discountCodeID)
	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &DiscountCodeDetailResponse{BaseResponse: BaseResponse{Code: Success}}
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

// DeleteDiscountCode 删除优惠码
// 文档：https://docs.creem.io/api-reference/discount-code#delete-discount-code
func (c *Client) DeleteDiscountCode(ctx context.Context, discountCodeID string) (rsp *BaseResponse, err error) {
	if discountCodeID == "" {
		return nil, errors.New("discount code id is required")
	}

	path := fmt.Sprintf(discountCodeDelete, discountCodeID)
	res, bs, err := c.doCreemDelete(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &BaseResponse{Code: Success}
	if res.StatusCode != http.StatusNoContent {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}
