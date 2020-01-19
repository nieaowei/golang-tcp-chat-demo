/*******************************************************
 *  File        :   user_mgr.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/15 4:50 上午
 *  Notes       :
 *******************************************************/
package main

import "golang-tcp-chat-demo/server/model"

var (
	mgr *model.UserMgr
)

func initUSerMgr() {
	mgr = model.NewUserMgr(pool)
}
