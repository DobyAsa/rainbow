package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request  *http.Request
	response http.ResponseWriter
	ctx      context.Context
	handler  ControllerHandler

	hasTimeout bool

	writeMux *sync.Mutex
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		request:  r,
		response: w,
		ctx:      r.Context(),

		writeMux: &sync.Mutex{},
	}
}

func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writeMux
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.response
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.request.Context().Done()
}

func (ctx *Context) Err() error {
	return ctx.request.Context().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.request.Context().Value(key)
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.request.Context().Deadline()
}
func (ctx *Context) QueryInt(key string, def int) (val int, ok bool) {
	params := ctx.QueryAll()
	if v, ok := params[key]; ok {
		len := len(v)
		if len > 0 {
			intVal, err := strconv.Atoi(v[len-1])
			if err != nil {
				return def, false
			}
			return intVal, true
		}
	}
	return def, false
}

func (ctx *Context) QueryString(key string, def string) (val string, ok bool) {
	params := ctx.QueryAll()
	if v, ok := params[key]; ok {
		len := len(v)
		if len > 0 {
			return v[len-1], true
		}
	}
	return def, false
}

func (ctx *Context) QueryArray(key string, def []string) (val []string, ok bool) {
	params := ctx.QueryAll()
	if v, ok := params[key]; ok {
		return v, true
	}
	return def, false
}
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) (val int, ok bool) {
	params := ctx.FormAll()
	if v, ok := params[key]; ok {
		len := len(v)
		if len > 0 {
			intVal, err := strconv.Atoi(v[len-1])
			if err != nil {
				return def, false
			}
			return intVal, true
		}
	}
	return def, false
}

func (ctx *Context) FormString(key string, def string) (val string, ok bool) {
	params := ctx.FormAll()
	if v, ok := params[key]; ok {
		len := len(v)
		if len > 0 {
			return v[len-1], true
		}
	}
	return def, false
}

func (ctx *Context) FormArray(key string, def []string) (val []string, ok bool) {
	params := ctx.FormAll()
	if v, ok := params[key]; ok {
		return v, true
	}
	return def, false
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.Form
	}
	return map[string][]string{}
}

func (ctx *Context) BindJSON(obj interface{}) error {
	if ctx.request != nil {
		body, err := io.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}

		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("ctx.request empty")
}

func (ctx *Context) JSON(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}

	ctx.response.Header().Set("Content-Type", "application/json")
	ctx.response.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	ctx.response.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}) error {
	return errors.New("ctx.HTML() method body is empty")
}

func (ctx *Context) Text(status int, text string) error {
	return errors.New("ctx.Text() method body is empty")
}
