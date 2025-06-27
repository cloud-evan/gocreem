package creem

import (
	"testing"
	"time"

	"github.com/go-pay/gopay"
)

func TestNewClient(t *testing.T) {
	// 测试正常初始化
	client, err := NewClient("test_api_key", "test_secret_key", false)
	if err != nil {
		t.Errorf("NewClient failed: %v", err)
	}

	if client.ApiKey != "test_api_key" {
		t.Errorf("Expected ApiKey to be 'test_api_key', got '%s'", client.ApiKey)
	}

	if client.SecretKey != "test_secret_key" {
		t.Errorf("Expected SecretKey to be 'test_secret_key', got '%s'", client.SecretKey)
	}

	if client.IsProd != false {
		t.Errorf("Expected IsProd to be false, got %v", client.IsProd)
	}

	// 测试空参数
	_, err = NewClient("", "test_secret_key", false)
	if err == nil {
		t.Error("Expected error for empty apiKey")
	}

	_, err = NewClient("test_api_key", "", false)
	if err == nil {
		t.Error("Expected error for empty secretKey")
	}
}

func TestClientOptions(t *testing.T) {
	// 测试自定义HTTP客户端
	customClient := gopay.NewClient()
	client, err := NewClient("test_api_key", "test_secret_key", false, WithHttpClient(customClient))
	if err != nil {
		t.Errorf("NewClient with custom HTTP client failed: %v", err)
	}

	if client.hc != customClient {
		t.Error("Custom HTTP client not set correctly")
	}

	// 测试代理URL
	client, err = NewClient("test_api_key", "test_secret_key", false,
		WithProxyUrl("https://proxy.prod.com", "https://proxy.sandbox.com"))
	if err != nil {
		t.Errorf("NewClient with proxy URL failed: %v", err)
	}

	if client.baseUrlProd != "https://proxy.prod.com" {
		t.Errorf("Expected prod proxy URL to be 'https://proxy.prod.com', got '%s'", client.baseUrlProd)
	}

	if client.baseUrlSandbox != "https://proxy.sandbox.com" {
		t.Errorf("Expected sandbox proxy URL to be 'https://proxy.sandbox.com', got '%s'", client.baseUrlSandbox)
	}
}

func TestGetBaseUrl(t *testing.T) {
	client, _ := NewClient("test_api_key", "test_secret_key", false)

	// 测试沙箱环境
	url := client.GetBaseUrl()
	if url != baseUrlSandbox {
		t.Errorf("Expected sandbox URL, got '%s'", url)
	}

	// 测试生产环境
	client.IsProd = true
	url = client.GetBaseUrl()
	if url != baseUrlProd {
		t.Errorf("Expected prod URL, got '%s'", url)
	}
}

func TestSetRequestHeader(t *testing.T) {
	client, _ := NewClient("test_api_key", "test_secret_key", false)

	// 测试设置header
	client.SetRequestHeader("X-Custom-Header", "custom_value")
	if client.headerKeyMap["X-Custom-Header"] != "custom_value" {
		t.Error("Header not set correctly")
	}

	// 测试设置空值header
	client.SetRequestHeader("X-Empty-Header")
	if client.headerKeyMap["X-Empty-Header"] != "" {
		t.Error("Empty header not set correctly")
	}

	// 测试清理header
	client.ClearRequestHeader()
	if len(client.headerKeyMap) != 0 {
		t.Error("Headers not cleared correctly")
	}
}

func TestProductCreateRequest(t *testing.T) {
	req := &ProductCreateRequest{
		Name:        "Test Product",
		Description: "A test product",
		Type:        ProductTypeOneTime,
		Price:       9.99,
		Currency:    CurrencyUSD,
		Active:      true,
	}

	if req.Name == "" {
		t.Error("Product name should not be empty")
	}

	if req.Description == "" {
		t.Error("Product description should not be empty")
	}

	if req.Price <= 0 {
		t.Error("Product price should be greater than 0")
	}

	if req.Currency == "" {
		t.Error("Product currency should not be empty")
	}
}

func TestCheckoutSessionCreateRequest(t *testing.T) {
	req := &CheckoutSessionCreateRequest{
		ProductID:  "prod_123",
		CustomerID: "cust_123",
		ReturnURL:  "https://example.com/return",
		CancelURL:  "https://example.com/cancel",
		SuccessURL: "https://example.com/success",
	}

	if req.ProductID == "" {
		t.Error("Product ID should not be empty")
	}

	if req.ReturnURL == "" {
		t.Error("Return URL should not be empty")
	}

	if req.CancelURL == "" {
		t.Error("Cancel URL should not be empty")
	}

	if req.SuccessURL == "" {
		t.Error("Success URL should not be empty")
	}
}

