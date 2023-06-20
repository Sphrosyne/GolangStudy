package Test

import (
	"crypto/tls"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestHttpsGet(t *testing.T) {
	urlx := "https://10.62.2.92/api/v2/loadbalancers?scope=system&show_fail_reason=true&public_cloud=true&details=false&limit=20&id=94eab444-fa53-4912-8ecb-e4b263e1905e"
	token := "lang=zh-CN; region=region0; tenant=7cef0ecb38904b338bc051b7dc4f80b7; scope=system; yunionauth=eyJleHAiOiIyMDIyLTA5LTIyVDA4OjE2OjU2WiIsImlzX3NzbyI6ZmFsc2UsInNlc3Npb24iOiJleUpoYkdjaU9pSlNVMEV4WHpVaUxDSjZhWEFpT2lKRVJVWWlMQ0psYm1NaU9pSkJNVEk0UjBOTkluMC5NclVOTUdWSHRPc3I3NG5SWURlbzNMT00tZm56YzduR2ZuYUN3RVpBRHg4bDIwanFrQzdzNURhSFF5QkNJd1k4ZmtDOXk1RXI2Nnp1dW51aHloZWRfR2NYZkQ0dlltYlpZVGxtbHl4RE9iaWpkaDJHQ2RTY3RRNWlmZkdJamlTaTh2dmxKWTN3UEVBdEJKWElqSExLOXNUdkl6TFlGQmZwSmNFSmtpM2F6dmNpejk5T1IyXzlyODJaeTctZnM2b3JHMHRHNXoxMDNqSExkMVk0bWpkN01lLU1tYTBsRGRoY3BXbm5iYUpiQ1Q2bUphWUN5SHhmb21lWEQ0NlFvdU42MHV1QWtBbWhSUlFwcjB6TVFsdi1RUWp3ZXlkSml3TlI0OUtsMGZsbnNXTkZmYWFSelpDUlVKVlZSdlhIMTJGdGczVGFRZDJYNmd5bVdhY243OUMzLUEucXlmSXlVUURuVnVTVjk0Uy5KS1MzemtycmtFWTlzN2lCNVlhVWhYMmtkMS1BZGZ2eXhlTVFrNkJvU1JQbEdQTGszeVAxbEd5VDRuMnpfSS1EZjl0Z3BpOUd1MGt1SG5VQ05XUWMtTU9hZ1lDNDBPekNvcURramFvZk0wRU1hSEhTUEV1MTcwU0lnMGZETW5UZ1YtTElSb2lSaC11N2VKeVlmeTA0elk0ekdYeDEzTG1mV0k3UlFsU1VBeVpJZ3FOSmxwTF80cFRiVXNqMmZxN3hYUzBZWHUwV1VwWE1OdlNWd2hUVktELVpSRDdRZGlXMkFORWVxUGJOeVNLQ2ctdUZTVG5ZdGhwXy03UTlvQl94ZFkxYUduMXguQk5TTVh2N0ZHYXpKTG1ELV9NbGFiQSIsInN5c3RlbV90b3RwX29uIjp0cnVlLCJ0b3RwX2luaXQiOmZhbHNlLCJ0b3RwX29uIjpmYWxzZSwidG90cF92ZXJpZmllZCI6ZmFsc2UsInVzZXIiOiJhZG1pbiIsInVzZXJfaWQiOiIwZjY0MjFkYTU3YjM0Yjg0OGNhMTE5ZTA3OTgxYWUwZiJ9; domain=Default"

	resp, _ := GetHttpsSkip(urlx, token)
	fmt.Println(string(resp))

	value1 := gjson.Get(string(resp), "data.0.address")
	fmt.Println("------", value1.String())
}
func GetHttpsSkip(url, token string) ([]byte, error) {
	// 创建各类对象
	var (
		client  *http.Client
		request *http.Request
		resp    *http.Response
		body    []byte
		err     error
	)

	// 这里请注意，使用 InsecureSkipVerify: true 来跳过证书验证
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("GetHttpSkip Request Error:", err)
		return nil, nil
	}

	// 加入 token; Authorization、accessToken看你接口的请求头是什么了
	request.Header.Add("cookie", token)
	resp, err = client.Do(request)

	if err != nil {
		log.Println("GetHttpSkip Response Error:", err)
		return nil, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("resp.Body.close Error", err)
		}
	}(resp.Body)

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("ReadAll resp.Body Error:", err)
		return nil, nil
	}
	fmt.Println(resp.StatusCode)

	// 延迟关闭
	defer client.CloseIdleConnections()
	return body, nil
}
