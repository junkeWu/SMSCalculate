package main

import (
	"fmt"

	util "github.com/junkeWu/SMSCalculate/server"
)

func main() {
	s1 := "hello word!2020 street 188#"
	unitCount := util.CalculateContent(s1) // 返回标准的多少单位
	fmt.Println("单位数量为：", unitCount)
	// todo 做计算
}
