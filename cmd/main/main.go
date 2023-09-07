package main

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/utils/excel2tag"
	"os"
)

func main() {
	file, err := os.Open("temp/lxy.xlsx")
	if err != nil {
		return
	}
	worker := excel2tag.New(file, excel2tag.WithCommentColumnIndex(11))
	worker.Parse()
	fmt.Println(worker.GetTagTypeList())
	fmt.Println(worker.GetTagList())
	fmt.Println(worker.GetWordList())
	fmt.Println(worker.GetWordTagLink())
}
