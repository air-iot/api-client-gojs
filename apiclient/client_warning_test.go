package apiclient

import (
	"testing"

	"github.com/air-iot/gojs"
)

func TestCreateWarning(t *testing.T) {
	vm, err := gojs.GetVm()
	if err != nil {
		t.Fatalf("Error creating gojs vm: %s", err)
	}

	client := NewClient(cli)
	jsClient, err := client.CreateJsClient(vm)
	if err != nil {
		t.Fatalf("CreateJsClient err: %v", err)
	}

	if err := vm.Set("jsClient", jsClient); err != nil {
		t.Fatalf("Set err: %v", err)
	}

	_, err = vm.RunScript("testCreateWarning", `
	const result = jsClient.CreateWarning("default", {
		  "alert": true,
		  "desc": "测试告警",
		  "fields": [
			{
			  "id": "acceleration",
			  "value": 35
			}
		  ],
		  "interval": 10,
		  "level": "中",
		  "remark": "测试",
		  "ruleid": "67cfff4c7418cdd90180cb98",
		  "table": {
			"id": "zigbee"
		  },
		  "tableData": {
			"id": "integratedLoad_001"
		  },
		  "time": "2026-08-12T17:30:10.000Z",
		  "type": "1f849be4-2907-4459-9665-3fda916bf286"
	});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestBatchCreateWarning(t *testing.T) {
	vm, err := gojs.GetVm()
	if err != nil {
		t.Fatalf("Error creating gojs vm: %s", err)
	}

	client := NewClient(cli)
	jsClient, err := client.CreateJsClient(vm)
	if err != nil {
		t.Fatalf("CreateJsClient err: %v", err)
	}

	if err := vm.Set("jsClient", jsClient); err != nil {
		t.Fatalf("Set err: %v", err)
	}

	_, err = vm.RunScript("testCreateWarning", `
	const result = jsClient.CreateWarning("default", [{
		  "alert": true,
		  "desc": "测试告警1",
		  "fields": [
			{
			  "id": "acceleration",
			  "value": 37
			}
		  ],
		  "interval": 10,
		  "level": "低",
		  "remark": "测试1",
		  "ruleid": "67cfff4c7418cdd90180cb98",
		  "table": {
			"id": "zigbee"
		  },
		  "tableData": {
			"id": "integratedLoad_001"
		  },
		  "time": "2026-08-13T11:30:10.000Z",
		  "type": "1f849be4-2907-4459-9665-3fda916bf286"
	}, {
		  "alert": true,
		  "desc": "测试告警2",
		  "fields": [
			{
			  "id": "acceleration",
			  "value": 38
			}
		  ],
		  "interval": 10,
		  "level": "高",
		  "remark": "测试2",
		  "ruleid": "67cfff4c7418cdd90180cb98",
		  "table": {
			"id": "zigbee"
		  },
		  "tableData": {
			"id": "integratedLoad_001"
		  },
		  "time": "2026-08-13T11:30:10.000Z",
		  "type": "1f849be4-2907-4459-9665-3fda916bf286"
	}]);
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestGetWarning(t *testing.T) {
	vm, err := gojs.GetVm()
	if err != nil {
		t.Fatalf("Error creating gojs vm: %s", err)
	}

	client := NewClient(cli)
	jsClient, err := client.CreateJsClient(vm)
	if err != nil {
		t.Fatalf("CreateJsClient err: %v", err)
	}

	if err := vm.Set("jsClient", jsClient); err != nil {
		t.Fatalf("Set err: %v", err)
	}

	_, err = vm.RunScript("testGetWarning", `
	const result = jsClient.GetWarning("default", "6a7c3dc38f4a3ee1c5b7feae");
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestQueryWarning(t *testing.T) {
	vm, err := gojs.GetVm()
	if err != nil {
		t.Fatalf("Error creating gojs vm: %s", err)
	}

	client := NewClient(cli)
	jsClient, err := client.CreateJsClient(vm)
	if err != nil {
		t.Fatalf("CreateJsClient err: %v", err)
	}

	if err := vm.Set("jsClient", jsClient); err != nil {
		t.Fatalf("Set err: %v", err)
	}

	_, err = vm.RunScript("testQueryWarning", `
	const result = jsClient.QueryWarning("default", {
		"project": {
			"time": 1, "level": 1, "remark": 1, "type": 1,
		},
		"filter": {
			"tableId": "zigbee"
		}
});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}
