package parsehtml

import (
	"testing"
	"fmt"
)

func TestGetUsersId(t *testing.T) {
	uids := GetUsersId("https://www.jianshu.com/recommendations/users?page=1")
	fmt.Println(len(uids))
	fmt.Println(uids[0])
}

func TestGetUserInfo(t *testing.T) {
	userInfo := GetUserInfo("https://www.jianshu.com/u/dbfdce352c0d")
	fmt.Println(userInfo.ArticalNum)
	fmt.Println(userInfo.UserFollower)
	fmt.Println(userInfo.UserFollowing)
	fmt.Println(userInfo.UserUrl)
	fmt.Println(userInfo.NickName)
}
