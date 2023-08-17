package mutex

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

func Redisync() {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.60.34:6379",
	})
	pool := goredis.NewPool(client)

	rs := redsync.New(pool)
	mutexname := "test"
	mutex := rs.NewMutex(mutexname)
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("开始获取锁")
			mutex.Lock()
			fmt.Println("获取锁成功")
			time.Sleep(10 * time.Second)
			fmt.Println("开始释放锁")
			mutex.Unlock()
			fmt.Println("释放锁成功")
		}()
	}
	wg.Wait()
}
