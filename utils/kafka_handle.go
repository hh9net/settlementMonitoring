package utils

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaConn1 = "127.0.0.1:9092"
	kafkaConn2 = "127.0.0.1:9093"
	kafkaConn3 = "127.0.0.1:9094"
	topic1     = "test_kafka"
)

//代理
var brokerAddrs = []string{kafkaConn1, kafkaConn2, kafkaConn3}

//kafka生产数据
func KafkaProduction() {
	// read command line input
	reader := bufio.NewReader(os.Stdin)
	writer := newKafkaWriter(brokerAddrs, topic1)
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

var (
	topic2 = flag.String("t", "test_kafka", "kafka topic")
	group  = flag.String("g", "test-group", "kafka consumer group")
)

func main() {
	flag.Parse()
	config := kafka.ReaderConfig{
		Brokers:  []string{kafkaConn1, kafkaConn2, kafkaConn3},
		Topic:    *topic2,
		MinBytes: 1e3,
		MaxBytes: 1e6,
		GroupID:  *group,
	}
	reader := kafka.NewReader(config)
	ctx := context.Background()
	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {

			log.Printf("fail to get msg:%v", err)
			continue
		}
		log.Printf("msg content:topic=%v,partition=%v,offset=%v,content=%v",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			log.Printf("fail to commit msg:%v", err)
		}
	}
}
