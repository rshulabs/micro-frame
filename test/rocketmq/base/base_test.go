package base

import "testing"

func TestProducer(t *testing.T) {
	Producer()
}

func TestConsumer(t *testing.T) {
	Consumer()
}

// 延迟消息
func TestDelay(t *testing.T) {
	Delay()
}

// 事务消息
func TestTx(t *testing.T) {
	Transaction()
}
