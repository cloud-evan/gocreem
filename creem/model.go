package creem

import (
	"time"
)

// 基础响应结构
type BaseResponse struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
}

// 错误响应
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// 产品相关模型
type Product struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Price       float64                `json:"price"`
	Currency    string                 `json:"currency"`
	Active      bool                   `json:"active"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type ProductCreateRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Price       float64                `json:"price"`
	Currency    string                 `json:"currency"`
	Active      bool                   `json:"active"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type ProductUpdateRequest struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Price       float64                `json:"price,omitempty"`
	Currency    string                 `json:"currency,omitempty"`
	Active      bool                   `json:"active,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type ProductsListResponse struct {
	BaseResponse
	Data       []Product `json:"data"`
	TotalCount int       `json:"total_count"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}

type ProductDetailResponse struct {
	BaseResponse
	Data Product `json:"data"`
}

type ProductCreateResponse struct {
	BaseResponse
	Data Product `json:"data"`
}

type ProductUpdateResponse struct {
	BaseResponse
	Data Product `json:"data"`
}

// 结账会话相关模型
type CheckoutSession struct {
	ID              string                 `json:"id"`
	ProductID       string                 `json:"product_id"`
	CustomerID      string                 `json:"customer_id,omitempty"`
	Status          string                 `json:"status"`
	Amount          float64                `json:"amount"`
	Currency        string                 `json:"currency"`
	ReturnURL       string                 `json:"return_url"`
	CancelURL       string                 `json:"cancel_url"`
	SuccessURL      string                 `json:"success_url"`
	PaymentMethodID string                 `json:"payment_method_id,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	ExpiresAt       time.Time              `json:"expires_at"`
}

type CheckoutSessionCreateRequest struct {
	ProductID       string                 `json:"product_id"`
	CustomerID      string                 `json:"customer_id,omitempty"`
	Amount          float64                `json:"amount,omitempty"`
	Currency        string                 `json:"currency,omitempty"`
	ReturnURL       string                 `json:"return_url"`
	CancelURL       string                 `json:"cancel_url"`
	SuccessURL      string                 `json:"success_url"`
	PaymentMethodID string                 `json:"payment_method_id,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

type CheckoutSessionUpdateRequest struct {
	Amount          float64                `json:"amount,omitempty"`
	Currency        string                 `json:"currency,omitempty"`
	ReturnURL       string                 `json:"return_url,omitempty"`
	CancelURL       string                 `json:"cancel_url,omitempty"`
	SuccessURL      string                 `json:"success_url,omitempty"`
	PaymentMethodID string                 `json:"payment_method_id,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

type CheckoutSessionResponse struct {
	BaseResponse
	Data CheckoutSession `json:"data"`
}

// 订单相关模型
type Order struct {
	ID                string                 `json:"id"`
	CheckoutSessionID string                 `json:"checkout_session_id"`
	ProductID         string                 `json:"product_id"`
	CustomerID        string                 `json:"customer_id"`
	Status            string                 `json:"status"`
	Amount            float64                `json:"amount"`
	Currency          string                 `json:"currency"`
	PaymentMethodID   string                 `json:"payment_method_id"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	PaidAt            *time.Time             `json:"paid_at,omitempty"`
}

type OrdersListResponse struct {
	BaseResponse
	Data       []Order `json:"data"`
	TotalCount int     `json:"total_count"`
	Page       int     `json:"page"`
	Limit      int     `json:"limit"`
}

type OrderDetailResponse struct {
	BaseResponse
	Data Order `json:"data"`
}

type OrderUpdateRequest struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type OrderUpdateResponse struct {
	BaseResponse
	Data Order `json:"data"`
}

// 订阅相关模型
type Subscription struct {
	ID                 string                 `json:"id"`
	CustomerID         string                 `json:"customer_id"`
	ProductID          string                 `json:"product_id"`
	Status             string                 `json:"status"`
	Amount             float64                `json:"amount"`
	Currency           string                 `json:"currency"`
	BillingCycle       string                 `json:"billing_cycle"`
	TrialDays          int                    `json:"trial_days"`
	CurrentPeriodStart time.Time              `json:"current_period_start"`
	CurrentPeriodEnd   time.Time              `json:"current_period_end"`
	CanceledAt         *time.Time             `json:"canceled_at,omitempty"`
	EndedAt            *time.Time             `json:"ended_at,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}

type SubscriptionCreateRequest struct {
	CustomerID   string                 `json:"customer_id"`
	ProductID    string                 `json:"product_id"`
	Amount       float64                `json:"amount,omitempty"`
	Currency     string                 `json:"currency,omitempty"`
	BillingCycle string                 `json:"billing_cycle"`
	TrialDays    int                    `json:"trial_days,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type SubscriptionUpdateRequest struct {
	Amount       float64                `json:"amount,omitempty"`
	Currency     string                 `json:"currency,omitempty"`
	BillingCycle string                 `json:"billing_cycle,omitempty"`
	TrialDays    int                    `json:"trial_days,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type SubscriptionsListResponse struct {
	BaseResponse
	Data       []Subscription `json:"data"`
	TotalCount int            `json:"total_count"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
}

type SubscriptionDetailResponse struct {
	BaseResponse
	Data Subscription `json:"data"`
}

type SubscriptionUpdateResponse struct {
	BaseResponse
	Data Subscription `json:"data"`
}

type SubscriptionUpgradeRequest struct {
	NewProductID string                 `json:"new_product_id"`
	Amount       float64                `json:"amount,omitempty"`
	Currency     string                 `json:"currency,omitempty"`
	BillingCycle string                 `json:"billing_cycle,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type SubscriptionUpgradeResponse struct {
	BaseResponse
	Data Subscription `json:"data"`
}

type SubscriptionCancelResponse struct {
	BaseResponse
	Data struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

// 客户相关模型
type Customer struct {
	ID        string                 `json:"id"`
	Email     string                 `json:"email"`
	Name      string                 `json:"name"`
	Phone     string                 `json:"phone,omitempty"`
	Company   string                 `json:"company,omitempty"`
	Address   *Address               `json:"address,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2,omitempty"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type CustomerCreateRequest struct {
	Email    string                 `json:"email"`
	Name     string                 `json:"name"`
	Phone    string                 `json:"phone,omitempty"`
	Company  string                 `json:"company,omitempty"`
	Address  *Address               `json:"address,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type CustomerUpdateRequest struct {
	Email    string                 `json:"email,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Phone    string                 `json:"phone,omitempty"`
	Company  string                 `json:"company,omitempty"`
	Address  *Address               `json:"address,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type CustomersListResponse struct {
	BaseResponse
	Data       []Customer `json:"data"`
	TotalCount int        `json:"total_count"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
}

type CustomerDetailResponse struct {
	BaseResponse
	Data Customer `json:"data"`
}

type CustomerCreateResponse struct {
	BaseResponse
	Data Customer `json:"data"`
}

type CustomerUpdateResponse struct {
	BaseResponse
	Data Customer `json:"data"`
}

// 客户门户相关模型
type CustomerPortalSession struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	URL        string    `json:"url"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

type CustomerPortalCreateRequest struct {
	CustomerID string `json:"customer_id"`
	ReturnURL  string `json:"return_url"`
}

type CustomerPortalCreateResponse struct {
	BaseResponse
	Data CustomerPortalSession `json:"data"`
}

// 支付方式相关模型
type PaymentMethod struct {
	ID         string                 `json:"id"`
	CustomerID string                 `json:"customer_id"`
	Type       string                 `json:"type"`
	Card       *Card                  `json:"card,omitempty"`
	Bank       *Bank                  `json:"bank,omitempty"`
	Default    bool                   `json:"default"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

type Card struct {
	Brand       string `json:"brand"`
	Last4       string `json:"last4"`
	ExpMonth    int    `json:"exp_month"`
	ExpYear     int    `json:"exp_year"`
	Fingerprint string `json:"fingerprint"`
	Country     string `json:"country,omitempty"`
}

type Bank struct {
	BankName      string `json:"bank_name"`
	Last4         string `json:"last4"`
	RoutingNumber string `json:"routing_number"`
	AccountType   string `json:"account_type"`
	Country       string `json:"country"`
}

type PaymentMethodCreateRequest struct {
	CustomerID string                 `json:"customer_id"`
	Type       string                 `json:"type"`
	Card       *CardCreateRequest     `json:"card,omitempty"`
	Bank       *BankCreateRequest     `json:"bank,omitempty"`
	Default    bool                   `json:"default,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type CardCreateRequest struct {
	Number   string   `json:"number"`
	ExpMonth int      `json:"exp_month"`
	ExpYear  int      `json:"exp_year"`
	Cvc      string   `json:"cvc"`
	Name     string   `json:"name,omitempty"`
	Address  *Address `json:"address,omitempty"`
}

type BankCreateRequest struct {
	AccountNumber     string `json:"account_number"`
	RoutingNumber     string `json:"routing_number"`
	AccountType       string `json:"account_type"`
	AccountHolderName string `json:"account_holder_name"`
	Country           string `json:"country"`
}

type PaymentMethodsListResponse struct {
	BaseResponse
	Data       []PaymentMethod `json:"data"`
	TotalCount int             `json:"total_count"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
}

type PaymentMethodDetailResponse struct {
	BaseResponse
	Data PaymentMethod `json:"data"`
}

type PaymentMethodCreateResponse struct {
	BaseResponse
	Data PaymentMethod `json:"data"`
}

type PaymentMethodUpdateRequest struct {
	Default  bool                   `json:"default,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type PaymentMethodUpdateResponse struct {
	BaseResponse
	Data PaymentMethod `json:"data"`
}

// Webhook相关模型
type Webhook struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Events    []string  `json:"events"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WebhookCreateRequest struct {
	URL    string   `json:"url"`
	Events []string `json:"events"`
}

type WebhookUpdateRequest struct {
	URL    string   `json:"url,omitempty"`
	Events []string `json:"events,omitempty"`
	Active bool     `json:"active,omitempty"`
}

type WebhooksListResponse struct {
	BaseResponse
	Data       []Webhook `json:"data"`
	TotalCount int       `json:"total_count"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}

type WebhookDetailResponse struct {
	BaseResponse
	Data Webhook `json:"data"`
}

type WebhookCreateResponse struct {
	BaseResponse
	Data Webhook `json:"data"`
}

type WebhookUpdateResponse struct {
	BaseResponse
	Data Webhook `json:"data"`
}

// 发票相关模型
type Invoice struct {
	ID             string                 `json:"id"`
	CustomerID     string                 `json:"customer_id"`
	OrderID        string                 `json:"order_id,omitempty"`
	SubscriptionID string                 `json:"subscription_id,omitempty"`
	Number         string                 `json:"number"`
	Status         string                 `json:"status"`
	Amount         float64                `json:"amount"`
	Currency       string                 `json:"currency"`
	TaxAmount      float64                `json:"tax_amount"`
	TotalAmount    float64                `json:"total_amount"`
	Items          []InvoiceItem          `json:"items"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
	DueDate        *time.Time             `json:"due_date,omitempty"`
	PaidAt         *time.Time             `json:"paid_at,omitempty"`
}

type InvoiceItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
	TaxAmount   float64 `json:"tax_amount"`
}

type InvoiceCreateRequest struct {
	CustomerID     string                 `json:"customer_id"`
	OrderID        string                 `json:"order_id,omitempty"`
	SubscriptionID string                 `json:"subscription_id,omitempty"`
	Amount         float64                `json:"amount"`
	Currency       string                 `json:"currency"`
	Items          []InvoiceItemRequest   `json:"items"`
	DueDate        *time.Time             `json:"due_date,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

type InvoiceItemRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

type InvoicesListResponse struct {
	BaseResponse
	Data       []Invoice `json:"data"`
	TotalCount int       `json:"total_count"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}

type InvoiceDetailResponse struct {
	BaseResponse
	Data Invoice `json:"data"`
}

type InvoiceCreateResponse struct {
	BaseResponse
	Data Invoice `json:"data"`
}

type InvoiceUpdateRequest struct {
	Amount   float64                `json:"amount,omitempty"`
	Currency string                 `json:"currency,omitempty"`
	Items    []InvoiceItemRequest   `json:"items,omitempty"`
	DueDate  *time.Time             `json:"due_date,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type InvoiceUpdateResponse struct {
	BaseResponse
	Data Invoice `json:"data"`
}

// 退款相关模型
type Refund struct {
	ID        string                 `json:"id"`
	OrderID   string                 `json:"order_id"`
	Amount    float64                `json:"amount"`
	Currency  string                 `json:"currency"`
	Status    string                 `json:"status"`
	Reason    string                 `json:"reason,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type RefundCreateRequest struct {
	OrderID  string                 `json:"order_id"`
	Amount   float64                `json:"amount,omitempty"`
	Currency string                 `json:"currency,omitempty"`
	Reason   string                 `json:"reason,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type RefundsListResponse struct {
	BaseResponse
	Data       []Refund `json:"data"`
	TotalCount int      `json:"total_count"`
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
}

type RefundDetailResponse struct {
	BaseResponse
	Data Refund `json:"data"`
}

type RefundCreateResponse struct {
	BaseResponse
	Data Refund `json:"data"`
}

type RefundUpdateRequest struct {
	Reason   string                 `json:"reason,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type RefundUpdateResponse struct {
	BaseResponse
	Data Refund `json:"data"`
}

// 账户相关模型
type Account struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Company   string                 `json:"company,omitempty"`
	Website   string                 `json:"website,omitempty"`
	LogoURL   string                 `json:"logo_url,omitempty"`
	Timezone  string                 `json:"timezone"`
	Locale    string                 `json:"locale"`
	Currency  string                 `json:"currency"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type AccountUpdateRequest struct {
	Name     string                 `json:"name,omitempty"`
	Email    string                 `json:"email,omitempty"`
	Company  string                 `json:"company,omitempty"`
	Website  string                 `json:"website,omitempty"`
	LogoURL  string                 `json:"logo_url,omitempty"`
	Timezone string                 `json:"timezone,omitempty"`
	Locale   string                 `json:"locale,omitempty"`
	Currency string                 `json:"currency,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type AccountDetailResponse struct {
	BaseResponse
	Data Account `json:"data"`
}

type AccountUpdateResponse struct {
	BaseResponse
	Data Account `json:"data"`
}

// 报告相关模型
type Report struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	Status      string     `json:"status"`
	URL         string     `json:"url,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type ReportCreateRequest struct {
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type ReportsListResponse struct {
	BaseResponse
	Data       []Report `json:"data"`
	TotalCount int      `json:"total_count"`
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
}

type ReportDetailResponse struct {
	BaseResponse
	Data Report `json:"data"`
}

type ReportCreateResponse struct {
	BaseResponse
	Data Report `json:"data"`
}

// 分页参数
type PaginationParams struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

// 列表查询参数
type ListParams struct {
	PaginationParams
	Status     string     `json:"status,omitempty"`
	CustomerID string     `json:"customer_id,omitempty"`
	ProductID  string     `json:"product_id,omitempty"`
	StartDate  *time.Time `json:"start_date,omitempty"`
	EndDate    *time.Time `json:"end_date,omitempty"`
}

// Transaction相关模型
type Transaction struct {
	ID                string                 `json:"id"`
	CheckoutSessionID string                 `json:"checkout_session_id"`
	ProductID         string                 `json:"product_id"`
	CustomerID        string                 `json:"customer_id"`
	Status            string                 `json:"status"`
	Amount            float64                `json:"amount"`
	Currency          string                 `json:"currency"`
	PaymentMethodID   string                 `json:"payment_method_id"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	PaidAt            *time.Time             `json:"paid_at,omitempty"`
}

