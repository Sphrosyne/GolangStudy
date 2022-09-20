package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urlx := "https://10.62.2.92/api/v2/loadbalancers?scope=system&show_fail_reason=true&public_cloud=true&details=false&limit=20&id=1d6a4861-0146-4a89-81e4-3546f30655c9"
	token := " lang=zh-CN; region=region0; tenant=7cef0ecb38904b338bc051b7dc4f80b7; scope=system; yunionauth=eyJleHAiOiIyMDIyLTA5LTIxVDA4OjEwOjMwWiIsImlzX3NzbyI6ZmFsc2UsInNlc3Npb24iOiJleUpoYkdjaU9pSlNVMEV4WHpVaUxDSjZhWEFpT2lKRVJVWWlMQ0psYm1NaU9pSkJNVEk0UjBOTkluMC5Ycmd3VzFaN3NrVDFPREZoTE5EeWJUUFVRbEhtQnVfeUxGUHJzaURDMEZuYU0tam1tN21RQ1RjNkRyQUh0bzRobDVYREVIdndMOFkxcXh6WVQtYTRyN2VnWHdSWDB0Q1RGelpVY21CWVRGU3ZvRHNER29RQXdiQWt6dS1MdEJfdHdPRWFoQ0IyaGFkYmktenZPU2oyNENnR3djcU4tTzdSeGtUNzA4Mjd5T1d3RVN6ZGh0VXp0cUxoaEhXd2VKSEU2aktPUUd1Y2pJNjFXTVgwVEpxWno4aVJEXzFzYkNlN3QzWXRvYmZCVmk2Slpadldsd0xBcV95YVhDY2ZNc3BFaDhxR1JlUnQwVGhMMG5iUjNrR3oxX3BQd3FqbWpiSHB0dXIwQ2ZuT3RQSGNkbDBmWUxSRm1jbUU3Mm1jaXppZ2lQMzRoQlp4QnNEQ08wOTdzbkpzZmcua202dkIyQllKT1FGMTFFcy5GWklKbGJqaG9iTGxoU1N6ZUFZWHlvUTJpaEx3NngzbkZ1ank0RkxzdmR5NnY1SWZhT2VJd01BajdDUk5PR0lBRlVnZVd1VFRLQ2JvbWhLcGRRMzVtLVhzZzVjVXJPcjMweFJ0cG1nSk9DTi1ORG4yRzEwTTdsRVJGMjVYNHMwMzBRNG42Y0xhRlg3dS1LNmFYSDhaZlVJV0hyNTlsUGhXdEVTSk1aTkhXZEJ3YUstVWJZN3lEcW8wX1FQc3JhNEQtRXJ1ZU91QTZ1anM2bThhMmlCdHlBekMxakZQUlJVNW9QY0o0aGVZdGI5dzFvLXg1cEU0TDZLVWU5cEQtLUZoYzVLWXF3LlJCSV9ZUHB1RW1qbDYyX0FQTUQ5bFEiLCJzeXN0ZW1fdG90cF9vbiI6dHJ1ZSwidG90cF9pbml0IjpmYWxzZSwidG90cF9vbiI6ZmFsc2UsInRvdHBfdmVyaWZpZWQiOmZhbHNlLCJ1c2VyIjoiYWRtaW4iLCJ1c2VyX2lkIjoiMGY2NDIxZGE1N2IzNGI4NDhjYTExOWUwNzk4MWFlMGYifQ; domain=Default"

	resp, _ := GetHttpsSkip(urlx, token)
	fmt.Println(string(resp))
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
