package parsehtml

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
)

func ParseAbs(url string) map[string]string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	abmap := make(map[string]string)
	doc.Find(".note-list li").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title").Text()
		path, _ := contentSelection.Find(".title").Attr("href")
		abmap[title] = doc.Url.Scheme + "://" + doc.Url.Host + path
	})
	return abmap
}

/**
获取该用户所有文章数
 */
func GetTotal(url string) int {
	doc, err := goquery.NewDocument(url)
	var totalCount int
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".main-top .info li").Each(func(i int, contentSelection *goquery.Selection) {
		//text := contentSelection.Find(".title").Text()
		//num,_ := contentSelection.Find(".title").Attr("href")
		if i == 2 {
			totalCount, _ = strconv.Atoi(contentSelection.Find(".meta-block a p").Text())

			return
		}

	})
	return totalCount
}
