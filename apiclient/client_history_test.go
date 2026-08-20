package apiclient

import (
	"testing"

	"github.com/air-iot/gojs"
)

func TestQueryLatestSingle(t *testing.T) {
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

	_, err = vm.RunScript("testQueryTagLatest", `
	const result = jsClient.QueryTagLatest("default", "zigbee", "integratedLoad_001", "batteryVoltage");
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestQueryLatestMany(t *testing.T) {
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

	_, err = vm.RunScript("testQueryTagLatest", `
	const result = jsClient.QueryTagLatest("default", "zigbee", "integratedLoad_001", ["CommunicationEfficiency", "batteryVoltage"]);
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestQueryLatestAll(t *testing.T) {
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

	_, err = vm.RunScript("testQueryTagLatest", `
	const result = jsClient.QueryTagLatest("default", "zigbee", "integratedLoad_001", true);
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestQueryHistory(t *testing.T) {
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

	_, err = vm.RunScript("testQueryTagHistory", `
	const result = jsClient.QueryTagHistory("default", [
		{
			"fields": ["MEAN(batteryVoltage) as bv", "MEAN(CommunicationEfficiency) as commEfficiency"], 
			"id": "integratedLoad_001",
			"tableId": "zigbee",
			"where": [
				"time >= '2026-08-01T00:00:00.000Z'",
				"time < '2026-08-11T00:00:00.000Z'"
			],
			"group": ["time(10m)"],
			"fill": 0,
			"order": "time asc"
		}
	]);
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestSaveHistory(t *testing.T) {
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

	_, err = vm.RunScript("testSaveHistory", `
	const result = jsClient.SaveHistory("default", {"tableId": "zigbee", "tableDataId": "integratedLoad_001", "time": 1786501482535, "fields": {"acceleration": 0.152, "load": 23.751}});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}
