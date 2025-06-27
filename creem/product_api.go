package creem

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-pay/gopay"
)

// CreateProduct 创建产品
// 文档：https://docs.creem.io/api-reference/product#create-product
func (c *Client) CreateProduct(ctx context.Context, req *ProductCreateRequest) (rsp *ProductCreateResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 参数校验
	if req.Name == "" {
		return nil, MissNameErr
	}
	if req.Description == "" {
		return nil, MissDescriptionErr
	}
	if req.Type == "" {
		return nil, errors.New("product type is required")
	}
	if req.Price <= 0 {
		return nil, MissPriceErr
	}
	if req.Currency == "" {
		return nil, MissCurrencyErr
	}

	res, bs, err := c.doCreemPost(ctx, req, productCreate)
	if err != nil {
		return nil, err
	}

	rsp = &ProductCreateResponse{BaseResponse: BaseResponse{Code: Success}}
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

// GetProduct 获取产品详情
// 文档：https://docs.creem.io/api-reference/product#get-product
func (c *Client) GetProduct(ctx context.Context, productID string) (rsp *ProductDetailResponse, err error) {
	if productID == "" {
		return nil, MissProductIdErr
	}

	path := fmt.Sprintf(productDetail, productID)
	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &ProductDetailResponse{BaseResponse: BaseResponse{Code: Success}}
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

// ListProducts 获取产品列表
// 文档：https://docs.creem.io/api-reference/product#list-products
func (c *Client) ListProducts(ctx context.Context, params *ListParams) (rsp *ProductsListResponse, err error) {
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

	path := productsList
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &ProductsListResponse{BaseResponse: BaseResponse{Code: Success}}
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
