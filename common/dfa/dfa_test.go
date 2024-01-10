package dfa

import (
	"fmt"
	"testing"
)

func TestDfa(t *testing.T) {
	filterManager := NewFilter(
		StoreOption{
			Type: StoreMemory,
		},
		FilterOption{
			Type: FilterDfa,
		},
	)

	// 加载字典

	err := filterManager.GetStore().LoadDictPath("dict/default_dict.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(filterManager.GetFilter().IsSensitive("操你阿玛"))
	fmt.Println(filterManager.GetFilter().IsSensitive("你好"))
}
