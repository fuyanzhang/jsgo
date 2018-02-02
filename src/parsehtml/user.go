package parsehtml

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

/**
 */
type UserInfo struct {
	UserId string  /** 用户IDeg."dbfdce352c0d"*/
	NickName string /** 用户昵称 eg." 遛遛心情的溜妈"*/
	UserUrl string  /** 用户主页url eg "https://www.jianshu.com/u/dbfdce352c0d"*/
	UserFollowing string  /**关注数*/
	UserFollower  string /**被关注数*/
	ArticalNum string    /**文章数*/
	AritcalList []AbsTitle
	Desc string  /**用户描述*/
}

type AbsTitle struct {
	ArticleUrl string
	Title      string
}

func GetUsersId(url string) []string{
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println("ERROR","get uid failed.",err)
	}
	var userids []string
	doc.Find(".container .wrap").Each(func(i int, contentSelection *goquery.Selection) {
		path, _ := contentSelection.Find("a").Attr("href")
		uid := path[7:]
		userids = append(userids,uid)

	})
	return userids
}
//get user info ,the param url is user home url ,eg https://www.jianshu.com/u/dbfdce352c0d
func GetUserInfo(url string) *UserInfo{
	doc, err := goquery.NewDocument(url)
	if err != nil{
		log.Println("ERROR","get userInfo failed.",err)
	}
	userInfo := new(UserInfo)
	nickName :=doc.Find(".main-top .title").Find("a").Text()
	userInfo.NickName = nickName
	userInfo.UserUrl = url
	doc.Find(".main-top .info li").Each(func(i int, selection *goquery.Selection) {
		temp := selection.Find("p").Text()
		switch i {
		case 0: userInfo.UserFollowing = temp
		case 1: userInfo.UserFollower = temp
		case 2: userInfo.ArticalNum = temp
		}
	})
	return userInfo

}