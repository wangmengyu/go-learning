package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	fmt.Println("path:", path)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		//http.Error(writer,err.Error(),http.StatusInternalServerError)
		return err
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err.Error())
		return err

	}

	writer.Write(all)
	return nil

}
