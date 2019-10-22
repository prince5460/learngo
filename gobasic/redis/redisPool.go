package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当程序启动时初始化连接池
func init() {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //客户端与数据库的最大连接数,0表示没有限制
		IdleTimeout: 100, //最大空闲时间
	}
}

func main() {
	//从pool中取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "job", "coder")
	if err != nil {
		fmt.Println("Set err=", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "job"))
	if err != nil {
		fmt.Println("Get err=", err)
		return
	}
	fmt.Println("r=", r)
}
