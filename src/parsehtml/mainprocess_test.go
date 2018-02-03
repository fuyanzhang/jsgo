package parsehtml

import (
	"testing"
	"fmt"
)

func TestMainProcess(t *testing.T) {
	MainProcess()
}

func TestGetUserArtiDetail(t *testing.T)  {
	userinfo := new(UserInfo)
	userinfo.UserUrl = "https://www.jianshu.com/u/8c84a932666e"
	userinfo.ArticalNum = 84
	userinfo.UserId = "8c84a932666e"
	userinfo.NickName = "共央君"
	result :=getUserArtiDetail(*userinfo)
	fmt.Println(len(result))
	
}