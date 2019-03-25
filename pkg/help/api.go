package help

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func DoPost(url string, params map[string]interface{}, headers map[string]string) ([]byte, int) {

	jsonByte, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonByte))
	if err != nil {
		fmt.Println(err.Error())
		return nil, -1
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 400 * time.Millisecond, //400ms
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, -1
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode
}
