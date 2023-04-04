package main

import (
	"fmt"
	"yunion.io/x/jsonutils"
)

var (
	allowResult = jsonutils.NewString("allow")
	denyResult  = jsonutils.NewString("deny")
)

//结果实例：{"compute":{"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}}}
func main() {
	services := map[string][]string{
		"compute": {"servers"},
	}

	policy := jsonutils.NewDict()
	if len(services) == 1 {
		for k := range services {
			if k == "*" {
				ns := make(map[string][]string)
				ns[k] = services[k]
				// expand adminPerformActions
				for s, resActions := range adminPerformActions {
					resList := make([]string, 0, len(resActions)+1)
					resList = append(resList, "*")
					for res := range resActions {
						resList = append(resList, res)
					}
					ns[s] = resList
				}
				services = ns
			}
		}
	}
	for k, resList := range services {
		resPolicy := jsonutils.NewDict()
		if len(resList) == 0 {
			resList = []string{"*"}
		}
		if len(resList) == 1 && resList[0] == "*" {
			if resActions, ok := adminPerformActions[k]; ok {
				for res := range resActions {
					resList = append(resList, res)
				}
			}
		}
		for i := range resList {
			resPolicy.Add(getEditActionPolicy(k, resList[i]), resList[i])
		}
		policy.Add(resPolicy, k)
	}
	fmt.Println(policy)
}

//获取
//结果实例：{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}
func getEditActionPolicy(service, resource string) jsonutils.JSONObject {
	p := jsonutils.NewDict()
	p.Add(denyResult, "create")
	p.Add(denyResult, "delete")
	perform := jsonutils.NewDict()
	perform.Add(denyResult, "purge")
	perform.Add(denyResult, "clone")
	if resActions, ok := adminPerformActions["compute"]; ok {
		if actions, ok := resActions["servers"]; ok {
			for _, action := range actions {
				//服务-》资源-》动作，对指定服务下的资源下的动作全部禁止
				perform.Add(denyResult, action)
			}
		}
	}
	perform.Add(allowResult, "*")
	p.Add(perform, "perform")
	p.Add(allowResult, "*")
	return p
}

var (
	adminPerformActions = map[string]map[string][]string{
		"compute": map[string][]string{
			"servers": []string{
				"snapshot-and-clone",
				"createdisk",
				"create-eip",
				"create-backup",
				"save-image",
				"delete-disk",
				"delete-eip",
				"delete-backup",
			},
			"buckets": []string{
				"upload",
				"delete",
			},
		},
		"k8s": map[string][]string{
			"kubeclusters": []string{
				"add-machines",
				"delete-machines",
			},
		},
	}
)
