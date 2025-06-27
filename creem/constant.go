package creem

const (
	NULL = ""
	// SUCCESS  = "SUCCESS"
	// FAIL     = "FAIL"
	// OK       = "OK"
	DebugOff = 0
	DebugOn  = 1
	Version  = "v1.0.0"
	Success  = 0

	HeaderApiKey = "x-api-key"            // Creem API认证头
	baseUrlProd  = "https://api.creem.io" // 正式 URL

	// Checkout相关
	checkoutSessionCreate = "/v1/checkout-sessions"    // 创建结账会话 POST
	checkoutSessionDetail = "/v1/checkout-sessions/%s" // session_id 获取结账会话 GET

	// Product相关
	productCreate = "/v1/products"    // 创建产品 POST
	productDetail = "/v1/products/%s" // product_id 获取产品 GET
	productsList  = "/v1/products"    // 获取产品列表 GET

	// Customer相关
	customerCreate = "/v1/customers"    // 创建客户 POST
	customerDetail = "/v1/customers/%s" // customer_id 获取客户 GET
	customerUpdate = "/v1/customers/%s" // customer_id 更新客户 PUT
	customerDelete = "/v1/customers/%s" // customer_id 删除客户 DELETE
	customersList  = "/v1/customers"    // 获取客户列表 GET

	// Customer Portal相关
	customerPortalCreate = "/v1/customer-portal/sessions" // 创建客户门户会话 POST

	// Transactions相关
	transactionsList = "/v1/transactions" // 获取交易列表 GET

	// License相关
	licenseValidate   = "/v1/licenses/validate"   // 校验授权密钥 POST
	licenseActivate   = "/v1/licenses/activate"   // 激活授权密钥 POST
	licenseDeactivate = "/v1/licenses/deactivate" // 注销授权密钥 POST

	// Discount Code相关
	discountCodeCreate = "/v1/discount-codes"    // 创建优惠码 POST
	discountCodeDelete = "/v1/discount-codes/%s" // discount_code_id 删除优惠码 DELETE
	discountCodeDetail = "/v1/discount-codes/%s" // discount_code_id 获取优惠码 GET

	// Subscription相关
	subscriptionDetail  = "/v1/subscriptions/%s"         // subscription_id 获取订阅 GET
	subscriptionUpdate  = "/v1/subscriptions/%s"         // subscription_id 更新订阅 POST
	subscriptionUpgrade = "/v1/subscriptions/%s/upgrade" // subscription_id 升级订阅 POST
	subscriptionCancel  = "/v1/subscriptions/%s/cancel"  // subscription_id 取消订阅 POST

	// 状态常量
	StatusActive    = "active"
	StatusInactive  = "inactive"
	StatusPending   = "pending"
	StatusCanceled  = "canceled"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusRefunded  = "refunded"
	StatusPaused    = "paused"
	StatusResumed   = "resumed"

	// 支付状态
	PaymentStatusPending   = "pending"
	PaymentStatusSucceeded = "succeeded"
	PaymentStatusFailed    = "failed"
	PaymentStatusCanceled  = "canceled"
	PaymentStatusRefunded  = "refunded"

	// 订阅状态
	SubscriptionStatusActive   = "active"
	SubscriptionStatusCanceled = "canceled"
	SubscriptionStatusPaused   = "paused"
	SubscriptionStatusPastDue  = "past_due"
	SubscriptionStatusUnpaid   = "unpaid"
	SubscriptionStatusTrialing = "trialing"

	// 订单状态
	OrderStatusPending   = "pending"
	OrderStatusCompleted = "completed"
	OrderStatusCanceled  = "canceled"
	OrderStatusRefunded  = "refunded"
	OrderStatusFailed    = "failed"

	// 发票状态
	InvoiceStatusDraft         = "draft"
	InvoiceStatusOpen          = "open"
	InvoiceStatusPaid          = "paid"
	InvoiceStatusVoid          = "void"
	InvoiceStatusUncollectible = "uncollectible"

	// 退款状态
	RefundStatusPending   = "pending"
	RefundStatusSucceeded = "succeeded"
	RefundStatusFailed    = "failed"
	RefundStatusCanceled  = "canceled"

	// 货币代码
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
	CurrencyGBP = "GBP"
	CurrencyJPY = "JPY"
	CurrencyCAD = "CAD"
	CurrencyAUD = "AUD"
	CurrencyCHF = "CHF"
	CurrencyCNY = "CNY"

	// 支付方式类型
	PaymentMethodTypeCard      = "card"
	PaymentMethodTypeBank      = "bank"
	PaymentMethodTypePaypal    = "paypal"
	PaymentMethodTypeApplePay  = "apple_pay"
	PaymentMethodTypeGooglePay = "google_pay"

	// 计费周期
	BillingCycleDaily   = "daily"
	BillingCycleWeekly  = "weekly"
	BillingCycleMonthly = "monthly"
	BillingCycleYearly  = "yearly"

	// 产品类型
	ProductTypeOneTime   = "one_time"
	ProductTypeRecurring = "recurring"
	ProductTypeService   = "service"
	ProductTypeDigital   = "digital"
	ProductTypePhysical  = "physical"
)

type DebugSwitch int8
