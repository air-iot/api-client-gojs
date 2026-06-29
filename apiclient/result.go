package apiclient

import (
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
