package apiclient

import (
	"context"
	"strings"

	"github.com/air-iot/api-client-gojs/utils"
	"github.com/air-iot/gojs"
	"github.com/dop251/goja"
)

// GetRole 根据角色ID查询角色信息
//
// projectId: 项目ID
//
// roleId: 角色ID
func (a Client) GetRole(ctx context.Context, vm *goja.Runtime, projectId string, roleId string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	roleId = strings.TrimSpace(roleId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if roleId == "" {
		return Failure("roleId 不能为空").ToObject(vm)
	}

	var role map[string]any
	if _, err := a.cli.GetRole(ctx, projectId, roleId, &role); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(role).ToObject(vm)
}

// GetRoleUserIds 根据角色ID查询所有拥有该角色的用户列表
//
// projectId: 项目ID
//
// roleId: 角色ID
func (a Client) GetRoleUserIds(ctx context.Context, vm *goja.Runtime, projectId, roleId string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	roleId = strings.TrimSpace(roleId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if roleId == "" {
		return Failure("roleId 不能为空").ToObject(vm)
	}

	var role map[string]any
	if _, err := a.cli.GetRole(ctx, projectId, roleId, &role); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}

	if v, ok := role["userIds"]; ok {
		return Success(v).ToObject(vm)
	}

	return Success([]string{}).ToObject(vm)
}

// QueryRole 根据条件查询角色信息
//
// projectId: 项目ID
//
// query: 查询条件
func (a Client) QueryRole(ctx context.Context, vm *goja.Runtime, projectId string, query goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(query) {
		return Failure("无效的 query 参数").ToObject(vm)
	} else if !utils.IsObject(query) {
		return Failure("无效的 query 参数, 不是有效的对象").ToObject(vm)
	}

	var roles []map[string]any
	if err := a.cli.QueryRole(ctx, projectId, query, &roles); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}

	return Success(roles).ToObject(vm)
}
