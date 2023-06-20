package Test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestLoadBalance(t *testing.T) {
	//变化性参数
	var (
		//申请下来的虚拟机id
		//vmId = make([]string, 5)
		vmId = []string{"22aefe8b-23d0-477b-8e21-17ee02db92b7", "a5491ba8-574d-4bb2-896b-f0c94d2b6dd3"}
		//全局登录校验token
		token = "lang=zh-CN; tenant=7cef0ecb38904b338bc051b7dc4f80b7; scope=system; yunionauth=eyJleHAiOiIyMDIyLTA5LTIyVDA4OjE2OjU2WiIsImlzX3NzbyI6ZmFsc2UsInNlc3Npb24iOiJleUpoYkdjaU9pSlNVMEV4WHpVaUxDSjZhWEFpT2lKRVJVWWlMQ0psYm1NaU9pSkJNVEk0UjBOTkluMC5NclVOTUdWSHRPc3I3NG5SWURlbzNMT00tZm56YzduR2ZuYUN3RVpBRHg4bDIwanFrQzdzNURhSFF5QkNJd1k4ZmtDOXk1RXI2Nnp1dW51aHloZWRfR2NYZkQ0dlltYlpZVGxtbHl4RE9iaWpkaDJHQ2RTY3RRNWlmZkdJamlTaTh2dmxKWTN3UEVBdEJKWElqSExLOXNUdkl6TFlGQmZwSmNFSmtpM2F6dmNpejk5T1IyXzlyODJaeTctZnM2b3JHMHRHNXoxMDNqSExkMVk0bWpkN01lLU1tYTBsRGRoY3BXbm5iYUpiQ1Q2bUphWUN5SHhmb21lWEQ0NlFvdU42MHV1QWtBbWhSUlFwcjB6TVFsdi1RUWp3ZXlkSml3TlI0OUtsMGZsbnNXTkZmYWFSelpDUlVKVlZSdlhIMTJGdGczVGFRZDJYNmd5bVdhY243OUMzLUEucXlmSXlVUURuVnVTVjk0Uy5KS1MzemtycmtFWTlzN2lCNVlhVWhYMmtkMS1BZGZ2eXhlTVFrNkJvU1JQbEdQTGszeVAxbEd5VDRuMnpfSS1EZjl0Z3BpOUd1MGt1SG5VQ05XUWMtTU9hZ1lDNDBPekNvcURramFvZk0wRU1hSEhTUEV1MTcwU0lnMGZETW5UZ1YtTElSb2lSaC11N2VKeVlmeTA0elk0ekdYeDEzTG1mV0k3UlFsU1VBeVpJZ3FOSmxwTF80cFRiVXNqMmZxN3hYUzBZWHUwV1VwWE1OdlNWd2hUVktELVpSRDdRZGlXMkFORWVxUGJOeVNLQ2ctdUZTVG5ZdGhwXy03UTlvQl94ZFkxYUduMXguQk5TTVh2N0ZHYXpKTG1ELV9NbGFiQSIsInN5c3RlbV90b3RwX29uIjp0cnVlLCJ0b3RwX2luaXQiOmZhbHNlLCJ0b3RwX29uIjpmYWxzZSwidG90cF92ZXJpZmllZCI6ZmFsc2UsInVzZXIiOiJhZG1pbiIsInVzZXJfaWQiOiIwZjY0MjFkYTU3YjM0Yjg0OGNhMTE5ZTA3OTgxYWUwZiJ9; domain=Default; region=region0"
		//所有请求访问地址
		requestAddress = "https://10.62.2.92/"

		//项目
		project  = "d57c337ed18c4e0187f361c944ccbdbd"
		provider = "Qcloud"
		//云相关
		cloudregion = "457472a3-7b91-4d64-8bcc-887e5a2502cf"
		zone        = "6f6dc679-65ff-4639-8db0-8ed843cca5d3"
		vpc         = "e423071c-558d-49c8-8fb5-6bd6930b14a3"
		manager     = "tencent"

		//访问相关
		path = "/image"
	)

	//1、新建负载均衡器，获得负载均衡器id（异步）
	lbId := newLoadBalance(token, requestAddress, project, provider, cloudregion, zone, vpc, manager)

	//2、根据负载均衡器id，查询公网ip
	lbIp := getLbIp(token, requestAddress, lbId)
	if len(lbIp) < 4 {
		log.Println("负载均衡器创建失败，结束")
		return
	}
	//3、指定负载均衡器id，创建后端服务器组，获得后端服务器组id
	groupId := newGroup(token, requestAddress, lbId)
	if len(groupId) < 4 {
		log.Println("新建后端服务器组失败，结束")
		return
	}
	time.Sleep(3 * time.Second)

	//4、指定后端服务器组backend_group（就是后端服务器组id），backend为虚拟机的id，guest_backend为虚拟机的id，添加服务器
	groupAddVM(token, requestAddress, groupId, vmId)
	time.Sleep(3 * time.Second)

	//5、指定负载均衡器loadbalancer（就是负载均衡器id），指定backend_group（就是后端服务器组id）,新建监听，返回id
	listenId := newListen(token, requestAddress, lbId, groupId)
	if len(listenId) < 4 {
		log.Println("新建监听失败，结束")
		return
	}
	time.Sleep(3 * time.Second)
	//6、指定listener（监听id），创建转发规则
	newTransmit(token, requestAddress, listenId, groupId, path)
}

