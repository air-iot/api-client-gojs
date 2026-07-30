package apiclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/air-iot/api-client-gojs/v4/utils"
	"github.com/air-iot/gojs"
	"github.com/dop251/goja"
)

// GetSystemVariable 根据系统变量的ID查询变量数据
//
// projectId 项目ID
//
// varId 系统变量的ID
func (a Client) GetSystemVariable(ctx context.Context, vm *goja.Runtime, projectId, varId string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	varId = strings.TrimSpace(varId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if varId == "" {
		return Failure("varId 不能为空").ToObject(vm)
	}

	var result map[string]any
	if _, err := a.cli.GetSystemVariable(ctx, projectId, varId, &result); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// GetSystemVariableByName 根据系统变量的名称查询变量数据
//
// projectId 项目ID
//
// varName 系统变量的名称
func (a Client) GetSystemVariableByName(ctx context.Context, vm *goja.Runtime, projectId, varName string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	varName = strings.TrimSpace(varName)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if varName == "" {
		return Failure("varName 不能为空").ToObject(vm)
	}

	var results []map[string]any
	if err := a.cli.QuerySystemVariable(ctx, projectId, map[string]any{
		"project": map[string]any{
			"id":    1,
			"uid":   1,
			"name":  1,
			"type":  1,
			"value": 1,
		},
		"filter": map[string]any{
			"name": varName,
		},
	}, &results); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	} else if len(results) == 0 {
		return Failure(fmt.Sprintf("未根据名称 '%s' 查询到系统变量", varName)).ToObject(vm)
	}
	return Success(results[0]).ToObject(vm)
}

// GetSystemVariableByUid 根据系统变量的 uid 查询变量数据
//
// projectId 项目ID
//
// varUid 系统变量的名称
func (a Client) GetSystemVariableByUid(ctx context.Context, vm *goja.Runtime, projectId, varUid string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	varUid = strings.TrimSpace(varUid)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if varUid == "" {
		return Failure("varUid 不能为空").ToObject(vm)
	}

	var results []map[string]any
	if err := a.cli.QuerySystemVariable(ctx, projectId, map[string]any{
		"project": map[string]any{
			"id":    1,
			"uid":   1,
			"name":  1,
			"type":  1,
			"value": 1,
		},
		"filter": map[string]any{
			"uid": varUid,
		},
	}, &results); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	} else if len(results) == 0 {
		return Failure(fmt.Sprintf("未根据uid '%s' 查询到系统变量", varUid)).ToObject(vm)
	}
	return Success(results[0]).ToObject(vm)
}

// QuerySystemVariable 自定义查询系统变量
//
// projectId 项目ID
//
// query 查询条件
func (a Client) QuerySystemVariable(ctx context.Context, vm *goja.Runtime, projectId string, query goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(query) {
		return Failure("无效的查询参数").ToObject(vm)
	} else if !utils.IsObject(query) {
		return Failure("无效的查询参数, 不是有效的对象").ToObject(vm)
	}

	var result map[string]any
	if err := a.cli.QuerySystemVariable(ctx, projectId, query.Export(), &result); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// CreateSystemVariable 新增系统变量数据
//
// projectId 项目ID
//
// variable 系统变量数据
func (a Client) CreateSystemVariable(ctx context.Context, vm *goja.Runtime, projectId string, variable goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(variable) {
		return Failure("variable 不能为空").ToObject(vm)
	} else if !utils.IsObject(variable) {
		return Failure("variable 有是有效的对象").ToObject(vm)
	}

	var result map[string]any
	if err := a.cli.CreateSystemVariable(ctx, projectId, variable.Export(), &result); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// UpdateSystemVariable 更新系统变量数据
//
// projectId 项目ID
//
// variable 系统变量数据
func (a Client) UpdateSystemVariable(ctx context.Context, vm *goja.Runtime, projectId, varId string, variable goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	varId = strings.TrimSpace(varId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if varId == "" {
		return Failure("varId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(variable) {
		return Failure("variable 不能为空").ToObject(vm)
	} else if !utils.IsObject(variable) {
		return Failure("variable 有是有效的对象").ToObject(vm)
	}

	var result StatusResult
	if err := a.cli.UpdateSystemVariable(ctx, projectId, varId, variable.Export(), &result); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// ReplaceSystemVariable 替换系统变量数据
//
// projectId 项目ID
//
// varId 系统变量的ID
//
// variable 系统变量数据
func (a Client) ReplaceSystemVariable(ctx context.Context, vm *goja.Runtime, projectId, varId string, variable goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	varId = strings.TrimSpace(varId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if varId == "" {
		return Failure("varId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(variable) {
		return Failure("variable 不能为空").ToObject(vm)
	} else if !utils.IsObject(variable) {
		return Failure("variable 有是有效的对象").ToObject(vm)
	}

	var result StatusResult
	if err := a.cli.ReplaceSystemVariable(ctx, projectId, varId, variable.Export(), &result); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// DeleteSystemVariable 删除系统变量数据
//
// projectId 项目ID
//
// varId 系统变量的ID
//
// variable 系统变量数据
func (a Client) DeleteSystemVariable(ctx context.Context, vm *goja.Runtime, projectId, varId string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	varId = strings.TrimSpace(varId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if varId == "" {
		return Failure("varId 不能为空").ToObject(vm)
	}

	var result StatusResult
	if err := a.cli.DeleteSystemVariable(ctx, projectId, varId, &result); err != nil {
		return Failure(fmt.Sprintf("调用数据接口失败, %+v", err)).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}
