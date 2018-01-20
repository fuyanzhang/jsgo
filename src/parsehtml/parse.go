package parsehtml

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	//"analysis"
)

func ParseDetail(url string) string{
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("open failed.%v",err)
	}
	sele := doc.Find(".show-content")
	fmt.Println(sele.Text())
	//analysis.Analysis(sele.Text())
	return sele.Text()
}
