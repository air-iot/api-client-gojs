package apiclient

import (
	"encoding/json"

	"github.com/air-iot/api-client-go/v4/api"
	internalError "github.com/air-iot/api-client-go/v4/errors"
	"github.com/air-iot/errors"
	"github.com/dop251/goja"
)

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Count   *int64 `json:"count,omitempty"`
	Data    any    `json:"data"`
}

func Success(data any) *Result {
	return &Result{Success: true, Message: "OK", Data: data}
}

func SuccessEmpty() *Result {
	return &Result{Success: true}
}

func Failure(message string) *Result {
	return &Result{Success: false, Message: message}
}

func (r *Result) WithCount(count int64) *Result {
	r.Count = &count
	return r
}

func (r *Result) ToObject(vm *goja.Runtime) *goja.Object {
	obj := vm.NewObject()
	_ = obj.Set("success", r.Success)
	_ = obj.Set("message", r.Message)
	_ = obj.Set("data", r.Data)
	if r.Count != nil {
		_ = obj.Set("count", r.Count)
	}
	return obj
}

func ParseRes(err error, res *api.Response, result any) ([]byte, error) {
	if err != nil {
		return nil, errors.NewResErrorMsg(err, "请求错误")
	}
	if !res.GetStatus() {
		return nil, internalError.ParseResponse(res)
	}
	if result != nil && res.GetResult() != nil {
		if err := json.Unmarshal(res.GetResult(), result); err != nil {
			return nil, errors.Wrap(err, "解析请求结果错误")
		}
	}
	return res.GetResult(), nil
}
