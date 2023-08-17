package tx

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Tx() {
	cli := redis.NewClient(&redis.Options{
		Addr: "192.168.60.34:6379",
	})
	defer cli.Close()
	p := cli.TxPipeline()
	// 执行多个命令
	p.Get(context.Background(), "k2")
	p.Set(context.Background(), "k3", "444", -1)
	if _, err := p.Del(context.Background(), "k1").Result(); err != nil {
		fmt.Println(err)
	}
	_, err := p.Exec(context.Background())
	if err != nil {
		fmt.Println("提交事务失败", err)
	} else {
		fmt.Println("提交成功")
	}
	if val, err := cli.Get(context.Background(), "k3").Result(); err == nil {
		fmt.Println(val)
	}
}
