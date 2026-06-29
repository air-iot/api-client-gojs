package apiclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/air-iot/api-client-gojs/v4/utils"
	"github.com/air-iot/gojs"
	"github.com/dop251/goja"
)

// CallDataInterface 调用数据接口
//
// projectId: 项目ID
//
// dsId: 接口标识
//
// params: 接口参数
func (a Client) CallDataInterface(ctx context.Context, vm *goja.Runtime, projectId, dsId string, params goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	dsId = strings.TrimSpace(dsId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if dsId == "" {
		return Failure("dsId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(params) {
		return Failure("params 不能为空").ToObject(vm)
	} else if !utils.IsObject(params) {
		return Failure("params 不是有效的对象").ToObject(vm)
	}

	paramObj, ok := params.Export().(map[string]any)
	if !ok {
		return Failure("params 不是有效的对象, 与 map[string]any 类型不一致").ToObject(vm)
	}

	result, err := a.cli.DataInterfaceProxy(ctx, projectId, dsId, paramObj)
	if err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	}

	return Success(result).ToObject(vm)
}
