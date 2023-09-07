package excel2tag_test

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/utils/excel2tag"
	"os"
	"testing"
)

func TestExcel2Tag(t *testing.T) {
	file, err := os.Open("temp/lxy.xlsx")
	if err != nil {
		return
	}
	worker := excel2tag.New(file)
	worker.Parse()
	tagTypeSet := worker.GetTagTypeList()
	fmt.Println(tagTypeSet)
}
