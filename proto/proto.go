/*******************************************************
 *  File        :   proto.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/19 5:01 下午
 *  Notes       :
 *******************************************************/
package proto

import "golang-tcp-chat-demo/server/model"

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int    `json:"id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User model.User `json:"user"`
}

type LoginCmdRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
