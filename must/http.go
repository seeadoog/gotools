package must

import (
	"io"
	"net/http"
)

//@throw error
func NewHttpRequest(method, url string, bd io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, bd)
	throw(err)
	return req
}
