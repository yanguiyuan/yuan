package main

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/datastruct/vec"
)

func main() {
	//file, err := os.Open("temp/lxy.xlsx")
	//if err != nil {
	//	return
	//}
	//worker := excel2tag.New(file, excel2tag.WithCommentColumnIndex(11))
	//worker.Parse()
	//fmt.Println(worker.GetTagTypeList())
	//fmt.Println(worker.GetTagList())
	//fmt.Println(worker.GetWordList())
	//fmt.Println(worker.GetWordTagLink())
	v := vec.FromSlice([]int{1, 2, 3, 4, 5})
	res := v.Iter().Unfold(0, func(result *int, e int) {
		*result += e
	})
	fmt.Println(res)
}
