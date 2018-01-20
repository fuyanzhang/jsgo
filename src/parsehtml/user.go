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
	UserFollowing int32 /**关注数*/
	UserFollower  int32 /**被关注数*/
	ArticalNum int32    /**文章数*/
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

func GetUserInfo(url string) UserInfo{
	doc, err := goquery.NewDocument(url)
	if err != nil{
		log.Println("ERROR","get userInfo failed.",err)
	}
	
}