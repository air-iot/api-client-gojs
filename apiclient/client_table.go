package apiclient

import (
	"context"
	"encoding/json"

	"github.com/dop251/goja"
)

func (a Client) GetTableSchema(ctx context.Context, vm *goja.Runtime, projectId string, tableId string) goja.Value {
	var result json.RawMessage
	_, err := a.cli.GetTableSchema(ctx, projectId, tableId, &result)
	if err != nil {
		return Failure(err.Error()).ToObject(vm)
	}
	return Success(string(result)).ToObject(vm)
}
