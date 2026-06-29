package apiclient

import (
	"testing"

	"github.com/air-iot/gojs"
)

func TestGetTableData(t *testing.T) {
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

	_, err = vm.RunScript("testGetTableData1", `
	const result1 = jsClient.GetTableData("default", "detect_history", "6971f7a08e8e1354fa21299a");
	console.log("result1:", JSON.stringify(result1));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}

	_, err = vm.RunScript("testGetTableData2", `
	const result2 = jsClient.GetTableData("default", "detect_history");
	if(!result2.success) {
		throw new Error(result2.message);
	}
	console.log("result2:", JSON.stringify(result2));
`)

	if err == nil {
		t.Fatalf("RunScript: expect error")
	}
}

func TestQueryTableData(t *testing.T) {
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

	_, err = vm.RunScript("testQueryTableData", `
	const result = jsClient.QueryTableData("default", "detect_history", {"project": {"id": 1, "flight_id0": 1}, "filter": {"event_id0": "1747453267264"}, "withCount": true});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestCreateTableData(t *testing.T) {
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

	_, err = vm.RunScript("testCreateTableData", `
	const result = jsClient.CreateTableData("default", "ccs_devices", {"id": "ccs_001", "name": "船级社_001"});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestCreateManyTableData(t *testing.T) {
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

	_, err = vm.RunScript("testCreateManyTableData", `
	const result = jsClient.CreateManyTableData("default", "ccs_devices", [{"id": "ccs_002", "name": "船级社_002"}]);
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestUpdateTableData(t *testing.T) {
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

	_, err = vm.RunScript("testUpdateTableData", `
	const result = jsClient.UpdateTableData("default", "ccs_devices", "ccs_001", { "name": "船级社_101"});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestUpdateManyTableData(t *testing.T) {
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

	_, err = vm.RunScript("testUpdateManyTableData", `
	const result = jsClient.UpdateManyTableData("default", "ccs_devices", {"id": {"$in": ["ccs_001", "ccs_002"]}}, { "age": 11});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestReplaceTableData(t *testing.T) {
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

	_, err = vm.RunScript("testReplaceTableData", `
	const result = jsClient.ReplaceTableData("default", "ccs_devices", "ccs_001", { "name": "船级社111", "age": 111});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestDeleteTableData(t *testing.T) {
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

	_, err = vm.RunScript("testDeleteTableData", `
	const result = jsClient.DeleteTableData("default", "ccs_devices", "ccs_001");
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestDeleteManyTableData(t *testing.T) {
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

	_, err = vm.RunScript("testDeleteManyTableData", `
	const result = jsClient.DeleteManyTableData("default", "ccs_devices", {"age": {"$gt": "10"}});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestProjectGetTableData(t *testing.T) {
	vm, err := gojs.GetVm()
	if err != nil {
		t.Fatalf("Error creating gojs vm: %s", err)
	}

	client := NewClient(cli)
	jsClient, err := client.CreateProjectJsClient("default", vm)
	if err != nil {
		t.Fatalf("CreateJsClient err: %v", err)
	}

	if err := vm.Set("jsClient", jsClient); err != nil {
		t.Fatalf("Set err: %v", err)
	}

	_, err = vm.RunScript("testGetTableData1", `
	const result1 = jsClient.GetTableData("detect_history", "6971f7a08e8e1354fa21299a");
	console.log("result1:", JSON.stringify(result1));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}

	_, err = vm.RunScript("testGetTableData2", `
	const result2 = jsClient.GetTableData("detect_history");
	if(!result2.success) {
		throw new Error(result2.message);
	}
	console.log("result2:", JSON.stringify(result2));
`)

	if err == nil {
		t.Fatalf("RunScript: expected error")
	}
}

func TestProjectQueryTableData(t *testing.T) {
	vm, err := gojs.GetVm()
	if err != nil {
		t.Fatalf("Error creating gojs vm: %s", err)
	}

	client := NewClient(cli)
	jsClient, err := client.CreateProjectJsClient("default", vm)
	if err != nil {
		t.Fatalf("CreateJsClient err: %v", err)
	}

	if err := vm.Set("jsClient", jsClient); err != nil {
		t.Fatalf("Set err: %v", err)
	}

	_, err = vm.RunScript("testQueryTableData", `
	const result = jsClient.QueryTableData("detect_history", {"project": {"id": 1, "flight_id0": 1}, "filter": {"event_id0": "1747453267264"}, "withCount": true});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}