func TestCustomerCreateRequest(t *testing.T) {
	req := &CustomerCreateRequest{
		Email:   "test@example.com",
		Name:    "Test Customer",
		Phone:   "+1234567890",
		Company: "Test Company",
		Address: &Address{
			Line1:      "123 Test St",
			City:       "Test City",
			State:      "Test State",
			PostalCode: "12345",
			Country:    "US",
		},
	}

	if req.Email == "" {
		t.Error("Customer email should not be empty")
	}

	if req.Name == "" {
		t.Error("Customer name should not be empty")
	}

	if req.Address == nil {
		t.Error("Customer address should not be nil")
	}

	if req.Address.Line1 == "" {
		t.Error("Address line1 should not be empty")
	}
}

func TestSubscriptionCreateRequest(t *testing.T) {
	req := &SubscriptionCreateRequest{
		CustomerID:   "cust_123",
		ProductID:    "prod_123",
		Amount:       19.99,
		Currency:     CurrencyUSD,
		BillingCycle: BillingCycleMonthly,
		TrialDays:    7,
	}

	if req.CustomerID == "" {
		t.Error("Customer ID should not be empty")
	}

	if req.ProductID == "" {
		t.Error("Product ID should not be empty")
	}

	if req.BillingCycle == "" {
		t.Error("Billing cycle should not be empty")
	}

	if req.Amount <= 0 {
		t.Error("Amount should be greater than 0")
	}
}

func TestWebhookCreateRequest(t *testing.T) {
	req := &WebhookCreateRequest{
		URL:    "https://example.com/webhook",
		Events: []string{"checkout.session.completed", "subscription.created"},
	}

	if req.URL == "" {
		t.Error("Webhook URL should not be empty")
	}

	if len(req.Events) == 0 {
		t.Error("Webhook events should not be empty")
	}
}

func TestListParams(t *testing.T) {
	params := &ListParams{
		PaginationParams: PaginationParams{
			Page:  1,
			Limit: 10,
		},
		Status:     StatusActive,
		CustomerID: "cust_123",
		ProductID:  "prod_123",
		StartDate:  &time.Time{},
		EndDate:    &time.Time{},
	}

	if params.Page != 1 {
		t.Error("Page should be 1")
	}

	if params.Limit != 10 {
		t.Error("Limit should be 10")
	}

	if params.Status != StatusActive {
		t.Error("Status should be active")
	}

	if params.CustomerID != "cust_123" {
		t.Error("Customer ID should be cust_123")
	}

	if params.ProductID != "prod_123" {
		t.Error("Product ID should be prod_123")
	}
}

func TestConstants(t *testing.T) {
	// 测试状态常量
	if StatusActive != "active" {
		t.Error("StatusActive should be 'active'")
	}

	if StatusPending != "pending" {
		t.Error("StatusPending should be 'pending'")
	}

	if StatusCanceled != "canceled" {
		t.Error("StatusCanceled should be 'canceled'")
	}

	// 测试货币常量
	if CurrencyUSD != "USD" {
		t.Error("CurrencyUSD should be 'USD'")
	}

	if CurrencyEUR != "EUR" {
		t.Error("CurrencyEUR should be 'EUR'")
	}

	// 测试计费周期常量
	if BillingCycleMonthly != "monthly" {
		t.Error("BillingCycleMonthly should be 'monthly'")
	}

	if BillingCycleYearly != "yearly" {
		t.Error("BillingCycleYearly should be 'yearly'")
	}

	// 测试产品类型常量
	if ProductTypeOneTime != "one_time" {
		t.Error("ProductTypeOneTime should be 'one_time'")
	}

	if ProductTypeRecurring != "recurring" {
		t.Error("ProductTypeRecurring should be 'recurring'")
	}
}

func TestErrorConstants(t *testing.T) {
	// 测试错误常量
	if MissCreemInitParamErr.Error() != "missing creem init parameter" {
		t.Error("MissCreemInitParamErr message incorrect")
	}

	if MissProductIdErr.Error() != "missing product id" {
		t.Error("MissProductIdErr message incorrect")
	}

	if MissCustomerIdErr.Error() != "missing customer id" {
		t.Error("MissCustomerIdErr message incorrect")
	}

	if MissEmailErr.Error() != "missing email" {
		t.Error("MissEmailErr message incorrect")
	}

	if MissNameErr.Error() != "missing name" {
		t.Error("MissNameErr message incorrect")
	}
}
