/*******************************************************
 *  File        :   user.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/13 7:16 下午
 *  Notes       :
 *******************************************************/
package model

const (
	UserStatusOnline  = 1
	UserStatusOffline = iota
)

type User struct {
	UserId    int    `json:"user_id"`
	Passwd    string `json:"passwd"`
	NickName  string `json:"nick_name"`
	Sex       string `json:"sex"`
	Header    string `json:"header"`
	LastLogin string `json:"last_login"`
	Status    int    `json:"status"`
}
