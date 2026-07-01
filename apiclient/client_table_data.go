package apiclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/air-iot/api-client-gojs/v4/utils"
	"github.com/air-iot/errors"
	"github.com/air-iot/gojs"
	"github.com/dop251/goja"
)

// GetTableData 根据记录ID查询表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// tableDataId: 表记录ID
func (a Client) GetTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId, tableDataId string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)
	tableDataId = strings.TrimSpace(tableDataId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}
	if tableDataId == "" {
		return Failure("tableDataId 不能为空").ToObject(vm)
	}

	var result map[string]interface{}
	if _, err := a.cli.GetTableData(ctx, projectId, tableId, tableDataId, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// QueryTableData 根据自定义查询参数查询表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// query: 查询参数
func (a Client) QueryTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId string, query goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)
	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(query) {
		return Failure("无效的查询参数").ToObject(vm)
	} else if !utils.IsObject(query) {
		return Failure("无效的查询参数, 不是有效的对象").ToObject(vm)
	}

	var result []map[string]interface{}
	count, err := a.cli.QueryTableData(ctx, projectId, tableId, query.Export(), &result)
	if err != nil {
		return Failure(err.Error()).ToObject(vm)
	}
	return Success(vm.ToValue(result)).WithCount(count).ToObject(vm)
}

