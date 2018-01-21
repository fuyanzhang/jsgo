package parsehtml

import (
	"testing"
	"fmt"
	"strconv"
)

func TestGetUsersId(t *testing.T) {
	uids := GetUsersId("https://www.jianshu.com/recommendations/users?page=1")
	fmt.Println(uids)
	fmt.Println(len(*uids))
	fmt.Println(*uids)
	fmt.Println((*uids)[0])
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
	var allUIds []*[]string
	path := "https://www.jianshu.com/recommendations/users?page="
	for i:=1;i<=100 ;i++  {
		url := path+strconv.Itoa(i)
		uids := GetUsersId(url)
		allUIds[i] = uids

	}


}
