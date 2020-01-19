/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/13 7:14 下午
 *  Notes       :
 *******************************************************/
package main

import "time"

func main() {
	initRedis("localhost:6379", 16, 1024, time.Second*300)
	initUSerMgr()
	runServer("0.0.0.0:10000")

}
