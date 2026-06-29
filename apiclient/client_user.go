package apiclient

import (
	"context"
	"strings"

	"github.com/air-iot/api-client-gojs/utils"
	"github.com/air-iot/gojs"
	"github.com/dop251/goja"
)

// GetUser 根据用户ID查询用户信息
//
// projectId: 项目ID
//
// userId: 用户ID
func (a Client) GetUser(ctx context.Context, vm *goja.Runtime, projectId string, userId string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	userId = strings.TrimSpace(userId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if userId == "" {
		return Failure("userId 不能为空").ToObject(vm)
	}

	var user map[string]any
	if _, err := a.cli.GetUser(ctx, projectId, userId, &user); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(user).ToObject(vm)
}

// QueryUser 根据条件查询用户信息
//
// projectId: 项目ID
//
// query: 查询参数
func (a Client) QueryUser(ctx context.Context, vm *goja.Runtime, projectId string, query goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(query) {
		return Failure("无效的 query 参数").ToObject(vm)
	} else if !utils.IsObject(query) {
		return Failure("无效的 query 参数, 不是有效的对象").ToObject(vm)
	}

	var users []map[string]any
	if err := a.cli.QueryUser(ctx, projectId, query, &users); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}

	return Success(users).ToObject(vm)
}
