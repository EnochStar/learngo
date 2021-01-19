package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/learngo/"

type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

func Handler(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start" +
			"width" + prefix)
	}
	path := request.URL.Path[len(prefix):]
	open, err := os.Open(path)
	if err != nil {
		return err
	}
	defer open.Close()
	all, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
