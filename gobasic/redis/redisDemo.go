package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dail err=", err)
		return
	}
	defer conn.Close() //关闭

	//testString(err, conn)

	//testMString(err, conn)

	//testHash(err, conn)

	//testMHash(err, conn)

	//testExpire(err, conn)

	//testList(err, conn)

	testSet(err, conn)

}

//操作Set
func testSet(err error, conn redis.Conn) {
	_, err = conn.Do("SAdd", "zoo", "cat", "dog", "pig")
	if err != nil {
		fmt.Println("SAdd err=", err)
		return
	}

	_, err = conn.Do("SRem", "zoo", "pig")
	if err != nil {
		fmt.Println("SRem err=", err)
		return
	}

	r, err := redis.Strings(conn.Do("SMembers", "zoo"))
	if err != nil {
		fmt.Println("SMembers err=", err)
		return
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}

//操作List
func testList(err error, conn redis.Conn) {
	_, err = conn.Do("LPush", "fruits", "apple", "orange", "banana")
	if err != nil {
		fmt.Println("LPush err=", err)
		return
	}
	r, err := redis.Strings(conn.Do("LRange", "fruits", 0, -1))
	if err != nil {
		fmt.Println("LRange err=", err)
		return
	}
	for _, v := range r {
		fmt.Println(v)
	}
}

//设置有效时间
func testExpire(err error, conn redis.Conn) {
	_, err = conn.Do("expire", "name", 10)
	if err != nil {
		fmt.Println("expire err=", err)
		return
	}
}

//批量操作Hash
func testMHash(err error, conn redis.Conn) {
	_, err = conn.Do("HMSet", "user2", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMSet err=", err)
		return
	}

	r, err := redis.Strings(conn.Do("HMGet", "user2", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err=", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d]=%v\n", i, v)
	}
}

//操作Hash
func testHash(err error, conn redis.Conn) {
	//向redis写入数据
	_, err = conn.Do("HSet", "user1", "name", "tom")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	_, err = conn.Do("HSet", "user1", "age", 18)
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//从redis中读取数据
	r1, err := redis.String(conn.Do("HGet", "user1", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	r2, err := redis.Int(conn.Do("HGet", "user1", "age"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Printf("r1 name=%v,r1 age=%v\n", r1, r2)
}

//批量操作String
func testMString(err error, conn redis.Conn) {
	_, err = conn.Do("MSet", "name", "James", "age", 20, "address", "beijing")
	if err != nil {
		fmt.Println("MSet err=", err)
		return
	}

	strings, err := redis.Strings(conn.Do("MGet", "name", "age", "address"))
	if err != nil {
		fmt.Println("MGet err=", err)
		return
	}
	for _, v := range strings {
		fmt.Println(v)
	}
}

//操作String
func testString(err error, conn redis.Conn) {
	//向redis写入数据string[k-v]
	_, err = conn.Do("Set", "name", "tom&jerry")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//从redis中读取数据string[k-v]
	reply, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("success:", reply)
}
