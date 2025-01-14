/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package http

import (
	"github.com/nuclio/nuclio-sdk"
	"github.com/nuclio/nuclio/pkg/util/common"

	"github.com/valyala/fasthttp"
)

// allows accessing fasthttp.RequestCtx as a event.Sync
type Event struct {
	nuclio.AbstractSync
	ctx *fasthttp.RequestCtx
}

func (e *Event) GetContentType() string {
	return common.ByteArrayToString(e.ctx.Request.Header.ContentType())
}

func (e *Event) GetBody() []byte {
	return e.ctx.Request.Body()
}

func (e *Event) GetHeaderByteSlice(key string) []byte {

	// TODO: copy underlying by default? huge gotcha
	return e.ctx.Request.Header.Peek(key)
}

func (e *Event) GetHeaderString(key string) string {
	return string(e.GetHeaderByteSlice(key))
}

func (e *Event) GetMethod() string {
	return string(e.ctx.Request.Header.Method())
}

func (e *Event) GetPath() string {
	return string(e.ctx.Request.URI().Path())
}

func (e *Event) GetFieldByteSlice(key string) []byte {
	return e.ctx.QueryArgs().Peek(key)
}

func (e *Event) GetFieldString(key string) string {
	return string(e.GetFieldByteSlice(key))
}

func (e *Event) GetFieldInt(key string) (int, error) {
	return e.ctx.QueryArgs().GetUint(key)
}

func (e *Event) GetFields() map[string]interface{} {
	fields := make(map[string]interface{})
	e.ctx.QueryArgs().VisitAll(func(key, value []byte) {
		fields[common.ByteArrayToString(key)] = common.ByteArrayToString(value)
	})

	return fields
}
