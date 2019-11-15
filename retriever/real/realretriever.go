package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration //毫秒

}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	res, err := httputil.DumpResponse(resp, true)

	_ = resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(res)
}
