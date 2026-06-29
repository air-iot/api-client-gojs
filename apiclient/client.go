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

	return nil
}
