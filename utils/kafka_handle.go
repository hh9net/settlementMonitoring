package utils

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaConn1 = "127.0.0.1:9092"
	kafkaConn2 = "127.0.0.1:9093"
	kafkaConn3 = "127.0.0.1:9094"
	topic      = "test_kafka"
)

//代理
var brokerAddrs = []string{kafkaConn1, kafkaConn2, kafkaConn3}

//kafka生产数据
func KafkaProduction() {
	// read command line input
	reader := bufio.NewReader(os.Stdin)
	writer := newKafkaWriter(brokerAddrs, topic)
	defer writer.Close()
	for {
		fmt.Print("Enter msg: ")
		msgStr, _ := reader.ReadString('\n')

		msg := kafka.Message{
			Value: []byte(msgStr),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//消息分发策略默认使用轮训策略
func newKafkaWriter(kafkaURL []string, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: kafkaURL,
		Topic:   topic,
	})
}
