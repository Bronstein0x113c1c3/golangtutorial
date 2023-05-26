package work

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func get_image(ep int, link string) error {
	resp, _ := http.Get(link)
	data, err := ioutil.ReadAll(resp.Body)
	err = os.WriteFile(fmt.Sprintf("dir/%v.jpg", ep), data, 0644)
	return err
}
