package creem

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// ValidateLicense 校验授权密钥
// 文档：https://docs.creem.io/api-reference/license#validate-license-key
func (c *Client) ValidateLicense(ctx context.Context, req *LicenseValidateRequest) (rsp *LicenseValidateResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.LicenseKey == "" {
		return nil, errors.New("license key is required")
	}

	res, bs, err := c.doCreemPost(ctx, req, licenseValidate)
	if err != nil {
		return nil, err
	}

	rsp = &LicenseValidateResponse{BaseResponse: BaseResponse{Code: Success}}
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

// ActivateLicense 激活授权密钥
// 文档：https://docs.creem.io/api-reference/license#activate-license-key
func (c *Client) ActivateLicense(ctx context.Context, req *LicenseActivateRequest) (rsp *LicenseActivateResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.LicenseKey == "" {
		return nil, errors.New("license key is required")
	}
	if req.CustomerID == "" {
		return nil, MissCustomerIdErr
	}

	res, bs, err := c.doCreemPost(ctx, req, licenseActivate)
	if err != nil {
		return nil, err
	}

	rsp = &LicenseActivateResponse{BaseResponse: BaseResponse{Code: Success}}
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

// DeactivateLicense 注销授权密钥
// 文档：https://docs.creem.io/api-reference/license#deactivate-license-key
func (c *Client) DeactivateLicense(ctx context.Context, req *LicenseDeactivateRequest) (rsp *LicenseDeactivateResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.LicenseKey == "" {
		return nil, errors.New("license key is required")
	}

	res, bs, err := c.doCreemPost(ctx, req, licenseDeactivate)
	if err != nil {
		return nil, err
	}

	rsp = &LicenseDeactivateResponse{BaseResponse: BaseResponse{Code: Success}}
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