func newLoadBalance(token, requestAddress, project, provider, cloudregion, zone, vpc, manager string) string {
	urlNewLoadBalance := requestAddress + "api/v2/loadbalancers"
	data1 := make(map[string]interface{})
	data1["project"] = project
	var builderLoadBalance strings.Builder
	builderLoadBalance.WriteString("lb")
	builderLoadBalance.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	data1["name"] = builderLoadBalance.String()
	data1["provider"] = provider
	data1["cloudregion"] = cloudregion
	data1["zone"] = zone
	data1["ip"] = "ipv4"
	data1["vpc"] = vpc
	data1["charge_type"] = "bandwidth"
	data1["bandwidth"] = 1
	data1["address_type"] = "internet"
	data1["manager"] = manager
	data1["admin"] = "true"
	data1["egress_mbps"] = 1
	// Marshal
	bytesData, _ := json.Marshal(data1)
	resp1, _ := postHttpsSkip(urlNewLoadBalance, token, bytesData)
	//获取新建的负载均衡器id
	log.Println("新建负载均衡器：", string(resp1))
	lbId := gjson.Get(string(resp1), "id").String()
	log.Println("新建的负载均衡器id：", lbId)
	return lbId
}

func getLbIp(token, requestAddress, lbId string) string {
	var builderGetLbIp strings.Builder
	builderGetLbIp.WriteString(requestAddress)
	builderGetLbIp.WriteString("api/v2/loadbalancers?scope=system&show_fail_reason=true&public_cloud=true&details=false&limit=20&id=")
	builderGetLbIp.WriteString(lbId)
	urlGetLoadBalance := builderGetLbIp.String()
	i := 0
	var lbIp string
	for {
		resp2, _ := getHttpsSkip(urlGetLoadBalance, token)
		log.Println("获取负载均衡器信息：", string(resp2))
		lbIp = gjson.Get(string(resp2), "data.0.address").String()
		i++
		if i >= 30 || len(lbIp) > 8 {
			log.Println("最后获取到负载均衡器ip:", lbIp)
			break
		}
		time.Sleep(2 * time.Second)
	}
	return lbIp
}

func newGroup(token, requestAddress, lbId string) string {
	urlNewGroup := requestAddress + "api/v2/loadbalancerbackendgroups"
	data1 := make(map[string]interface{})
	data1["name"] = "vgroup"
	data1["loadbalancer"] = lbId
	// Marshal
	bytesData, _ := json.Marshal(data1)
	resp1, _ := postHttpsSkip(urlNewGroup, token, bytesData)
	//获取负载均衡器添加后端服务器组反馈信息
	log.Println("负载均衡器添加后端服务器组：", string(resp1))
	groupId := gjson.Get(string(resp1), "id").String()
	log.Println("新建的后端服务器组id：", groupId)
	return groupId
}

