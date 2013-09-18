package gotcdata

import (
	"io/ioutil"
	"net/http"
)

const (
	UrlCoderList string = "http://www.topcoder.com/tc?module=BasicData&c=dd_coder_list"
)


func NewCoderList() (cl CoderList, err error) {
	resp, err := http.Get(UrlCoderList)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = unmarshal(body, &cl)

	return
}
