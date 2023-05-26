package work

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBook(ep int) (*EpInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%v/info.0.json", ep))
	if err != nil {
		return nil, err
	}
	info := &EpInfo{}
	err = json.NewDecoder(resp.Body).Decode(info)
	err = get_image(ep, info.Img)
	return info, err
}
