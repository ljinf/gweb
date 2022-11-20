package jin

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
	ctx     context.Context

	//写保护机制
	writerMux *sync.Mutex
	params    map[string]string // url路由匹配的参数
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer:  w,
		request: r,
		params:  make(map[string]string),
	}
}

// 设置参数
func (c *Context) SetParams(params map[string]string) {
	c.params = params
}

func (c *Context) Param(key string) string {
	return c.params[key]
}

func (c *Context) Status(code int) {
	c.writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.writer.Header().Set(key, value)
}

func (c *Context) String(code int, str string) {
	c.writer.WriteHeader(code)
	c.writer.Header().Set("Content-Type", "text/plain")
	c.writer.Write([]byte(str))
}

func (c *Context) JSON(code int, obj ...interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.writer, err.Error(), 500)
	}
}
