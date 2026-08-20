package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/air-iot/api-client-go/v4/apicontext"
	"github.com/air-iot/api-client-go/v4/config"
	"github.com/air-iot/api-client-go/v4/warning"
	"github.com/air-iot/api-client-gojs/v4/utils"
	"github.com/dop251/goja"
)

// CreateWarning 新增报警记录
//
// projectId 项目ID
//
// warning 报警数据. 对象或对象数组
func (a Client) CreateWarning(ctx context.Context, vm *goja.Runtime, projectId string, warning goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}

	if !utils.IsObject(warning) && !utils.IsObjectArray(warning) {
		return Failure("warning 不是有效的对象或对象数组").ToObject(vm)
	}

	var warnings []any
	if utils.IsObject(warning) {
		warnings = append(warnings, warning.Export())
	} else if utils.IsObjectArray(warning) {
		if v, ok := utils.GetAsArrayObject(warning); ok {
			for i := range v {
				warnings = append(warnings, v[i])
			}
		} else {
			return Failure("warning 不是有效的对象或对象数组").ToObject(vm)
		}
	} else {
		return Failure("warning 不是有效的对象或对象数组").ToObject(vm)
	}

	var result map[string]any
	if err := a.cli.BatchCreateWarn(ctx, projectId, warnings, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// GetWarning 根据ID查询报警记录
//
// projectId 项目ID
//
// warningId 报警记录ID
//
// archive 是否查询归档记录
func (a Client) GetWarning(ctx context.Context, vm *goja.Runtime, projectId, warningId string, archive bool) goja.Value {
	projectId = strings.TrimSpace(projectId)
	warningId = strings.TrimSpace(warningId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if warningId == "" {
		return Failure("warningId 不能为空").ToObject(vm)
	}

	var result map[string]any
	if _, err := a.cli.GetWarn(ctx, projectId, strconv.FormatBool(archive), warningId, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// QueryWarning 查询报警记录
//
// projectId 项目ID
//
// query 报警记录ID
//
// archive 是否查询归档记录
func (a Client) QueryWarning(ctx context.Context, vm *goja.Runtime, projectId string, query goja.Value, archive bool) goja.Value {
	projectId = strings.TrimSpace(projectId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if !utils.IsObject(query) {
		return Failure("query 不是有效的对象").ToObject(vm)
	}

	bts, err := json.Marshal(query)
	if err != nil {
		return Failure(fmt.Sprintf("序列化查询参数错误, %+v", err)).ToObject(vm)
	}
	cli, err := a.cli.WarningClient.GetWarnServiceClient()
	if err != nil {
		return Failure(fmt.Sprintf("获取平台客户端失败, %+v", err)).ToObject(vm)
	}

	var result []map[string]any
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&warning.QueryWarningRequest{Query: bts, Archive: strconv.FormatBool(archive)})
	if _, err := ParseRes(err, res, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}

	return Success(result).WithCount(res.GetCount()).ToObject(vm)
}
