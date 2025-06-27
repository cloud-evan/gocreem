# Creem Go SDK

基于 [Creem API 文档](https://docs.creem.io/api-reference/introduction) 实现的 Go 语言 SDK，严格按照 gopay 的设计模式封装。

## 特性

- 完整的 Creem API 接口封装
- 统一的错误处理和响应解析
- 支持自定义 HTTP 客户端和代理
- 详细的参数校验
- 完整的单元测试

## 安装

```bash
go get github.com/your-repo/creem
```

## 快速开始

### 初始化客户端

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/your-repo/creem"
)

func main() {
    // 初始化客户端
    client, err := creem.NewClient("your_api_key", "your_secret_key", true)
    if err != nil {
        log.Fatal(err)
    }
    
    // 启用调试模式
    client.DebugSwitch = creem.DebugOn
    
    ctx := context.Background()
    
    // 使用示例...
}
```

## API 接口

### Checkout（结账会话）

#### 创建结账会话

```go
req := &creem.CheckoutSessionCreateRequest{
    ProductID:  "prod_123",
    CustomerID: "cust_123",
    ReturnURL:  "https://example.com/return",
    CancelURL:  "https://example.com/cancel",
    SuccessURL: "https://example.com/success",
}

rsp, err := client.CreateCheckoutSession(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Checkout Session ID: %s\n", rsp.Data.ID)
```

#### 获取结账会话详情

```go
rsp, err := client.GetCheckoutSession(ctx, "cs_123")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Status: %s\n", rsp.Data.Status)
```

### Product（产品）

#### 创建产品

```go
req := &creem.ProductCreateRequest{
    Name:        "Premium Plan",
    Description: "Premium subscription plan",
    Type:        creem.ProductTypeRecurring,
    Price:       29.99,
    Currency:    creem.CurrencyUSD,
    Active:      true,
}

rsp, err := client.CreateProduct(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Product ID: %s\n", rsp.Data.ID)
```

#### 获取产品详情

```go
rsp, err := client.GetProduct(ctx, "prod_123")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Product Name: %s\n", rsp.Data.Name)
```

#### 获取产品列表

```go
params := &creem.ListParams{
    Page:  1,
    Limit: 10,
}

rsp, err := client.ListProducts(ctx, params)
if err != nil {
    log.Fatal(err)
}

for _, product := range rsp.Data {
    fmt.Printf("Product: %s - %s\n", product.ID, product.Name)
}
```

### Customer（客户）

#### 获取客户详情

```go
rsp, err := client.GetCustomer(ctx, "cust_123")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Customer Email: %s\n", rsp.Data.Email)
```

### Transactions（交易）

#### 获取交易列表

```go
params := &creem.ListParams{
    Page:       1,
    Limit:      10,
    CustomerID: "cust_123",
    StartDate:  &startDate,
    EndDate:    &endDate,
}

rsp, err := client.ListTransactions(ctx, params)
if err != nil {
    log.Fatal(err)
}

for _, transaction := range rsp.Data {
    fmt.Printf("Transaction: %s - %f %s\n", transaction.ID, transaction.Amount, transaction.Currency)
}
```

### License（授权）

#### 校验授权密钥

```go
req := &creem.LicenseValidateRequest{
    LicenseKey: "license_key_123",
}

rsp, err := client.ValidateLicense(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Valid: %t\n", rsp.Data.Valid)
```

#### 激活授权密钥

```go
req := &creem.LicenseActivateRequest{
    LicenseKey: "license_key_123",
    CustomerID: "cust_123",
}

rsp, err := client.ActivateLicense(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Success: %t\n", rsp.Data.Success)
```

#### 注销授权密钥

```go
req := &creem.LicenseDeactivateRequest{
    LicenseKey: "license_key_123",
}

rsp, err := client.DeactivateLicense(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Success: %t\n", rsp.Data.Success)
```

### Discount Code（优惠码）

#### 创建优惠码

```go
req := &creem.DiscountCodeCreateRequest{
    Code:  "SAVE20",
    Type:  "percentage",
    Value: 20.0,
    MaxUses: 100,
}

rsp, err := client.CreateDiscountCode(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Discount Code ID: %s\n", rsp.Data.ID)
```

#### 获取优惠码详情

```go
rsp, err := client.GetDiscountCode(ctx, "dc_123")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Code: %s, Type: %s\n", rsp.Data.Code, rsp.Data.Type)
```

#### 删除优惠码

```go
rsp, err := client.DeleteDiscountCode(ctx, "dc_123")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Deleted successfully\n")
```

### Subscription（订阅）

#### 获取订阅详情

```go
rsp, err := client.GetSubscription(ctx, "sub_123")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Subscription Status: %s\n", rsp.Data.Status)
```

#### 更新订阅

```go
req := &creem.SubscriptionUpdateRequest{
    Amount:       39.99,
    Currency:     creem.CurrencyUSD,
    BillingCycle: creem.BillingCycleMonthly,
}

rsp, err := client.UpdateSubscription(ctx, "sub_123", req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Updated successfully\n")
```

#### 升级订阅

```go
req := &creem.SubscriptionUpgradeRequest{
    NewProductID: "prod_456",
    Amount:       49.99,
    Currency:     creem.CurrencyUSD,
}

rsp, err := client.UpgradeSubscription(ctx, "sub_123", req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Upgraded successfully\n")
```

#### 取消订阅

```go
rsp, err := client.CancelSubscription(ctx, "sub_123")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Success: %t\n", rsp.Data.Success)
```

## 配置选项

### 自定义 HTTP 客户端

```go
client, err := creem.NewClient("api_key", "secret_key", true,
    creem.WithHttpClient(customHttpClient),
)
```

### 设置代理

```go
client, err := creem.NewClient("api_key", "secret_key", true,
    creem.WithProxyUrl("https://proxy.example.com"),
)
```

### 设置自定义请求头

```go
client.SetRequestHeader("X-Custom-Header", "custom_value")
```

### 设置响应体大小限制

```go
client.SetBodySize(10) // 10MB
```

## 错误处理

所有 API 方法都会返回统一的错误格式：

```go
rsp, err := client.GetProduct(ctx, "prod_123")
if err != nil {
    log.Fatal(err)
}

if rsp.Code != creem.Success {
    fmt.Printf("Error: %s\n", rsp.Error)
    if rsp.ErrorResponse != nil {
        fmt.Printf("Error Message: %s\n", rsp.ErrorResponse.Message)
    }
}
```

## 状态码

Creem API 使用标准 HTTP 状态码：

- `200` - 成功
- `400` - 参数错误
- `401` - API 密钥缺失
- `403` - API 密钥无效
- `404` - 资源未找到
- `429` - 请求频率超限
- `500` - 服务器错误

## 常量定义

### 货币代码

```go
creem.CurrencyUSD // USD
creem.CurrencyEUR // EUR
creem.CurrencyGBP // GBP
creem.CurrencyJPY // JPY
creem.CurrencyCAD // CAD
creem.CurrencyAUD // AUD
creem.CurrencyCHF // CHF
creem.CurrencyCNY // CNY
```

### 产品类型

```go
creem.ProductTypeOneTime   // one_time
creem.ProductTypeRecurring // recurring
creem.ProductTypeService   // service
creem.ProductTypeDigital   // digital
creem.ProductTypePhysical  // physical
```

### 计费周期

```go
creem.BillingCycleDaily   // daily
creem.BillingCycleWeekly  // weekly
creem.BillingCycleMonthly // monthly
creem.BillingCycleYearly  // yearly
```

### 状态常量

```go
creem.StatusActive    // active
creem.StatusInactive  // inactive
creem.StatusPending   // pending
creem.StatusCanceled  // canceled
creem.StatusCompleted // completed
creem.StatusFailed    // failed
```

## 测试

运行测试：

```bash
go test ./...
```

## 许可证

MIT License

## 参考文档

- [Creem API Reference](https://docs.creem.io/api-reference/introduction)
- [gopay](https://github.com/go-pay/gopay) - 参考设计模式 