func groupAddVM(token, requestAddress, groupId string, vmId []string) {
	for index, value := range vmId {
		log.Println("后端组添加服务器", index, value)
		urlGroupAddVM := requestAddress + "api/v2/loadbalancerbackends"
		data1 := make(map[string]interface{})
		data1["backend_type"] = "guest"
		data1["guest_backend"] = value
		data1["port"] = 80
		data1["weight"] = 1
		data1["backend_group"] = groupId
		data1["ssl"] = "off"
		data1["backend"] = value

		// Marshal
		bytesData, _ := json.Marshal(data1)
		resp1, _ := postHttpsSkip(urlGroupAddVM, token, bytesData)
		//获取后端服务器组添加服务器反馈信息
		log.Println("后端服务器组添加服务器：", string(resp1))
	}

}

func newListen(token, requestAddress, lbId, groupId string) string {
	url := requestAddress + "api/v2/loadbalancerlisteners"
	data1 := make(map[string]interface{})
	data1["loadbalancer"] = lbId
	data1["scheduler"] = "wrr"
	data1["generate_name"] = "listen03"
	data1["listener_port"] = 80
	data1["listener_type"] = "http"
	data1["gzip"] = false
	data1["xforwarded_for"] = true
	data1["backend_group"] = groupId
	data1["health_check"] = "on"
	data1["health_check_type"] = "http"
	data1["health_check_uri"] = "/"
	data1["health_check_http_code"] = "http_2xx,http_3xx"
	data1["health_check_timeout"] = 2
	data1["health_check_interval"] = 5
	data1["health_check_rise"] = 3
	data1["health_check_fall"] = 3
	data1["acl_status"] = "off"
	data1["sticky_session"] = "off"
	data1["redirect"] = "off"

	// Marshal
	bytesData, _ := json.Marshal(data1)
	resp1, _ := postHttpsSkip(url, token, bytesData)
	//获取创建监听反馈信息
	log.Println("负载均衡器创建监听：", string(resp1))
	listenId := gjson.Get(string(resp1), "id").String()
	log.Println("新建的监听id：", listenId)
	return listenId
}
func newTransmit(token, requestAddress, listenId, groupId, path string) {
	url := requestAddress + "api/v2/loadbalancerlistenerrules"
	i := 1
	for {
		data1 := make(map[string]interface{})
		data1["listener_type"] = "http"
		data1["name"] = "strategy" + strconv.FormatInt(time.Now().Unix(), 10)
		data1["domain"] = "www.luvs" + strconv.FormatInt(time.Now().Unix(), 10) + ".com"
		data1["path"] = path
		data1["backend_group"] = groupId
		data1["redirect"] = "off"
		data1["listener"] = listenId

		// Marshal
		bytesData, _ := json.Marshal(data1)
		resp1, _ := postHttpsSkip(url, token, bytesData)
		//获取创建转发规则反馈信息
		log.Println("负载均衡器创建转发规则：", string(resp1))
		ruleId := gjson.Get(string(resp1), "id").String()
		time.Sleep(3 * time.Second)

		//获取转发规则创建的最终结果create_failed、enabled、creating
		//7、查询转发规则创建情况，失败则继续创建几次
		status := getTransmitStatus(requestAddress, ruleId, listenId, token, i)
		if status == "enabled" || i >= 5 {
			log.Println("最终转发规则创建状态：", status)
			break
		}
		i++
	}
}
func getTransmitStatus(requestAddress, ruleId, listenId, token string, k int) string {
	var builderGetRuleStatus strings.Builder
	builderGetRuleStatus.WriteString(requestAddress)
	builderGetRuleStatus.WriteString("api/v2/loadbalancerlistenerrules?scope=system&show_fail_reason=true&details=false&limit=20")
	builderGetRuleStatus.WriteString("&listener=")
	builderGetRuleStatus.WriteString(listenId)
	builderGetRuleStatus.WriteString("&id=")
	builderGetRuleStatus.WriteString(ruleId)
	urlGetRuleStatus := builderGetRuleStatus.String()
	i := 0
	var status string
	for {
		resp2, _ := getHttpsSkip(urlGetRuleStatus, token)
		log.Println("获取转发规则状态信息：", string(resp2))
		status = gjson.Get(string(resp2), "data.0.status").String()
		i++
		if i >= 40 || status != "creating" {
			log.Println("第", k, "次转发规则创建状态:", status)
			break
		}
		time.Sleep(1 * time.Second)
	}
	return status
}

func postHttpsSkip(url, token string, bytesData []byte) ([]byte, error) {
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

func getHttpsSkip(url, token string) ([]byte, error) {
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
