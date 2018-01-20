package main

import (
	"fmt"
	"parsehtml"
	"strconv"
	"sync"
)

func main() {
	total := parsehtml.GetTotal("https://www.jianshu.com/u/b52ff888fd17")
	fmt.Println(total)

	var pageSize int = 9
	var totalPage int

	if total%pageSize == 0 {
		totalPage = total / pageSize
	} else {
		totalPage = (total / pageSize) + 1
	}
	fmt.Println(totalPage)
	var wg sync.WaitGroup
	var resultSet [] map[string]string

	for i := 1; i <= totalPage; i++ {
		url := "https://www.jianshu.com/u/b52ff888fd17?order_by=shared_at&page=" + strconv.Itoa(i)
		wg.Add(1)
		go func() {
			resultSet = append(resultSet, parsehtml.ParseAbs(url))
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(resultSet))

	//获取文章详情
	for _, result := range resultSet {
		for _, v := range result {

			rs := parsehtml.ParseDetail(v)
			fmt.Println(rs)

		}
	}
}

