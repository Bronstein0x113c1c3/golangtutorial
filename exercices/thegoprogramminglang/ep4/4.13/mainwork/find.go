package mainwork

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Find(s string, api_key string) error {
	resp, err := http.Get(fmt.Sprintf("http://www.omdbapi.com/?apikey=%v&t=%v", api_key, s))
	data := &FilmInfo{}
	err = json.NewDecoder(resp.Body).Decode(data)
	download(data.Poster)
	return err
}
