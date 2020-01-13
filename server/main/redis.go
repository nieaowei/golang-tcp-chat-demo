/*******************************************************
 *  File        :   redis.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/13 6:54 下午
 *  Notes       :	Use redis in client and
					init redis connection pool.
 *******************************************************/
package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

//InitRedis is to inital the Redis connection pool.
func InitRedis(addr string, idleConnNum, maxConnNum int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", addr)
		},
		TestOnBorrow:    nil,
		MaxIdle:         idleConnNum, //the max idel connection number.
		MaxActive:       maxConnNum,  //the max active connection number.
		IdleTimeout:     idleTimeout, //the idel connection timeout.
		Wait:            false,
		MaxConnLifetime: 0,
	}
	return
}

//GetConn is to return accessiable connection in redis connection pool.
func GetConn() redis.Conn {
	return pool.Get()
}

func PutConn(conn redis.Conn) {
	conn.Close()
}
