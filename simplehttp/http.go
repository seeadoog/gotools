package simplehttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	method  string
	urls     []string
	headers map[string]string
	body    io.Reader
	timeout time.Duration
}

func New() *Request {
	return &Request{}
}

func GET() *Request {
	return New().GET()
}

func POST() *Request {
	return New().POST()
}

func (r *Request) GET() *Request {
	r.Method("GET")
	return r
}

func (r *Request) POST() *Request {
	r.Method("POST")
	return r
}

func (r *Request) DELETE() *Request {
	r.Method("DELETE")
	return r
}

func (r *Request) PATCH() *Request {
	r.Method("PATCH")
	return r
}

func (r *Request) Method(method string) *Request {
	r.method = method
	return r
}

func (r *Request) Urls(url ...string) *Request {
	r.urls = append(r.urls,url...)
	return r
}

func (r *Request) Timeout(t time.Duration) *Request {
	r.timeout = t
	return r
}

func (r *Request) SetHeader(key, val string) *Request {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[key] = val
	return r
}

func (r *Request) Body(body interface{}) *Request {
	var bd io.Reader
	switch t := body.(type) {
	case io.Reader:
		bd = t
	case []byte:
		bd = bytes.NewReader(t)
	case string:
		bd = strings.NewReader(t)
	default:
		bs, err := json.Marshal(t)
		if err != nil {
			throw(InvalidBodyType, err.Error())
		}
		bd = bytes.NewReader(bs)
	}
	r.body = bd
	return r
}

func (r *Request) do(url string) *Response {
	ctx := context.Background()
	if r.timeout > 0 {
		ctx, _ = context.WithTimeout(ctx, r.timeout)
	}
	req, err := http.NewRequestWithContext(ctx, r.method, url, r.body)
	if err != nil {
		throw(CreateRequestError, err.Error())
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		throw(DoRequestError, err.Error())
	}
	return &Response{rsp: rsp}
}


func (r *Request) Do() (resp *Response ){
	var err error
	for _, url := range r.urls {
		err = TryR(func() {
			resp = r.do(url)
		})
		if err == nil{
			return
		}
	}
	throwError(err)
	return
}

func (r *Request) DoWithRetry(retry int) *Response {
	if retry < 0 {
		throw(InvalidRetryNumber, "retry should be lager than 0")
	}
	var err error
	var resp *Response
	for i := 0; i <= retry; i++ {
		err = TryR(func() {
			resp = r.Do()
		})
		if err == nil {
			return resp
		}
	}
	throwError(err)
	return nil
}

type Response struct {
	rsp  *http.Response
	body []byte
}

func (r *Response) StatusCode() int {
	return r.rsp.StatusCode
}

func (r *Response) Body() []byte {
	if r.body != nil {
		return r.body
	}
	bd, err := ioutil.ReadAll(r.rsp.Body)
	if err != nil {
		throw(ResponseReadBodyError, err.Error())
	}
	r.body = bd
	return bd
}

func (r *Response) Text() string {
	return string(r.Body())
}

func (r *Response) Into(i interface{}) *Response {
	err := json.Unmarshal(r.Body(), i)
	if err != nil {
		throw(ResponseUnmarshalError, err.Error())
	}
	return r
}

func (r *Response) WriteTo(w io.Writer) *Response {
	err := r.rsp.Write(w)
	if err != nil {
		throw(ResponseWriteError, err.Error())
	}
	return r
}

func (r *Response) Header(key string) string {
	return r.rsp.Header.Get(key)
}
