package creem

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-pay/gopay"
)

// ListTransactions 获取交易列表
// 文档：https://docs.creem.io/api-reference/transactions#list-transactions
func (c *Client) ListTransactions(ctx context.Context, params *ListParams) (rsp *TransactionsListResponse, err error) {
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
	if params.CustomerID != "" {
		queryParams.Set("customer_id", params.CustomerID)
	}
	if params.ProductID != "" {
		queryParams.Set("product_id", params.ProductID)
	}
	if params.StartDate != nil {
		queryParams.Set("start_date", params.StartDate.Format("2006-01-02"))
	}
	if params.EndDate != nil {
		queryParams.Set("end_date", params.EndDate.Format("2006-01-02"))
	}

	path := transactionsList
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	res, bs, err := c.doCreemGet(ctx, path)
	if err != nil {
		return nil, err
	}

	rsp = &TransactionsListResponse{BaseResponse: BaseResponse{Code: Success}}
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
