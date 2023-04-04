package main

import (
	"fmt"
	"yunion.io/x/jsonutils"
)

var (
	allowResult = jsonutils.NewString("allow")
	denyResult  = jsonutils.NewString("deny")
)

//结果示例：{"compute":{"servers":{"*":"deny","get":"allow","list":"allow"}}}
func main() {
	services := map[string][]string{
		"compute": {"servers"},
	}

	policy := jsonutils.NewDict()
	for k, resList := range services {
		if len(resList) == 0 {
			resList = []string{"*"}
		}
		resPolicy := jsonutils.NewDict()
		for i := range resList {
			resPolicy.Add(getViewerActionPolicy(), resList[i])
		}
		policy.Add(resPolicy, k)
	}
	fmt.Println(policy)
}

func getViewerActionPolicy() jsonutils.JSONObject {
	p := jsonutils.NewDict()
	p.Add(allowResult, "get")
	p.Add(allowResult, "list")
	p.Add(denyResult, "*")
	return p
}
