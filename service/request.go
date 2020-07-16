package service

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
)

//请求服务
func RequestPost(method string, rurl string, postArgs map[string]string, headerArgs map[string]string) ([]byte, error) {
	var clusterinfo = url.Values{}
	for key, value := range postArgs {
		clusterinfo.Add(key, value)
	}
	data := clusterinfo.Encode()
	reader := strings.NewReader(data)
	request, err := http.NewRequest(method, rurl, reader)
	if err != nil {
		beego.Debug(err.Error())
		return nil, err
	}
	for key, value := range headerArgs {
		request.Header.Set(key, value)
	}
	client := http.Client{}
	resp, err := client.Do(request)
	beego.Debug(resp, err)
	if err != nil {
		beego.Debug(err.Error())
		return nil, err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	return respBytes, nil
}
