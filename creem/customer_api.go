package creem

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cloud-evan/gocreem"
)

// CustomersList 获取客户列表
// 文档：https://docs.creem.io/api/customers#list-customers
func (c *Client) CustomersList(ctx context.Context, params *ListParams) (rsp *CustomersListResponse, err error) {
	if params == nil {
		params = &ListParams{}
	}

	// 构建查询参数
	queryParams := url.Values{}
	if params.Page > 0 {
		queryParams.Set("page", strconv.Itoa(params.Page))
	}
	if params.Limit > 0 {
		queryParams.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.Status != "" {
		queryParams.Set("status", params.Status)
	}

	path := customersList
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &CustomersListResponse{BaseResponse: BaseResponse{Code: gocreem.Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gocreem.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// GetCustomer 获取客户详情
// 文档：https://docs.creem.io/api-reference/customer#get-customer
func (c *Client) GetCustomer(ctx context.Context, customerID string) (rsp *CustomerDetailResponse, err error) {
	if customerID == "" {
		return nil, MissCustomerIdErr
	}

	path := fmt.Sprintf(customerDetail, customerID)
	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &CustomerDetailResponse{BaseResponse: BaseResponse{Code: gocreem.Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gocreem.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// CustomerCreate 创建客户
// 文档：https://docs.creem.io/api/customers#create-customer
func (c *Client) CustomerCreate(ctx context.Context, req *CustomerCreateRequest) (rsp *CustomerCreateResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.Email == "" {
		return nil, MissEmailErr
	}
	if req.Name == "" {
		return nil, MissNameErr
	}

	res, bs, err := c.doCreemPost(ctx, req, customerCreate)
	if err != nil {
		return nil, err
	}

	rsp = &CustomerCreateResponse{BaseResponse: BaseResponse{Code: gocreem.Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gocreem.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusCreated {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// CustomerUpdate 更新客户
// 文档：https://docs.creem.io/api/customers#update-customer
func (c *Client) CustomerUpdate(ctx context.Context, customerID string, req *CustomerUpdateRequest) (rsp *CustomerUpdateResponse, err error) {
	if customerID == "" {
		return nil, MissCustomerIdErr
	}
	if req == nil {
		return nil, errors.New("request is nil")
	}

	path := fmt.Sprintf(customerUpdate, customerID)
	res, bs, err := c.doCreemPut(ctx, req, path)
	if err != nil {
		return nil, err
	}

	rsp = &CustomerUpdateResponse{BaseResponse: BaseResponse{Code: gocreem.Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gocreem.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusOK {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// CustomerDelete 删除客户
// 文档：https://docs.creem.io/api/customers#delete-customer
func (c *Client) CustomerDelete(ctx context.Context, customerID string) (rsp *BaseResponse, err error) {
	if customerID == "" {
		return nil, MissCustomerIdErr
	}

	path := fmt.Sprintf(customerDelete, customerID)
	res, bs, err := c.doCreemDelete(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &BaseResponse{Code: gocreem.Success}
	if res.StatusCode != http.StatusNoContent {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}

// CustomerPortalCreate 创建客户门户会话
// 文档：https://docs.creem.io/api/customer-portal#create-portal-session
func (c *Client) CustomerPortalCreate(ctx context.Context, req *CustomerPortalCreateRequest) (rsp *CustomerPortalCreateResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.CustomerID == "" {
		return nil, MissCustomerIdErr
	}
	if req.ReturnURL == "" {
		return nil, MissReturnUrlErr
	}

	res, bs, err := c.doCreemPost(ctx, req, customerPortalCreate)
	if err != nil {
		return nil, err
	}

	rsp = &CustomerPortalCreateResponse{BaseResponse: BaseResponse{Code: gocreem.Success}}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gocreem.UnmarshalErr, err, string(bs))
	}

	if res.StatusCode != http.StatusCreated {
		rsp.Code = res.StatusCode
		rsp.Error = string(bs)
		rsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, rsp.ErrorResponse)
	}

	return rsp, nil
}
