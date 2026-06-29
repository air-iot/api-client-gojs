package apiclient

import (
	"fmt"
	"os"
	"testing"
	"time"

	apiclient "github.com/air-iot/api-client-go/v4"
	"github.com/air-iot/api-client-go/v4/config"
	"github.com/air-iot/service/v4/etcd"
)

var cli *apiclient.Client

func TestMain(m *testing.M) {
	etcdCli, _, err := etcd.New(etcd.Config{
		Endpoints:        []string{"http://127.0.0.1:2379"},
		DialTimeout:      10,
		Username:         "root",
		Password:         "dell123",
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

	if err != nil {
		panic(fmt.Errorf("failed to create api client: %v", err))
	}

	os.Exit(m.Run())
}

func Test_GetTableSchema(t *testing.T) {

}
