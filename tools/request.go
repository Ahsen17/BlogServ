/*
  Package teleport
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/6/2 22:31
  @Description: ...
*/

package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func BaseRequest(method string, uri string, param any) []byte {
	data, _ := json.Marshal(param)
	req := httptest.NewRequest(method, uri, bytes.NewReader(data))
	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func GetRequest(uri string, param any) []byte {
	return BaseRequest(http.MethodGet, uri, param)
}

func PostRequest(uri string, param any) []byte {
	return BaseRequest(http.MethodPost, uri, param)
}
