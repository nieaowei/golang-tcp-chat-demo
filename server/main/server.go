/*******************************************************
 *  File        :   server.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/15 4:46 上午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"net"
)

//Run the tcp server.
//Listen to the client request.
func runServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen failed")
		return
	}
	//Wait client-side connection.
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}
		go proccess(conn)
	}
}

func proccess(conn net.Conn) {
	defer conn.Close() // Close the connection.
	client := &Client{
		conn: conn,
	}
	err := client.Process()
	if err != nil {
		fmt.Println("client process failed.")
		return
	}
}