// CreateTableData 新增表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// rowData: 表记录数据
//
// closeRequire: 是否关闭字段必填校验
func (a Client) CreateTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId string, rowData goja.Value, closeRequire goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}

	var notRequired bool
	if utils.IsBool(closeRequire) {
		notRequired = closeRequire.ToBoolean()
	} else if gojs.IsValid(closeRequire) {
		return Failure("closeRequire 必须为 boolean 值").ToObject(vm)
	}

	if !gojs.IsValid(rowData) {
		return Failure("rowData 不能为空").ToObject(vm)
	} else if !utils.IsObject(rowData) {
		return Failure("rowData 不是有效的对象").ToObject(vm)
	}

	var result InsertResult
	if err := a.cli.CreateTableData(ctx, projectId, tableId, notRequired, rowData.Export(), &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// CreateManyTableData 批量新增表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// rowsData: 表记录数据
//
// closeRequire: 是否关闭字段必填校验
func (a Client) CreateManyTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId string, rowsData goja.Value, closeRequire goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}

	var notRequired bool
	if utils.IsBool(closeRequire) {
		notRequired = closeRequire.ToBoolean()
	} else if gojs.IsValid(closeRequire) {
		return Failure("closeRequire 必须为 boolean 值").ToObject(vm)
	}

	if !gojs.IsValid(rowsData) {
		return Failure("rowsData 不能为空").ToObject(vm)
	}

	rows, ok := utils.GetAsArrayObject(rowsData)
	if !ok {
		return Failure("rowsData 不是有效的对象数组").ToObject(vm)
	}

	var result InsertsResult
	if err := a.cli.CreateManyTableData(ctx, projectId, tableId, notRequired, rows, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// UpdateTableData 更新单条表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// tableDataId: 表记录ID
//
// rowData: 表记录数据
//
// closeRequire: 是否关闭字段必填校验
func (a Client) UpdateTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId, tableDataId string, rowData goja.Value, closeRequire goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)
	tableDataId = strings.TrimSpace(tableDataId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}
	if tableDataId == "" {
		return Failure("tableDataId 不能为空").ToObject(vm)
	}

	var notRequired bool
	if utils.IsBool(closeRequire) {
		notRequired = closeRequire.ToBoolean()
	} else if gojs.IsValid(closeRequire) {
		return Failure("closeRequire 必须为 boolean 值").ToObject(vm)
	}

	if !gojs.IsValid(rowData) {
		return Failure("rowData 不能为空").ToObject(vm)
	} else if !utils.IsObject(rowData) {
		return Failure("rowData 不是有效的对象").ToObject(vm)
	}

	var result StatusResult
	if err := a.cli.UpdateTableData(ctx, projectId, tableId, tableDataId, notRequired, rowData.Export(), &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// UpdateManyTableData 批量更新表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// query: 更新条件
//
// updateFields: 更新内容
//
// closeRequire: 是否关闭字段必填校验
func (a Client) UpdateManyTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId string, query, updateFields goja.Value, closeRequire goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}

	var notRequired bool
	if utils.IsBool(closeRequire) {
		notRequired = closeRequire.ToBoolean()
	} else if gojs.IsValid(closeRequire) {
		return Failure("closeRequire 必须为 boolean 值").ToObject(vm)
	}

	if !gojs.IsValid(updateFields) {
		return Failure("updateFields 不能为空").ToObject(vm)
	} else if !utils.IsObject(updateFields) {
		return Failure("updateFields 不是有效的对象").ToObject(vm)
	}

	if !gojs.IsValid(query) {
		return Failure("query 不能为空").ToObject(vm)
	} else if !utils.IsObject(query) {
		return Failure("query 不是有效的对象").ToObject(vm)
	}

	var result StatusResult
	if err := a.cli.UpdateManyTableData(ctx, projectId, tableId, notRequired, query.Export(), updateFields.Export(), &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// ReplaceTableData 替换表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// tableDataId: 表记录ID
//
// rowData: 表记录数据
//
// closeRequire: 是否关闭字段必填校验
func (a Client) ReplaceTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId, tableDataId string, rowData goja.Value, closeRequire goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)
	tableDataId = strings.TrimSpace(tableDataId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}
	if tableDataId == "" {
		return Failure("tableDataId 不能为空").ToObject(vm)
	}

	var notRequired bool
	if utils.IsBool(closeRequire) {
		notRequired = closeRequire.ToBoolean()
	} else if gojs.IsValid(closeRequire) {
		return Failure("closeRequire 必须为 boolean 值").ToObject(vm)
	}

	if !gojs.IsValid(rowData) {
		return Failure("rowData 不能为空").ToObject(vm)
	} else if !utils.IsObject(rowData) {
		return Failure("rowData 不是有效的对象").ToObject(vm)
	}

	var result StatusResult
	if err := a.cli.ReplaceTableData(ctx, projectId, tableId, tableDataId, notRequired, rowData.Export(), &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// DeleteTableData 表除单条表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// tableDataId: 表记录ID
func (a Client) DeleteTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId, tableDataId string) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)
	tableDataId = strings.TrimSpace(tableDataId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}
	if tableDataId == "" {
		return Failure("tableDataId 不能为空").ToObject(vm)
	}

	var result StatusResult
	if err := a.cli.DeleteTableData(ctx, projectId, tableId, tableDataId, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// DeleteManyTableData 批量删除表记录
//
// projectId: 项目ID
//
// tableId: 表标识
//
// query: 删除条件
func (a Client) DeleteManyTableData(ctx context.Context, vm *goja.Runtime, projectId, tableId string, query goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)
	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}

	if !gojs.IsValid(query) {
		return Failure("无效的查询参数").ToObject(vm)
	} else if !utils.IsObject(query) {
		return Failure("无效的查询参数, 不是有效的对象").ToObject(vm)
	}

	var result StatusResult
	err := a.cli.DeleteManyTableData(ctx, projectId, tableId, query.Export(), &result)
	if err != nil {
		return Failure(err.Error()).ToObject(vm)
	}
	return Success(vm.ToValue(result)).ToObject(vm)
}

func ToFailureResult(err error) *Result {
	var responseErr *errors.ResponseError
	if errors.As(err, &responseErr) {
		if responseErr.ERR != nil {
			return Failure(fmt.Sprintf("%s, %s", responseErr.Message, responseErr.ERR))
		}
		return Failure(responseErr.Message)
	}
	return Failure(err.Error())
}
