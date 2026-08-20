package apiclient

import (
	"context"
	"fmt"

	apiclient "github.com/air-iot/api-client-go/v4"
	"github.com/dop251/goja"
)

type Client struct {
	cli *apiclient.Client
}

func NewClient(cli *apiclient.Client) *Client {
	return &Client{
		cli: cli,
	}
}

// CreateJsClient 创建 js 客户端对象
func (a Client) CreateJsClient(vm *goja.Runtime) (*goja.Object, error) {
	jsClient := vm.NewObject()
	if err := a.AttachJsClient(vm, jsClient); err != nil {
		return nil, err
	}
	return jsClient, nil
}

func (a Client) CreateProjectJsClient(projectId string, vm *goja.Runtime) (*goja.Object, error) {
	if projectId == "" {
		return nil, fmt.Errorf("projectId 不能为空")
	}

	jsClient := vm.NewObject()
	if err := a.AttachProjectJsClient(projectId, vm, jsClient); err != nil {
		return nil, err
	}
	return jsClient, nil
}

func (a Client) AttachJsClient(vm *goja.Runtime, jsClient *goja.Object) error {
	{
		if err := jsClient.DefineDataProperty("GetTableData", vm.ToValue(func(projectId, tableId, tableDataId string) goja.Value {
			return a.GetTableData(context.Background(), vm, projectId, tableId, tableDataId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("QueryTableData", vm.ToValue(func(projectId, tableId string, query goja.Value) goja.Value {
			return a.QueryTableData(context.Background(), vm, projectId, tableId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("CreateTableData", vm.ToValue(func(projectId, tableId string, rowData, closeRequire goja.Value) goja.Value {
			return a.CreateTableData(context.Background(), vm, projectId, tableId, rowData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CreateTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("CreateManyTableData", vm.ToValue(func(projectId, tableId string, rowsData, closeRequire goja.Value) goja.Value {
			return a.CreateManyTableData(context.Background(), vm, projectId, tableId, rowsData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CreateManyTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("UpdateTableData", vm.ToValue(func(projectId, tableId, tableDataId string, rowData, closeRequire goja.Value) goja.Value {
			return a.UpdateTableData(context.Background(), vm, projectId, tableId, tableDataId, rowData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'UpdateTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("UpdateManyTableData", vm.ToValue(func(projectId, tableId string, query, updateFields, closeRequire goja.Value) goja.Value {
			return a.UpdateManyTableData(context.Background(), vm, projectId, tableId, query, updateFields, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'UpdateManyTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("ReplaceTableData", vm.ToValue(func(projectId, tableId, tableDataId string, rowData, closeRequire goja.Value) goja.Value {
			return a.ReplaceTableData(context.Background(), vm, projectId, tableId, tableDataId, rowData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'ReplaceTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("DeleteTableData", vm.ToValue(func(projectId, tableId, tableDataId string) goja.Value {
			return a.DeleteTableData(context.Background(), vm, projectId, tableId, tableDataId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'DeleteTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("DeleteManyTableData", vm.ToValue(func(projectId, tableId string, query goja.Value) goja.Value {
			return a.DeleteManyTableData(context.Background(), vm, projectId, tableId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'DeleteManyTableData' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("GetUser", vm.ToValue(func(projectId, userId string) goja.Value {
			return a.GetUser(context.Background(), vm, projectId, userId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetUser' 方法失败, %+v", err)
		}

		if err := jsClient.DefineDataProperty("QueryUser", vm.ToValue(func(projectId string, query goja.Value) goja.Value {
			return a.QueryUser(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryUser' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("GetRole", vm.ToValue(func(projectId, roleId string) goja.Value {
			return a.GetRole(context.Background(), vm, projectId, roleId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetRole' 方法失败, %+v", err)
		}

		if err := jsClient.DefineDataProperty("QueryRole", vm.ToValue(func(projectId string, query goja.Value) goja.Value {
			return a.QueryRole(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryRole' 方法失败, %+v", err)
		}

		if err := jsClient.DefineDataProperty("GetRoleUserIds", vm.ToValue(func(projectId, roleId string) goja.Value {
			return a.GetRoleUserIds(context.Background(), vm, projectId, roleId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetRoleUserIds' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("CallDataInterface", vm.ToValue(func(projectId, dsId string, params goja.Value) goja.Value {
			return a.CallDataInterface(context.Background(), vm, projectId, dsId, params)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CallDataInterface' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("GetSystemVariable", vm.ToValue(func(projectId, varId string) goja.Value {
			return a.GetSystemVariable(context.Background(), vm, projectId, varId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetSystemVariable' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("GetSystemVariableByName", vm.ToValue(func(projectId, varName string) goja.Value {
			return a.GetSystemVariableByName(context.Background(), vm, projectId, varName)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetSystemVariableByName' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("GetSystemVariableByUid", vm.ToValue(func(projectId, varUid string) goja.Value {
			return a.GetSystemVariableByUid(context.Background(), vm, projectId, varUid)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetSystemVariableByUid' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("QuerySystemVariable", vm.ToValue(func(projectId string, query goja.Value) goja.Value {
			return a.QuerySystemVariable(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QuerySystemVariable' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("CreateSystemVariable", vm.ToValue(func(projectId string, query goja.Value) goja.Value {
			return a.CreateSystemVariable(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CreateSystemVariable' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("UpdateSystemVariable", vm.ToValue(func(projectId, varId string, variable goja.Value) goja.Value {
			return a.UpdateSystemVariable(context.Background(), vm, projectId, varId, variable)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'UpdateSystemVariable' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("ReplaceSystemVariable", vm.ToValue(func(projectId, varId string, variable goja.Value) goja.Value {
			return a.ReplaceSystemVariable(context.Background(), vm, projectId, varId, variable)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'ReplaceSystemVariable' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("DeleteSystemVariable", vm.ToValue(func(projectId, varId string) goja.Value {
			return a.DeleteSystemVariable(context.Background(), vm, projectId, varId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'DeleteSystemVariable' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("QueryTagLatest", vm.ToValue(func(projectId, tableId, deviceId string, tags goja.Value) goja.Value {
			return a.QueryTagLatest(context.Background(), vm, projectId, tableId, deviceId, tags)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryTagLatest' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("QueryTagHistory", vm.ToValue(func(projectId string, query goja.Value) goja.Value {
			return a.QueryTagHistory(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryTagHistory' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("SaveHistory", vm.ToValue(func(projectId string, data goja.Value) goja.Value {
			return a.SaveHistory(context.Background(), vm, projectId, data)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'SaveHistory' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("CreateWarning", vm.ToValue(func(projectId string, warning goja.Value) goja.Value {
			return a.CreateWarning(context.Background(), vm, projectId, warning)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CreateWarning' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("GetWarning", vm.ToValue(func(projectId, warningId string, archive bool) goja.Value {
			return a.GetWarning(context.Background(), vm, projectId, warningId, archive)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetWarning' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("QueryWarning", vm.ToValue(func(projectId string, query goja.Value, archive bool) goja.Value {
			return a.QueryWarning(context.Background(), vm, projectId, query, archive)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryWarning' 方法失败, %+v", err)
		}
	}

	return nil
}

func (a Client) AttachProjectJsClient(projectId string, vm *goja.Runtime, jsClient *goja.Object) error {
	if projectId == "" {
		return fmt.Errorf("projectId 不能为空")
	}

	{
		if err := jsClient.DefineDataProperty("GetTableData", vm.ToValue(func(tableId, tableDataId string) goja.Value {
			return a.GetTableData(context.Background(), vm, projectId, tableId, tableDataId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("QueryTableData", vm.ToValue(func(tableId string, query goja.Value) goja.Value {
			return a.QueryTableData(context.Background(), vm, projectId, tableId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("CreateTableData", vm.ToValue(func(tableId string, rowData, closeRequire goja.Value) goja.Value {
			return a.CreateTableData(context.Background(), vm, projectId, tableId, rowData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CreateTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("CreateManyTableData", vm.ToValue(func(tableId string, rowsData, closeRequire goja.Value) goja.Value {
			return a.CreateManyTableData(context.Background(), vm, projectId, tableId, rowsData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CreateManyTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("UpdateTableData", vm.ToValue(func(tableId, tableDataId string, rowData, closeRequire goja.Value) goja.Value {
			return a.UpdateTableData(context.Background(), vm, projectId, tableId, tableDataId, rowData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'UpdateTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("UpdateManyTableData", vm.ToValue(func(tableId string, query, updateFields, closeRequire goja.Value) goja.Value {
			return a.UpdateManyTableData(context.Background(), vm, projectId, tableId, query, updateFields, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'UpdateManyTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("ReplaceTableData", vm.ToValue(func(tableId, tableDataId string, rowData, closeRequire goja.Value) goja.Value {
			return a.ReplaceTableData(context.Background(), vm, projectId, tableId, tableDataId, rowData, closeRequire)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'ReplaceTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("DeleteTableData", vm.ToValue(func(tableId, tableDataId string) goja.Value {
			return a.DeleteTableData(context.Background(), vm, projectId, tableId, tableDataId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'DeleteTableData' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("DeleteManyTableData", vm.ToValue(func(tableId string, query goja.Value) goja.Value {
			return a.DeleteManyTableData(context.Background(), vm, projectId, tableId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'DeleteManyTableData' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("GetUser", vm.ToValue(func(userId string) goja.Value {
			return a.GetUser(context.Background(), vm, projectId, userId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetUser' 方法失败, %+v", err)
		}

		if err := jsClient.DefineDataProperty("QueryUser", vm.ToValue(func(query goja.Value) goja.Value {
			return a.QueryUser(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryUser' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("GetRole", vm.ToValue(func(roleId string) goja.Value {
			return a.GetRole(context.Background(), vm, projectId, roleId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetRole' 方法失败, %+v", err)
		}

		if err := jsClient.DefineDataProperty("QueryRole", vm.ToValue(func(query goja.Value) goja.Value {
			return a.QueryRole(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryRole' 方法失败, %+v", err)
		}

		if err := jsClient.DefineDataProperty("GetRoleUserIds", vm.ToValue(func(roleId string) goja.Value {
			return a.GetRoleUserIds(context.Background(), vm, projectId, roleId)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetRoleUserIds' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("CallDataInterface", vm.ToValue(func(dsId string, params goja.Value) goja.Value {
			return a.CallDataInterface(context.Background(), vm, projectId, dsId, params)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CallDataInterface' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("QueryTagLatest", vm.ToValue(func(tableId, deviceId string, tags goja.Value) goja.Value {
			return a.QueryTagLatest(context.Background(), vm, projectId, tableId, deviceId, tags)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryTagLatest' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("QueryTagHistory", vm.ToValue(func(query goja.Value) goja.Value {
			return a.QueryTagHistory(context.Background(), vm, projectId, query)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryTagHistory' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("SaveHistory", vm.ToValue(func(data goja.Value) goja.Value {
			return a.SaveHistory(context.Background(), vm, projectId, data)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'SaveHistory' 方法失败, %+v", err)
		}
	}

	{
		if err := jsClient.DefineDataProperty("CreateWarning", vm.ToValue(func(warning goja.Value) goja.Value {
			return a.CreateWarning(context.Background(), vm, projectId, warning)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'CreateWarning' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("GetWarning", vm.ToValue(func(warningId string, archive bool) goja.Value {
			return a.GetWarning(context.Background(), vm, projectId, warningId, archive)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'GetWarning' 方法失败, %+v", err)
		}
		if err := jsClient.DefineDataProperty("QueryWarning", vm.ToValue(func(query goja.Value, archive bool) goja.Value {
			return a.QueryWarning(context.Background(), vm, projectId, query, archive)
		}), goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE); err != nil {
			return fmt.Errorf("注册 'QueryWarning' 方法失败, %+v", err)
		}
	}

	return nil
}
