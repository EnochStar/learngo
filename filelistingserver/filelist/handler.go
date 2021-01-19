package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
)

func Handler(writer http.ResponseWriter,
	request *http.Request) error {
	path := request.URL.Path[len("/learngo/"):]
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
