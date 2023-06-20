package Test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestHttpsPost(t *testing.T) {
	urlx := "https://10.62.2.92/api/v2/loadbalancers"
	token := " lang=zh-CN; region=region0; tenant=7cef0ecb38904b338bc051b7dc4f80b7; scope=system; yunionauth=eyJleHAiOiIyMDIyLTA5LTIxVDA4OjEwOjMwWiIsImlzX3NzbyI6ZmFsc2UsInNlc3Npb24iOiJleUpoYkdjaU9pSlNVMEV4WHpVaUxDSjZhWEFpT2lKRVJVWWlMQ0psYm1NaU9pSkJNVEk0UjBOTkluMC5Ycmd3VzFaN3NrVDFPREZoTE5EeWJUUFVRbEhtQnVfeUxGUHJzaURDMEZuYU0tam1tN21RQ1RjNkRyQUh0bzRobDVYREVIdndMOFkxcXh6WVQtYTRyN2VnWHdSWDB0Q1RGelpVY21CWVRGU3ZvRHNER29RQXdiQWt6dS1MdEJfdHdPRWFoQ0IyaGFkYmktenZPU2oyNENnR3djcU4tTzdSeGtUNzA4Mjd5T1d3RVN6ZGh0VXp0cUxoaEhXd2VKSEU2aktPUUd1Y2pJNjFXTVgwVEpxWno4aVJEXzFzYkNlN3QzWXRvYmZCVmk2Slpadldsd0xBcV95YVhDY2ZNc3BFaDhxR1JlUnQwVGhMMG5iUjNrR3oxX3BQd3FqbWpiSHB0dXIwQ2ZuT3RQSGNkbDBmWUxSRm1jbUU3Mm1jaXppZ2lQMzRoQlp4QnNEQ08wOTdzbkpzZmcua202dkIyQllKT1FGMTFFcy5GWklKbGJqaG9iTGxoU1N6ZUFZWHlvUTJpaEx3NngzbkZ1ank0RkxzdmR5NnY1SWZhT2VJd01BajdDUk5PR0lBRlVnZVd1VFRLQ2JvbWhLcGRRMzVtLVhzZzVjVXJPcjMweFJ0cG1nSk9DTi1ORG4yRzEwTTdsRVJGMjVYNHMwMzBRNG42Y0xhRlg3dS1LNmFYSDhaZlVJV0hyNTlsUGhXdEVTSk1aTkhXZEJ3YUstVWJZN3lEcW8wX1FQc3JhNEQtRXJ1ZU91QTZ1anM2bThhMmlCdHlBekMxakZQUlJVNW9QY0o0aGVZdGI5dzFvLXg1cEU0TDZLVWU5cEQtLUZoYzVLWXF3LlJCSV9ZUHB1RW1qbDYyX0FQTUQ5bFEiLCJzeXN0ZW1fdG90cF9vbiI6dHJ1ZSwidG90cF9pbml0IjpmYWxzZSwidG90cF9vbiI6ZmFsc2UsInRvdHBfdmVyaWZpZWQiOmZhbHNlLCJ1c2VyIjoiYWRtaW4iLCJ1c2VyX2lkIjoiMGY2NDIxZGE1N2IzNGI4NDhjYTExOWUwNzk4MWFlMGYifQ; domain=Default"
	data := make(map[string]interface{})
	data["project"] = "d57c337ed18c4e0187f361c944ccbdbd"
	data["name"] = "loadkkx"
	data["provider"] = "Qcloud"
	data["cloudregion"] = "457472a3-7b91-4d64-8bcc-887e5a2502cf"
	data["zone"] = "6f6dc679-65ff-4639-8db0-8ed843cca5d3"
	data["ip"] = "ipv4"
	data["vpc"] = "e423071c-558d-49c8-8fb5-6bd6930b14a3"
	data["charge_type"] = "bandwidth"
	data["bandwidth"] = 1
	data["address_type"] = "internet"
	data["manager"] = "tencent"
	data["admin"] = "true"
	data["egress_mbps"] = "1"

	// Marshal
	bytesData, _ := json.Marshal(data)

	resp, _ := PostHttpsSkip(urlx, token, bytesData)
	fmt.Println(string(resp))
}

func PostHttpsSkip(url, token string, bytesData []byte) ([]byte, error) {
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

	// 获取 request请求
	request, err = http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		log.Println("PostHttpSkip Request Error:", err)
		return nil, nil
	}

	// 加入 token; Authorization、accessToken看你接口的请求头是什么了
	request.Header.Add("cookie", token)
	request.Header.Add("content-type", " application/json;charset=UTF-8")

	resp, err = client.Do(request)
	if err != nil {
		log.Println("PostHttpSkip Response Error:", err)
		return nil, nil
	}
	defer resp.Body.Close()

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
