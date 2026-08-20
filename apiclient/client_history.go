package apiclient

import (
	"context"
	nethttp "net/http"
	"net/url"
	"strings"

	"github.com/air-iot/api-client-gojs/v4/utils"
	"github.com/dop251/goja"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type restCallOption struct {
	before func()
}

// QueryTagLatest 查询数据点最新数据
//
// projectId 项目ID
//
// tableId 表标识
//
// deviceId 设备编号
//
// tags 数据点. 可以传递 string, []string, 或 true, 分别表示: 查询单个指定数据点的最新数据, 多个数据点的最新数据 或 该设备所有数据点的最新数据
func (a Client) QueryTagLatest(ctx context.Context, vm *goja.Runtime, projectId, tableId, deviceId string, tags goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)
	tableId = strings.TrimSpace(tableId)
	deviceId = strings.TrimSpace(deviceId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}
	if tableId == "" {
		return Failure("tableId 不能为空").ToObject(vm)
	}
	if deviceId == "" {
		return Failure("deviceId 不能为空").ToObject(vm)
	}

	if !utils.IsString(tags) && !utils.IsStringArray(tags) && !utils.IsBool(tags) {
		return Failure("tags 必须为 string 或 string[] 或 true").ToObject(vm)
	}

	var queryTags []map[string]any
	if utils.IsString(tags) {
		queryTags = append(queryTags, map[string]any{
			"tableId": tableId,
			"id":      deviceId,
			"tagId":   tags.Export(),
		})
	} else if utils.IsStringArray(tags) {
		if tagsIds, ok := utils.GetAsArrayString(tags); ok {
			if len(tagsIds) == 0 {
				return Failure("tags 必须为 string 或 string[] 或 true").ToObject(vm)
			}
			for _, tagId := range tagsIds {
				queryTags = append(queryTags, map[string]any{
					"tableId": tableId,
					"id":      deviceId,
					"tagId":   tagId,
				})
			}
		} else {
			return Failure("tags 必须为 string 或 string[] 或 true").ToObject(vm)
		}
	} else if utils.IsBool(tags) {
		allTags, ok := tags.Export().(bool)
		if !ok || !allTags {
			return Failure("tags 必须为 string 或 string[] 或 true").ToObject(vm)
		}

		queryTags = append(queryTags, map[string]any{
			"tableId": tableId,
			"id":      deviceId,
			"allTag":  true,
		})
	} else {
		return Failure("tags 必须为 string 或 string[] 或 true").ToObject(vm)
	}

	var result []map[string]any
	if err := a.cli.PostLatest(ctx, projectId, queryTags, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// QueryTagHistory 查询数据点的历史数据
//
// projectId 项目ID
//
// query 查询参数
func (a Client) QueryTagHistory(ctx context.Context, vm *goja.Runtime, projectId string, query goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}

	if !utils.IsObject(query) && !utils.IsObjectArray(query) {
		return Failure("query 不是有效的对象或对象数组").ToObject(vm)
	}

	var queries []any
	if utils.IsObject(query) {
		queries = append(queries, query.Export())
	} else if utils.IsObjectArray(query) {
		if objs, ok := utils.GetAsArrayObject(query); ok {
			for _, obj := range objs {
				queries = append(queries, obj)
			}
		} else {
			return Failure("query 不是有效的对象或对象数组").ToObject(vm)
		}
	} else {
		return Failure("query 不是有效的对象或对象数组").ToObject(vm)
	}

	var result map[string]any
	if err := a.cli.GetQuery(ctx, projectId, queries, &result); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}
	return Success(result).ToObject(vm)
}

// SaveHistory 写入历史数据
//
// projectId 项目ID
//
// data 写入的历史数据, 可以为对象或对象数组. {"tableId": "表标识", "tableDataId": "设备编号", "time": 1786501482535, "fields": {"A": 1, "B": true, "C": "c", "D": 1.24}}
func (a Client) SaveHistory(ctx context.Context, vm *goja.Runtime, projectId string, data goja.Value) goja.Value {
	projectId = strings.TrimSpace(projectId)

	if projectId == "" {
		return Failure("projectId 不能为空").ToObject(vm)
	}

	if !utils.IsObject(data) {
		return Failure("data 不是有效的对象或对象数组").ToObject(vm)
	}

	var points []any
	if utils.IsObject(data) {
		points = append(points, data.Export())
	} else if utils.IsObjectArray(data) {
		if objs, ok := utils.GetAsArrayObject(data); ok {
			for _, obj := range objs {
				points = append(points, obj)
			}
		} else {
			return Failure("data 不是有效的对象或对象数组").ToObject(vm)
		}
	} else {
		return Failure("data 不是有效的对象或对象数组").ToObject(vm)
	}

	cli, err := a.cli.CoreClient.GetRestClient()
	if err != nil {
		return ToFailureResult(err).ToObject(vm)
	}

	u := url.URL{Path: "/core/data/save/many"}
	headers := &nethttp.Header{
		"x-request-project": []string{projectId},
	}

	var result map[string]any
	if err := cli.Invoke(ctx, nethttp.MethodPost, u.RequestURI(), points, &result, http.Header(headers)); err != nil {
		return ToFailureResult(err).ToObject(vm)
	}

	return Success(result).ToObject(vm)
}
