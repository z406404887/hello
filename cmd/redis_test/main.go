package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
)

func main() {
	host := "192.168.3.33:6379"
	redisPool := &redis.Pool{
		MaxIdle:   10,
		MaxActive: 100,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialPassword(""), redis.DialDatabase(0))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	conn := redisPool.Get()
	conn.Do("SET", "Hello", "World")
	conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		c := redisPool.Get()
		reply, err := c.Do("GET", "Hello")
		if err != nil {
			log.Printf("get reply error, %s", err)
			return
		}

		log.Printf("the valud of 'Hello' is %s", reply)
	}(&wg)
	wg.Wait()
}
