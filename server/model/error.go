/*******************************************************
 *  File        :   error.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/15 4:22 上午
 *  Notes       :
 *******************************************************/
package model

import "errors"

// Customized errors.
var (
	ErrorUserNotExist  = errors.New("user not exist")
	ErrorUserExist     = errors.New("user exist.")
	ErrorInvalidPasswd = errors.New("password is error.")
	ErrorInvalidParams = errors.New("invalid params.")
)
