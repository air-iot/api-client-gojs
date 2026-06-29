package utils

import (
	"testing"

	"github.com/dop251/goja"
)

func TestIsBool(t *testing.T) {
	vm := goja.New()
	if !IsBool(vm.ToValue(true)) {
		t.Fatalf("true -> expected: true, got: false")
	}
	if !IsBool(vm.ToValue(false)) {
		t.Fatalf("false -> expected: true, got: false")
	}

	if IsBool(vm.ToValue("false")) {
		t.Fatalf("string -> expected: false, got: true")
	}

	if IsBool(vm.ToValue(0)) {
		t.Fatalf("int -> expected: false, got: true")
	}
	if IsBool(vm.ToValue(1)) {
		t.Fatalf("int -> expected: false, got: true")
	}

	if IsBool(goja.Undefined()) {
		t.Fatalf("undefined -> expected: false, got: true")
	}
}

func TestIsObject(t *testing.T) {
	vm := goja.New()
	if IsObject(vm.ToValue(true)) {
		t.Fatalf("true -> expected: false, got: true")
	}
	if IsObject(vm.ToValue("string")) {
		t.Fatalf("string -> expected: false, got: true")
	}
	if IsObject(vm.ToValue("string")) {
		t.Fatalf("string -> expected: false, got: true")
	}
	if IsObject(vm.ToValue([]string{})) {
		t.Fatalf("[]string -> expected: false, got: true")
	}
	if IsObject(vm.ToValue([]map[string]string{})) {
		t.Fatalf("[]map[string]string -> expected: false, got: true")
	}
	if !IsObject(vm.ToValue(map[string]string{})) {
		t.Fatalf("map[string]string -> expected: true, got: false")
	}
}

func TestIsArray(t *testing.T) {
	vm := goja.New()
	if IsArray(vm.ToValue(true)) {
		t.Fatalf("true -> expected: false, got: true")
	}
	if IsArray(vm.ToValue("string")) {
		t.Fatalf("string -> expected: false, got: true")
	}
	if IsArray(vm.ToValue("string")) {
		t.Fatalf("string -> expected: false, got: true")
	}
	if IsArray(vm.ToValue(map[string]string{})) {
		t.Fatalf("map[string]string -> expected: false, got: true")
	}
	if !IsArray(vm.ToValue([]string{})) {
		t.Fatalf("[]string -> expected: true, got: false")
	}
	if !IsArray(vm.ToValue([]map[string]string{})) {
		t.Fatalf("[]map[string]string -> expected: false, got: true")
	}

}

func TestIsObjectArray(t *testing.T) {
	vm := goja.New()
	if IsObjectArray(vm.ToValue(true)) {
		t.Fatalf("true -> expected: false, got: true")
	}
	if IsObjectArray(vm.ToValue("string")) {
		t.Fatalf("string -> expected: false, got: true")
	}
	if IsObjectArray(vm.ToValue("string")) {
		t.Fatalf("string -> expected: false, got: true")
	}
	if IsObjectArray(vm.ToValue(map[string]string{})) {
		t.Fatalf("map[string]string -> expected: false, got: true")
	}
	if IsObjectArray(vm.ToValue([]string{})) {
		t.Fatalf("[]string -> expected: false, got: true")
	}
	if !IsObjectArray(vm.ToValue([]map[string]string{})) {
		t.Fatalf("[]map[string]string -> expected: true, got: false")
	}

}
