package apiclient

import (
	"testing"

	"github.com/air-iot/gojs"
)

func TestGetRole(t *testing.T) {
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

	_, err = vm.RunScript("testGetRole", `
	const result = jsClient.GetRole("default", "68931ff17bd2c6ecc0cfaca4");
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestQueryRole(t *testing.T) {
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

	_, err = vm.RunScript("testQueryRole", `
	const result = jsClient.QueryRole("default", {"project": {"id": 1, "name": 1}});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}

func TestGetRoleUserIds(t *testing.T) {
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

	_, err = vm.RunScript("testGetRoleUserIds", `
	const result = jsClient.GetRoleUserIds("default", "68931ff17bd2c6ecc0cfaca4");
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}
