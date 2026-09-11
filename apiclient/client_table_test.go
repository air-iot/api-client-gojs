package apiclient

import (
	"encoding/json"
	"reflect"
	"testing"

	apiclient "github.com/air-iot/api-client-go/v4"
	"github.com/air-iot/gojs"
	"github.com/dop251/goja"
)

var cli *apiclient.Client

//func TestMain(m *testing.M) {
//	etcdCli, _, err := etcd.New(etcd.Config{
//		Endpoints:        []string{"http://127.0.0.1:2379"},
//		DialTimeout:      10,
//		Username:         "root",
//		Password:         "dell123",
//		AutoSyncInterval: 0,
//	})
//	if err != nil {
//		panic(fmt.Errorf("failed to create etcd client: %v", err))
//	}
//
//	cli, _, err = apiclient.NewClient(etcdCli, config.Config{
//		LiteMode:        false,
//		Gateway:         "",
//		GatewayGrpc:     "",
//		EtcdConfig:      "/config/pro.json",
//		Metadata:        nil,
//		Services:        nil,
//		Type:            "",
//		ProjectId:       "",
//		AK:              "",
//		SK:              "",
//		Timeout:         0,
//		KeepAlive:       false,
//		MaxIdleConns:    0,
//		IdleConnTimeout: 0,
//		Limit:           0,
//		Debug:           true,
//		Service: struct {
//			Expire time.Duration `json:"expire"`
//		}{},
//		ExpirePrecision: 0,
//	})
//
//	if err != nil {
//		panic(fmt.Errorf("failed to create api client: %v", err))
//	}
//
//	os.Exit(m.Run())
//}

func Test_GetTableSchema(t *testing.T) {

}

func TestBigInt(t *testing.T) {

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

	_, err = vm.RunScript("testGetUser", `function handler(){
		const id = 43008827980319745n;
		return "{\"value\": " + id + "}";
}
`)

	if err != nil {
		t.Fatalf("RunScript err: %v", err)
	}

	handler, _ := goja.AssertFunction(vm.Get("handler"))
	value, _ := handler(goja.Undefined())
	t.Log(reflect.TypeOf(value))
	t.Log(reflect.TypeOf(value.Export()))
	t.Log(value.Export())
	data, err := json.Marshal(value.Export())
	if err != nil {
		t.Fatalf("Marshal err: %v", err)
	}
	t.Log(string(data))
}
