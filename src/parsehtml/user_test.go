package parsehtml

import (
	"testing"
	"fmt"
	"strconv"
)

func TestGetUsersId(t *testing.T) {
	uids := GetUsersId("https://www.jianshu.com/recommendations/users?page=1")
	fmt.Println(uids)
}

func TestGetUserInfo(t *testing.T) {
	userInfo := GetUserInfo("https://www.jianshu.com/u/dbfdce352c0d")
	fmt.Println(userInfo.ArticalNum)
	fmt.Println(userInfo.UserFollower)
	fmt.Println(userInfo.UserFollowing)
	fmt.Println(userInfo.UserUrl)
	fmt.Println(userInfo.NickName)
}

func TestGetUsersId2(t *testing.T) {
	//获取2400个用户id
	//存放的是数组类型的地址
	var allUIds []*([]string)
	ch := make(chan *([]string),100)
	path := "https://www.jianshu.com/recommendations/users?page="
	for i := 1; i <= 100; i++ {
		url := path + strconv.Itoa(i)
		//开100个协程，并行获取用户id
		go getUserIdSync(url,ch)
	}
	for i:=1;i<=100;i++  {
		allUIds = append(allUIds,<-ch)
	}
	fmt.Println(len(allUIds))
	fmt.Println(*allUIds[99])

}
func getUserIdSync(url string,ch chan  *([]string)){
	uids := GetUsersId(url)
	ch <- &uids
}
