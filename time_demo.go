/**
 @author:way
 @date:2021/11/24
 @note
**/

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("当前时间%v\n",now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)
}
