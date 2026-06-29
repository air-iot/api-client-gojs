package apiclient

type InsertResult struct {
	InsertedID string `json:"InsertedID"`
}

type InsertsResult struct {
	InsertedIDs []string `json:"InsertedIDs"`
}

type StatusResult struct {
	// 状态(OK)
	Status string `json:"status"`
}

// ErrorResult 响应错误
type ErrorResult struct {
	// 错误码
	Code int `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 错误字段
	Field string `json:"field"`
	// 错误详情信息
	Detail string `json:"detail"`
}
