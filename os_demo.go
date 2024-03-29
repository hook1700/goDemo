/**
 @author:way
 @date:2021/11/24
 @note
**/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	println(os.Hostname())
	println(os.Environ())
	println(os.Getpid())
	println(os.Getwd())


	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}

}
