package apiclient

import (
	"testing"

	"github.com/air-iot/gojs"
)

func TestCallDataSource(t *testing.T) {
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

	_, err = vm.RunScript("testCallDataSource", `
	const result = jsClient.CallDataSource("default", "baidu", {});
	console.log("result:", JSON.stringify(result));
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}
}
