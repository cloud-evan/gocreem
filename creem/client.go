package creem

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/gocreem/pkg/xhttp"
	"github.com/gocreem/pkg/xlog"
)

// Client Creem支付客户端
type Client struct {
	ApiKey       string
	SecretKey    string
	IsProd       bool
	DebugSwitch  DebugSwitch
	logger       xlog.XLogger
	hc           *xhttp.Client
	baseUrlProd  string
	headerKeyMap map[string]string
}

type Option func(*Client)

// NewClient 初始化Creem支付客户端
// apiKey: API密钥
// secretKey: 密钥
// isProd: 是否是正式环境（Creem只有生产环境）
func NewClient(apiKey, secretKey string, isProd bool, options ...Option) (client *Client, err error) {
	if apiKey == gopay.NULL || secretKey == gopay.NULL {
		return nil, gopay.MissParamErr
	}

	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)

	client = &Client{
		ApiKey:       apiKey,
		SecretKey:    secretKey,
		IsProd:       isProd,
		DebugSwitch:  DebugOff,
		logger:       logger,
		hc:           xhttp.NewClient(),
		baseUrlProd:  baseUrlProd,
		headerKeyMap: make(map[string]string),
	}

	for _, option := range options {
		option(client)
	}

	return client, nil
}

// WithProxyUrl 设置代理 URL
func WithProxyUrl(proxyUrlProd string) Option {
	return func(c *Client) {
		c.baseUrlProd = proxyUrlProd
	}
}

// WithHttpClient 设置自定义的xhttp.Client
func WithHttpClient(client *xhttp.Client) Option {
	return func(c *Client) {
		c.hc = client
	}
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}

// SetHttpClient 设置自定义的xhttp.Client
func (c *Client) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		c.hc = client
	}
}

// SetLogger 设置自定义的logger
func (c *Client) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		c.logger = logger
	}
}

// SetProxyUrl 设置代理 URL
func (c *Client) SetProxyUrl(proxyUrlProd string) {
	c.baseUrlProd = proxyUrlProd
}

// SetRequestHeader 设置自定义的header
func (c *Client) SetRequestHeader(key string, defaultVal ...string) {
	if key != "" {
		if len(defaultVal) > 0 {
			c.headerKeyMap[key] = defaultVal[0]
		} else {
			c.headerKeyMap[key] = ""
		}
	}
}

// ClearRequestHeader 清理自定义的header
func (c *Client) ClearRequestHeader() {
	c.headerKeyMap = make(map[string]string)
}

// GetBaseUrl 获取基础URL
func (c *Client) GetBaseUrl() string {
	return c.baseUrlProd
}

// doCreemGet 发送GET请求到Creem API
func (c *Client) doCreemGet(ctx context.Context, path string) (res *http.Response, bs []byte, err error) {
	url := c.GetBaseUrl() + path

	req := c.hc.Req()

	// 设置认证头 - Creem使用x-api-key
	req.Header.Set("x-api-key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// 设置自定义header
	for k, v := range c.headerKeyMap {
		req.Header.Set(k, v)
	}

	if c.DebugSwitch == DebugOn {
		c.logger.Debugf("Creem_Request: %s", url)
	}

	res, bs, err = req.Get(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("http.Do Error: %w", err)
	}

	if c.DebugSwitch == DebugOn {
		c.logger.Debugf("Creem_Response: %s", string(bs))
	}

	return res, bs, nil
}

// doCreemPost 发送POST请求到Creem API
func (c *Client) doCreemPost(ctx context.Context, body interface{}, path string) (res *http.Response, bs []byte, err error) {
	url := c.GetBaseUrl() + path

	req := c.hc.Req()

	// 设置认证头 - Creem使用x-api-key
	req.Header.Set("x-api-key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// 设置自定义header
	for k, v := range c.headerKeyMap {
		req.Header.Set(k, v)
	}

	if c.DebugSwitch == DebugOn {
		if body != nil {
			bodyBytes, _ := json.Marshal(body)
			c.logger.Debugf("Creem_Request: %s, Body: %s", url, string(bodyBytes))
		} else {
			c.logger.Debugf("Creem_Request: %s", url)
		}
	}

	if body != nil {
		res, bs, err = req.Post(url).SendStruct(body).EndBytes(ctx)
	} else {
		res, bs, err = req.Post(url).EndBytes(ctx)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("http.Do Error: %w", err)
	}

	if c.DebugSwitch == DebugOn {
		c.logger.Debugf("Creem_Response: %s", string(bs))
	}

	return res, bs, nil
}

// doCreemPut 发送PUT请求到Creem API
func (c *Client) doCreemPut(ctx context.Context, body interface{}, path string) (res *http.Response, bs []byte, err error) {
	url := c.GetBaseUrl() + path

	req := c.hc.Req()

	// 设置认证头 - Creem使用x-api-key
	req.Header.Set("x-api-key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// 设置自定义header
	for k, v := range c.headerKeyMap {
		req.Header.Set(k, v)
	}

	if c.DebugSwitch == DebugOn {
		if body != nil {
			bodyBytes, _ := json.Marshal(body)
			c.logger.Debugf("Creem_Request: %s, Body: %s", url, string(bodyBytes))
		} else {
			c.logger.Debugf("Creem_Request: %s", url)
		}
	}

	if body != nil {
		res, bs, err = req.Put(url).SendStruct(body).EndBytes(ctx)
	} else {
		res, bs, err = req.Put(url).EndBytes(ctx)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("http.Do Error: %w", err)
	}

	if c.DebugSwitch == DebugOn {
		c.logger.Debugf("Creem_Response: %s", string(bs))
	}

	return res, bs, nil
}

// doCreemDelete 发送DELETE请求到Creem API
func (c *Client) doCreemDelete(ctx context.Context, path string) (res *http.Response, bs []byte, err error) {
	url := c.GetBaseUrl() + path

	req := c.hc.Req()

	// 设置认证头 - Creem使用x-api-key
	req.Header.Set("x-api-key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// 设置自定义header
	for k, v := range c.headerKeyMap {
		req.Header.Set(k, v)
	}

	if c.DebugSwitch == DebugOn {
		c.logger.Debugf("Creem_Request: %s", url)
	}

	res, bs, err = req.Delete(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("http.Do Error: %w", err)
	}

	if c.DebugSwitch == DebugOn {
		c.logger.Debugf("Creem_Response: %s", string(bs))
	}

	return res, bs, nil
}