type TransactionsListResponse struct {
	BaseResponse
	Data       []Transaction `json:"data"`
	TotalCount int           `json:"total_count"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
}

// License相关模型
type LicenseValidateRequest struct {
	LicenseKey string `json:"license_key"`
}

type LicenseValidateResponse struct {
	BaseResponse
	Data struct {
		Valid    bool   `json:"valid"`
		Message  string `json:"message"`
		Customer string `json:"customer,omitempty"`
	} `json:"data"`
}

type LicenseActivateRequest struct {
	LicenseKey string `json:"license_key"`
	CustomerID string `json:"customer_id"`
}

type LicenseActivateResponse struct {
	BaseResponse
	Data struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

type LicenseDeactivateRequest struct {
	LicenseKey string `json:"license_key"`
}

type LicenseDeactivateResponse struct {
	BaseResponse
	Data struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

// Discount Code相关模型
type DiscountCode struct {
	ID         string                 `json:"id"`
	Code       string                 `json:"code"`
	Type       string                 `json:"type"` // percentage, fixed_amount
	Value      float64                `json:"value"`
	MaxUses    int                    `json:"max_uses"`
	UsedCount  int                    `json:"used_count"`
	Active     bool                   `json:"active"`
	ValidFrom  time.Time              `json:"valid_from"`
	ValidUntil time.Time              `json:"valid_until"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

type DiscountCodeCreateRequest struct {
	Code       string                 `json:"code"`
	Type       string                 `json:"type"`
	Value      float64                `json:"value"`
	MaxUses    int                    `json:"max_uses,omitempty"`
	ValidFrom  time.Time              `json:"valid_from,omitempty"`
	ValidUntil time.Time              `json:"valid_until,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type DiscountCodeCreateResponse struct {
	BaseResponse
	Data DiscountCode `json:"data"`
}

type DiscountCodeDetailResponse struct {
	BaseResponse
	Data DiscountCode `json:"data"`
}
