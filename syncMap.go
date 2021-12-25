/**
 @author:way
 @date:2021/11/24
 @note
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map

	scene.Store("way",97)
	scene.Store("kay","98")
	scene.Store("jay","99")
	data,_ := scene.Load("way")
	fmt.Printf("%T\n",data)
	fmt.Println(scene.Load("way"))

	scene.Delete("kay")
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key,value)
		return true
	})
}

