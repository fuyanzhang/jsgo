package parsehtml

import (
	"strconv"
	"fmt"
	"sync"
)

func MainProcess() {
	//获取所有uids
	allUids := getAllUids()
	urlPref := "https://www.jianshu.com/u/"
	//获取所有用户信息
	var userInfos []UserInfo
	var channel = make(chan *UserInfo, 50)
	for _, v := range allUids {
		go func() {
			for _, innerV := range *v {
				userInfo, err := GetUserInfo(urlPref + innerV)
				if err != nil {
					continue
				}
				userInfo.UserId = innerV

				channel <- userInfo
			}
		}()

	}

	for _, v := range allUids {
		for range *v {
			userInfos = append(userInfos, *(<-channel))
		}

	}
	fmt.Println(len(userInfos))
	fmt.Println(userInfos)

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

//获取文章详情
func getUserArtiDetail(userinfo UserInfo) map[string]string {
	total := userinfo.ArticalNum
	//分页
	var pageSize int = 9
	var totalPage int

	if total%pageSize == 0 {
		totalPage = total / pageSize
	} else {
		totalPage = (total / pageSize) + 1
	}
	var wg sync.WaitGroup
	var resultSet [] map[string]string

	for i := 1; i <= totalPage; i++ {
		url := userinfo.UserUrl + "?order_by=shared_at&page=" + strconv.Itoa(i)
		wg.Add(1)
		go func() {
			resultSet = append(resultSet, ParseAbs(url))
			defer wg.Done()
		}()
	}
	wg.Wait()
	detailArtil :=make(map[string]string)
	for _, result := range resultSet {
		for k, v := range result {
			rs := ParseDetail(v)
			detailArtil[k] = rs
		}
	}
	return detailArtil
}

//通过用户发表的文章，提取用户关键字，关键字为20个
func getUserKeyWord(detailArtil map[string]string) []string{

	return nil

}
