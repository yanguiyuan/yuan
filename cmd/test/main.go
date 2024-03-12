package main

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/utils/excel2tag"
	"os"
)

func main() {
	file, err := os.Open("temp/流行语分类终.xlsx")
	if err != nil {
		return
	}
	worker := excel2tag.New(excel2tag.WithCommentColumnIndex(11), excel2tag.WithPrimaryKeyColumnIndex(1))
	err = worker.Parse(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	tagTypeSet := worker.GetTagTypeSet()
	fmt.Println(tagTypeSet)
	list := worker.GetTagSet()
	fmt.Println(list[0])
}
