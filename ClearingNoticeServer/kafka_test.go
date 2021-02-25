package ClearingNoticeServer

import (
	"testing"
)

func TestClearingNoticeConsumer(t *testing.T) {
	KafkaIp = "192.168.128.222:9092"
	KafkaTopic = "com.etc.clear.Notice"
	ClearingNoticeConsumer()
}

func TestProducer(t *testing.T) {
	KafkaIp = "192.168.128.222:9092"
	KafkaTopic = "com.etc.clear.Notice"

	//for {
	ClearingNotice("2021-01-24", "565691")
	//time.Sleep(time.Second * 20)
	//}

}

func TestConsumer(t *testing.T) {
	//KafkaIp = "192.168.128.222:9092"
	//KafkaTopic = "com.etc.clear.Notice"
	//for {
	//	logrus.Print("消费数据")
	//	Consumer()
	//	time.Sleep(time.Second * 10)
	//}

}
func TestUseKafka(t *testing.T) {
	//UseKafka()
	//2.12-2.5.0
	//	consumer groups require Version to be >= V0_10_2_0
}
