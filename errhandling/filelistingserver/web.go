package main

import (
	"fmt"
	"github.com/wangmengyu/go-learning/errhandling/filelistingserver/filelisting"
	"net/http"
	_ "net/http/pprof"
	"os"
)

/**
定义一种类，类型是错误类型

*/
type appHandler func(writer http.ResponseWriter, request *http.Request) error

/**
定义一种方法，返回类型是 func(http.ResponseWriter,  *http.Request) 但是会处理其中的error
*/
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		code := http.StatusOK
		if err != nil {
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError

			}
			fmt.Println(err)
			http.Error(writer, http.StatusText(code), code)
		}

	}

}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	_ = http.ListenAndServe(":8888", nil)

}
