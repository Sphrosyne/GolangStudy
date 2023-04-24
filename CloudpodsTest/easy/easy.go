package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"yunion.io/x/jsonutils"
)

type Menu struct {
	Allow    bool           `json:"allow"`
	Id       uint32         `json:"id"`
	Label    string         `json:"label"`
	Children []ChildrenMenu `json:"children"`
}
type ChildrenMenu struct {
	Allow bool   `json:"allow"`
	Id    uint32 `json:"id"`
	Label string `json:"label"`
}

func main() {
	menuStruct, err := jsonutils.ParseString("[{\"id\":1,\"label\":\"资源总览\",\"allow\":true},{\"id\":2,\"label\":\"主机管理\",\"allow\":true,\"children\":[{\"id\":3,\"label\":\"公有云主机\",\"allow\":true},{\"id\":4,\"label\":\"本地IDC\",\"allow\":true},{\"id\":5,\"label\":\"宿主机\",\"allow\":true},{\"id\":6,\"label\":\"主机快照\",\"allow\":false},{\"id\":7,\"label\":\"快照策略\",\"allow\":false},{\"id\":8,\"label\":\"秘钥管理\",\"allow\":false},{\"id\":9,\"label\":\"系统镜像\",\"allow\":true},{\"id\":10,\"label\":\"主机镜像\",\"allow\":false},{\"id\":11,\"label\":\"硬盘管理\",\"allow\":true},{\"id\":12,\"label\":\"硬盘快照\",\"allow\":false},{\"id\":13,\"label\":\"安全组\",\"allow\":true},{\"id\":14,\"label\":\"虚拟机\",\"allow\":true},{\"id\":15,\"label\":\"硬盘\",\"allow\":true},{\"id\":16,\"label\":\"镜像\",\"allow\":true}]},{\"id\":17,\"label\":\"网络管理\",\"allow\":true,\"children\":[{\"id\":18,\"label\":\"专用网络\",\"allow\":true},{\"id\":19,\"label\":\"IP子网管理\",\"allow\":true},{\"id\":20,\"label\":\"二层网络管理\",\"allow\":true},{\"id\":21,\"label\":\"弹性公网IP\",\"allow\":true}]},{\"id\":22,\"label\":\"数据库管理\",\"allow\":false,\"children\":[{\"id\":23,\"label\":\"RDS实例管理\",\"allow\":false},{\"id\":24,\"label\":\"redis实例管理\",\"allow\":false},{\"id\":25,\"label\":\"MongoDB实例管理\",\"allow\":false}]},{\"id\":26,\"label\":\"中间件\",\"allow\":false,\"children\":[{\"id\":27,\"label\":\"Kafka实例管理\",\"allow\":false}]},{\"id\":28,\"label\":\"云账号管理\",\"allow\":true,\"children\":[{\"id\":29,\"label\":\"账号管理\",\"allow\":true},{\"id\":30,\"label\":\"用户组管理\",\"allow\":false},{\"id\":31,\"label\":\"代理管理\",\"allow\":false}]},{\"id\":32,\"label\":\"认证与安全\",\"allow\":false,\"children\":[{\"id\":33,\"label\":\"团队管理\",\"allow\":false},{\"id\":34,\"label\":\"项目管理\",\"allow\":false},{\"id\":35,\"label\":\"用户管理\",\"allow\":false},{\"id\":36,\"label\":\"角色管理\",\"allow\":false},{\"id\":36,\"label\":\"权限管理\",\"allow\":false}]}]")
	if err == nil {
		fmt.Println(menuStruct)
	}
	myMenu := []Menu{}
	err = menuStruct.Unmarshal(&myMenu)
	fmt.Println("!!!", myMenu)
	for i, menu := range myMenu {
		fmt.Println(menu.Id)
		random, _ := uuid.NewRandom()
		myMenu[i].Id = random.ID()
		for k, child := range menu.Children {
			fmt.Println(child.Id)
			random, _ := uuid.NewRandom()
			myMenu[i].Children[k].Id = random.ID()
		}
	}
	marshal, err := json.Marshal(myMenu)
	parse, err := jsonutils.Parse(marshal)
	fmt.Println(parse)
	/*array, err := menuStruct.GetArray()
	for _, menu := range array {
		id, err := menu.GetString("id")

		if err != nil {
			return
		}
		fmt.Println(id)
		random, err := uuid.NewRandom()
		id = random.String()
		if menu.Contains("children") {
			children, err := menu.GetArray("children")
			if err != nil {
				return
			}
			for _, submenu := range children {
				subId, err := submenu.GetString("id")
				if err != nil {
					return
				}
				fmt.Println(subId)
				random, err := uuid.NewRandom()
				subId = random.String()
			}
		}
	}*/
	fmt.Println("****", myMenu)
	/*at, err := menuStruct.GetAt(0)
	fmt.Println("**", at)*/
}

func getMyString(myMenu []Menu) (string, error) {
	var buffer bytes.Buffer
	var err error
	var b []byte

	for _, item := range myMenu {
		b, err = json.Marshal(item)
		if err != nil {
			return "", err
		}

		buffer.WriteString(string(b) + ",")
	}

	s := strings.TrimSpace(buffer.String())

	return s, nil
}
