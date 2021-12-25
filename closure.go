/**
 @author:way
 @date:2021/11/24
 @note
**/

package main

import "fmt"


// Accumulate 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}
// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string , int)  {
	//血量一直为150
	hp := 150

	//创建返回的闭包
	return func() (string, int) {
		return name,hp
	}

}

func main() {

	str := "hello world"
	fmt.Println(str)
	foo := func() {
		str ="hello way"
		fmt.Println(str)
	}
	foo()

	accumulator  := Accumulate(1)

	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)


	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)


	//创建一个玩家生成器
	generator := playerGen("hign noon")
	name,hp := generator()
	println(name,hp)
}
