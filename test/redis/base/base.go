package base

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func Redis() {
	cli := redis.NewClient(&redis.Options{
		Addr:         "192.168.60.34:6379",
		DB:           0,
		Password:     "",
		WriteTimeout: 600 * time.Millisecond, //写超时时间
		ReadTimeout:  300 * time.Millisecond, //读超时时间
		DialTimeout:  3 * time.Minute,        //连接超时时间
		PoolSize:     10,                     //最大连接数
		MinIdleConns: 3,                      //最小空闲连接数
		IdleTimeout:  1 * time.Minute,        //空闲连接超时时间
	})
	defer cli.Close()
	fmt.Println(cli)
	// Set
	if err := cli.Set("k1", "v1", 24*time.Hour).Err(); err != nil {
		fmt.Println(err)
	}
	// Get
	if val, err := cli.Get("k1").Result(); err == nil {
		fmt.Println(val) // v1
	}
	// 设置key值，返回旧值
	old, _ := cli.GetSet("k1", "v2").Result()
	fmt.Println(old) // v1
	// 若key不存在，设置key 0没有过期时间 做分布式锁-原子操作
	cli.SetNX("k2", "v2", 0)
	if val, err := cli.Get("k2").Result(); err == nil {
		fmt.Println(val) // v2
	}
	// 批量获取
	if vals, err := cli.MGet("k1", "k2").Result(); err == nil {
		fmt.Println("vals:", vals) // [v1,v2]
	}
	// 批量Set
	if _, err := cli.MSet("k3", "v3", "k4", "v4").Result(); err == nil {
		vals, _ := cli.MGet("k3", "k4").Result()
		fmt.Println(vals...) // v3 v4
	}
	// 每次为值累加一 仅限整数
	if _, err := cli.Set("k5", "1", 0).Result(); err == nil {
		val, _ := cli.Incr("k5").Result()
		fmt.Println(val) // 2
	}
	// 自定义累加
	if _, err := cli.Set("k6", "1", 0).Result(); err == nil {
		val, _ := cli.IncrBy("k6", 10).Result()
		fmt.Println(val) // 11
	}
	// 浮点数
	if _, err := cli.Set("k7", "1.0", 0).Result(); err == nil {
		val, _ := cli.IncrByFloat("k7", 7.6).Result()
		fmt.Println(val) // 8.6
	}
	// 累减
	if _, err := cli.Set("k8", "1", 0).Result(); err == nil {
		val, _ := cli.DecrBy("k8", 8).Result()
		fmt.Println(val) // -7
	}
	// 删除
	cli.Del("k8")
	if _, err := cli.Get("k8").Result(); err != nil {
		fmt.Println(err) // nil
	}
	// 设置过期时间
	cli.Expire("k1", 3*time.Minute)
	// 获取过期时间
	if val, err := cli.TTL("k1").Result(); err == nil {
		fmt.Println(val) // 3m
	}
}
