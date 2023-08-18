package base

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func Producer() {
	// 连接到客户端
	pro, err := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.60.120:9876"}))
	if err != nil {
		panic(err)
	}
	// 启动 producer
	err = pro.Start()
	if err != nil {
		panic(err)
	}
	// 发送消息 同步
	res, err := pro.SendSync(context.Background(), &primitive.Message{
		Topic: "test",
		Body:  []byte("hello world"),
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("发送成功：%s\n", res.String())
	}
	_ = pro.Shutdown()
}

func Consumer() {
	// 连接到客户端 Push服务器主动推 Pull主动向服务器请求
	c, _ := rocketmq.NewPushConsumer(consumer.WithNameServer([]string{"192.168.60.120:9876"}), consumer.WithGroupName("test"))
	if err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context, me ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range me {
			fmt.Printf("获取：%v\n", me[i])
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		panic(err)
	}
	_ = c.Start()
	// 不能让主协程退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}

// 延迟消息
func Delay() {
	// 连接到客户端
	pro, err := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.60.120:9876"}))
	if err != nil {
		panic(err)
	}
	// 启动 producer
	err = pro.Start()
	if err != nil {
		panic(err)
	}
	msg := primitive.NewMessage("test", []byte("延迟消息测试"))
	// 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	msg.WithDelayTimeLevel(3) // 设置延迟消息
	if res, err := pro.SendSync(context.Background(), msg); err == nil {
		fmt.Println("发送消息：", res.String())
	}
	_ = pro.Shutdown()
}

type TestListener struct{}

/*
ExecuteLocalTransaction(*Message) LocalTransactionState
// When no response to prepare(half) message. broker will send check message to check the transaction status, and this
// method will be invoked to get local transaction status.
CheckLocalTransaction(*MessageExt) LocalTransactionState
*/
/* 实现事务listener接口 */
// 执行
func (t *TestListener) ExecuteLocalTransaction(ms *primitive.Message) primitive.LocalTransactionState {
	// 本地执行逻辑
	/*fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second * 3)
	fmt.Println("执行逻辑成功")*/

	// return primitive.CommitMessageState // 本地执行逻辑成功，commit 消息不会执行回查

	// fmt.Println("开始执行本地逻辑")
	// time.Sleep(time.Second * 3)
	// fmt.Println("执行逻辑失败")
	// return primitive.RollbackMessageState // 本地执行逻辑失败，不会回查

	fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second * 3)
	fmt.Println("执行逻辑失败")
	// 场景 本地逻辑失败未知/代码异常
	return primitive.UnknowState // 本地执行逻辑出现问题，回查
}

// 回查 不断轮循同组其他producer
func (t *TestListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	// 回查逻辑
	fmt.Println("回查")
	time.Sleep(time.Second * 5)
	return primitive.RollbackMessageState // 回滚
}

// 事务消息
func Transaction() {
	// 连接到客户端 事务producer
	pro, err := rocketmq.NewTransactionProducer(
		&TestListener{},
		producer.WithNameServer([]string{"192.168.60.120:9876"}))
	if err != nil {
		panic(err)
	}
	// 启动 producer
	err = pro.Start()
	if err != nil {
		panic(err)
	}
	// 发送事务消息
	if res, err := pro.SendMessageInTransaction(context.Background(), primitive.NewMessage("tx-test", []byte("事务消息测试"))); err == nil {
		fmt.Println("发送成功：", res.String())
	}
	// 阻塞 回查不是实时性的
	time.Sleep(time.Minute)
	_ = pro.Shutdown()
}
