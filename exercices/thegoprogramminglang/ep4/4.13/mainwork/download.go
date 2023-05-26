package mainwork

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func download(s string) error {
	slices := strings.Split(s, "/")
	resp, err := http.Get(s)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	os.WriteFile(fmt.Sprintf("%v/%v", "dir", slices[len(slices)-1]), data, 0644)
	return err

}
