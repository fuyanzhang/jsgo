package parsehtml

import (
	"strconv"
	"fmt"
)

func MainProcess() {
	//获取所有uids
	allUids := getAllUids()
	urlPref :="https://www.jianshu.com/u/"
	i:=1
	for _,v := range allUids  {
		for _,innerV := range *v  {
			//fmt.Println(innerV)
			userInfo := GetUserInfo(urlPref+innerV)
			fmt.Println(userInfo)
			i++
		}
	}
	fmt.Println(i)

}

func getAllUids() []*([]string) {
	var allUIds []*([]string)
	ch := make(chan *([]string), 100)
	path := "https://www.jianshu.com/recommendations/users?page="
	for i := 1; i <= 100; i++ {
		url := path + strconv.Itoa(i)
		//开100个协程，并行获取用户id
		go getUserIdSync(url, ch)
	}
	for i := 1; i <= 100; i++ {
		allUIds = append(allUIds, <-ch)
	}
	return allUIds;
}
func getUserIdSync(url string, ch chan *([]string)) {
	uids := GetUsersId(url)
	ch <- &uids
}
