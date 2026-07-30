package apiclient

import (
	"testing"

	"github.com/air-iot/gojs"
)

func TestClient_CreateSystemVariable(t *testing.T) {
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

	result, err := vm.RunScript("testCreateSystemVariable", `function handler(){
	const result1 = jsClient.CreateSystemVariable("default", {"id": "6a6aa8b198a2fc66f4c7ace7", "uid": "a", "name": "a", "type": "number", "value": 2});
	console.log("result1:", JSON.stringify(result1));
	return result1;
}
handler()
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	} else {
		t.Logf("RunScript result: %v", result.Export())
	}
}

func TestClient_GetSystemVariable(t *testing.T) {
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

	result, err := vm.RunScript("testGetSystemVariable", `function handler(){
	const result1 = jsClient.GetSystemVariable("default", "6a6aa8b198a2fc66f4c7ace7");
	console.log("result1:", JSON.stringify(result1));
	return result1;
}
handler()
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	} else {
		t.Logf("RunScript result: %v", result.Export())
	}
}

func TestClient_GetSystemVariableByUid(t *testing.T) {
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

	result, err := vm.RunScript("testGetSystemVariableByUid", `function handler(){
	const result1 = jsClient.GetSystemVariableByUid("default", "a");
	console.log("result1:", JSON.stringify(result1));
	return result1;
}
handler()
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	} else {
		t.Logf("RunScript result: %v", result.Export())
	}
}

func TestClient_GetSystemVariableByName(t *testing.T) {
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

	result, err := vm.RunScript("testGetSystemVariableByName", `function handler(){
	const result1 = jsClient.GetSystemVariableByName("default", "a");
	console.log("result1:", JSON.stringify(result1));
	return result1;
}
handler()
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	} else {
		t.Logf("RunScript result: %v", result.Export())
	}
}

func TestClient_UpdateSystemVariable(t *testing.T) {
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

	result, err := vm.RunScript("testUpdateSystemVariable", `function handler(){
	const result1 = jsClient.UpdateSystemVariable("default", "6a6aa8b198a2fc66f4c7ace7", {"value": 22});
	console.log("result1:", JSON.stringify(result1));
	return result1;
}
handler()
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	} else {
		t.Logf("RunScript result: %v", result.Export())
	}
}

func TestClient_ReplaceSystemVariable(t *testing.T) {
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

	result, err := vm.RunScript("testReplaceSystemVariable", `function handler(){
	const result1 = jsClient.ReplaceSystemVariable("default", "6a6aa8b198a2fc66f4c7ace7", {"uid": "aa", "name": "aa", "type": "number", "value": 222});
	console.log("result1:", JSON.stringify(result1));
	return result1;
}
handler()
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	} else {
		t.Logf("RunScript result: %v", result.Export())
	}
}

func TestClient_DeleteSystemVariable(t *testing.T) {
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

	result, err := vm.RunScript("testDeleteSystemVariable", `function handler(){
	const result1 = jsClient.DeleteSystemVariable("default", "6a6aa8b198a2fc66f4c7ace7");
	console.log("result1:", JSON.stringify(result1));
	return result1;
}
handler()
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	} else {
		t.Logf("RunScript result: %v", result.Export())
	}
}
