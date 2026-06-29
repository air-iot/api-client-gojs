# api-client-gojs

## 示例

```golang
package main

import (
    "github.com/dop251/goja"
    apiclient "github.com/air-iot/api-client-go/v4"
    "github.com/air-iot/api-client-go/v4/config"
    "github.com/air-iot/service/v4/etcd"
)

etcdCli, _, err := etcd.New(etcd.Config{
		Endpoints:        []string{"http://127.0.0.1:2379"},
		DialTimeout:      10,
		Username:         "root",
		Password:         "abc123456",
		AutoSyncInterval: 0,
	})
	if err != nil {
		panic(fmt.Errorf("failed to create etcd client: %v", err))
	}

	cli, _, err = apiclient.NewClient(etcdCli, config.Config{
		LiteMode:        false,
		Gateway:         "",
		GatewayGrpc:     "",
		EtcdConfig:      "/config/pro.json",
		Metadata:        nil,
		Services:        nil,
		Type:            "",
		ProjectId:       "",
		AK:              "",
		SK:              "",
		Timeout:         0,
		KeepAlive:       false,
		MaxIdleConns:    0,
		IdleConnTimeout: 0,
		Limit:           0,
		Debug:           true,
		Service: struct {
			Expire time.Duration `json:"expire"`
		}{},
		ExpirePrecision: 0,
	})

vm := goja.New()
client := NewClient(cli)

// 创建通用客户端. 该客户端的所有接口方法都需要传 '项目ID'
// 例如: projectJsClient.GetTableData("default", "student", "张三");
jsClient := client.CreateJsClient(vm)

// 创建项目客户端. 该客户端调用接口时不用传 '项目ID'
// 例如: projectJsClient.GetTableData("student", "张三");
projectJsClient := client.CreateProjectJsClient("default", vm)

vm.Set("apiClient", jsClient)
```
在 js 脚本使用方式如下:

```js
// 查询学生信息
const result = apiClient.GetTableData("default", "student", "张三");
if(result.success) {
  // 查询成功
  const student = result.data;
} else {
  throw new Error("查询失败:", result.message);
}
```

## 接口方法

```typescript

interface Result {
    success: boolean;
    message: string;
    count?: number;
    data?: any;
}

interface Client {
    /**
     * 根据记录ID查询记录数据
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param tableDataId 表记录ID
     * @constructor
     */
    GetTableData(projectId: string, tableId: string, tableDataId: string): Result;
    
    /**
     * 根据条件查询表记录数据
     * 
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param query 表记录ID
     * @constructor
     */
    QueryTableData(projectId: string, tableId: string, query: Record<string, any>): Result;
    
    /**
     * 新增表记录
     * 
     * @param projectId 项目ID
     * @param tableId   表标识
     * @param rowData 记录数据
     * @param closeRequire 是否关闭必填校验
     * @constructor
     */
    CreateTableData(projectId: string, tableId: string, rowData: any, closeRequire?: boolean): Result;
    
    /**
     * 批量新增表记录
     * 
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param rowsData 表记录数据
     * @param closeRequire 是否关闭必填校验
     * @constructor
     */
    CreateManyTableData(projectId: string, tableId: string, rowsData: Array<any>, closeRequire?: boolean): Result;
    
    /**
     * 更新表记录
     * 
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param rowData 表记录数据
     * @param closeRequire 是否关闭必填校验
     * @constructor
     */
    UpdateTableData(projectId: string, tableId: string, rowData: any, closeRequire?: boolean): Result;
    
    /**
     * 批量更新表记录数据
     * 
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param query 更新条件
     * @param updateFields 更新内容
     * @param closeRequire 是否关闭必填校验
     * @constructor
     */
    UpdateManyTableData(projectId: string, tableId: string, query: Record<string, any>, updateFields: Record<string, any>, closeRequire?: boolean): Result;
    
    /**
     * 替换表记录数据
     * 
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param tableDataId 表记录ID
     * @param rowData 表记录数据
     * @param closeRequire 是否关闭必填校验
     * @constructor
     */
    ReplaceTableData(projectId: string, tableId: string, tableDataId: string, rowData: any, closeRequire?: boolean): Result;
    
    /**
     * 根据ID删除表记录
     * 
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param tableDataId 表记录ID
     * @constructor
     */
    DeleteTableData(projectId: string, tableId: string, tableDataId: string):Result;
    
    /**
     * 根据条件删除表记录
     * 
     * @param projectId 项目ID
     * @param tableId 表标识
     * @param query 删除条件
     * @constructor
     */
    DeleteManyTableData(projectId: string, tableId: string, query: Record<string, any>): Result;
    
    /**
     * 根据用户ID查询用户信息
     * 
     * @param projectId 项目ID
     * @param userId 用户ID
     * @constructor
     */
    GetUser(projectId: string, userId: string): Result;
    
    /**
     * 根据条件批量查询用户信息
     * 
     * @param projectId 项目ID
     * @param query 查询参数
     * @constructor
     */
    QueryUser(projectId: string, query: Record<string, any>):Result;

    /**
     * 根据角色ID查询角色信息
     * 
     * @param projectId 项目ID
     * @param roleId 角色ID
     * @constructor
     */
    GetRole(projectId: string, roleId: string):Result;
    
    /**
     * 查询所有拥有目标角色的用户ID列表
     * 
     * @param projectId 项目ID
     * @param roleId 角色ID
     * @constructor
     */
    GetRoleUserIds(projectId: string, roleId: string);
    
    /**
     * 根据条件批量查询角色信息
     *
     * @param projectId 项目ID
     * @param query 查询参数
     * @constructor
     */
    QueryRole(projectId: string, query: Record<string, any>):Result;
}
```