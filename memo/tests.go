package memo

import (
	"io/ioutil"
	"net/http"
)

func httpGetBody(url string) (interface{}, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	return ioutil.ReadAll(r.Body)
}
