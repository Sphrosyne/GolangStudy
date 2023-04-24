package captcha

import (
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"testing"
)

func TestCaptcha(t *testing.T) {

	s := captcha.New()
	bytes, _ := json.Marshal(map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": s})
	fmt.Println(bytes)
	fmt.Println(map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": s})
	/*res := captcha.VerifyString(k, v)
	if res { // 验证通过
		fmt.Println("-=-----")
	} else { // 验证未通过
		fmt.Println("******")
	}*/
}
