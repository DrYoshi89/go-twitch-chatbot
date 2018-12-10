package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/xphip/gjson"
)

func Get(url string) string {
	request, _ := http.Get( url )
	content, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	text := string(content[:])
	return text
}

func GetURL(url string, headers map[string]string) string {

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest("GET", url, nil)

	for i, v := range headers {
		req.Header.Add(i, v)
		fmt.Printf("'%s': '%s'\n",i,v)
	}
	resp, _ := client.Do(req)
	// defer req.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content[:])
}

func PostURL(url string, form map[string]string) string {

	values, _ := json.Marshal(form)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(values))
	// defer req.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content[:])
}

func Var_dump(result interface{}) {
	fmt.Printf("<pre>%#v</pre>", result)
}

func Json(json, index string) (gjson.Result, error) {
	return gjson.Get(json, index), nil
}

func Use(vals ...interface{}) {
	_ = vals
}

func CheckErr(err error) {
	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(0)
	}
}

func Defer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()
}
