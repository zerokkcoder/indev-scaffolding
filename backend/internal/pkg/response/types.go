package response

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`              // 状态码
	Message string      `json:"message"`           // 消息
	Data    interface{} `json:"data,omitempty"`    // 数据
	TraceID string      `json:"trace_id,omitempty"` // 追踪ID
}

// PageResult 分页结果
type PageResult struct {
	List     interface{} `json:"list"`      // 数据列表
	Total    int64      `json:"total"`      // 总数
	Page     int        `json:"page"`       // 当前页码
	PageSize int        `json:"page_size"`  // 每页数量
}

// ListResult 列表结果
type ListResult struct {
	List  interface{} `json:"list"`  // 数据列表
	Total int64      `json:"total"`  // 总数
}

// StatusCodes 状态码定义
const (
	StatusOK                  = 200   // 成功
	StatusBadRequest          = 400   // 请求错误
	StatusUnauthorized        = 401   // 未授权
	StatusForbidden           = 403   // 禁止访问
	StatusNotFound           = 404   // 资源不存在
	StatusMethodNotAllowed    = 405   // 方法不允许
	StatusTooManyRequests    = 429   // 请求过多
	StatusInternalServerError = 500   // 服务器错误
	StatusServiceUnavailable  = 503   // 服务不可用
)
