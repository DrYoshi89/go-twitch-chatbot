package tools

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	request, _ := http.Get( url )
	content, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	text := string(content[:])
	return text
}
