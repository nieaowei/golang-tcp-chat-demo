/*******************************************************
 *  File        :   mgr.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/15 3:52 上午
 *  Notes       :
 *******************************************************/
package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var UserTable = "users"

type UserMgr struct {
	pool *redis.Pool
}

//Initialize a users manager.
func NewUserMgr(pool *redis.Pool) (mgr *UserMgr) {
	mgr = &UserMgr{
		pool: pool,
	}
	return
}

//Get user's information in database.
func (p *UserMgr) getUser(conn redis.Conn, id int) (user *User, err error) {
	result, err := redis.String(conn.Do("hget", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorUserNotExist
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return
	}
	return
}

//Log in with users account and password.
func (p *UserMgr) Login(id int, passwd string) (user *User, err error) {
	conn := p.pool.Get()
	defer conn.Close()

	user, err = p.getUser(conn, id)
	if err != nil {
		return
	}
	if id != user.UserId && passwd != user.Passwd {
		err = ErrorInvalidPasswd
		return
	}
	user.Status = UserStatusOnline
	user.LastLogin = time.Now().String()
	return
}

//Register in with users info.
func (p *UserMgr) Register(user *User) (err error) {
	conn := p.pool.Get()
	defer conn.Close()

	if user == nil {
		fmt.Println("invalid user")
		err = ErrorInvalidParams
		return
	}

	_, err = p.getUser(conn, user.UserId)
	if err != nil {
		err = ErrorUserExist
		return
	}
	if err != ErrorUserNotExist {
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println(UserTable, fmt.Sprintf("%d", user.UserId), string(data))

	_, err = conn.Do("Hset", UserTable, fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		return
	}
	return
}
