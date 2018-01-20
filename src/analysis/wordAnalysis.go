package analysis

import (
	"github.com/huichen/sego"
	"fmt"
)

func Analysis(input string) {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("F:/go/src/goer/src/github.com/huichen/sego/data/dictionary.txt")
	segments := segmenter.Segment([]byte(input))

	//fmt.Println(sego.SegmentsToString(segments, false))
	output := sego.SegmentsToSlice(segments,true)
	for k,v := range output{
		fmt.Printf("the key is %v ,the value is %v \n",k,v)

	}
}
