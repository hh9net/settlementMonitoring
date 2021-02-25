package ClearingNoticeServer

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	KafkaIp    string
	KafkaTopic string
	wg         sync.WaitGroup
)

//生产数据
func Producer(msgdata *KafKaClearingNoticeMsg) {
	config := sarama.NewConfig()
	// 等待服务器 所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 使用给定代理地址和配置创建一个同步生产者
	//producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	producer, err := sarama.NewSyncProducer([]string{KafkaIp}, config)
	defer func() {
		_ = producer.Close()
	}()
	if err != nil {
		log.Println("sarama.NewSyncProducer errpr:", err)
		return
	}

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		//Topic: "test",//包含了消息的主题
		Topic:     KafkaTopic,
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}
	//fmt.Println("清分通知 构建发送的消息",msg)
	log.Printf("发送数据:%#v", msgdata)
	//将字符串转换为字节数组
	hdvalue, err := json.Marshal(msgdata)
	if err != nil {
		log.Println("KafKaClearingNoticeMsg json.Unmarshal error:", err)
		return
	}
	//log.Println("dvalue", string(hdvalue))
	msg.Value = sarama.ByteEncoder(hdvalue)
	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Send message Fail", err)
	}
	log.Printf("Partition = %d, offset=%d\n", partition, offset)
}

var brokers = []string{"172.18.70.21:9092"} //不参与编译
var topics = []string{KafkaTopic}           //不参与编译
var group = "39"                            //不参与编译

type consumerGroupHandler struct {
	name string //groupname
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

//消费主张
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("%s group Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		//消息处理msg.Value
		data := new(KafKaClearingNoticeMsg)
		umerr := json.Unmarshal(msg.Value, data)
		if umerr != nil {
			log.Println("处理清分通知消息json.Unmarshal error")
			return umerr
		}

		log.Println("消息处理完成", data.ClearingNum, data.ClearingDate)
	}
	return nil
}

//处理消息  msg 消息数据
func ProcessMessage(topic string, msg []byte) (error, string) {
	log.Println("正执行处理消息:ProcessMessage【topic,msg[0:10]的值】 :", topic, string(msg[0:10]))
	return nil, ""
}

//消费  group name == c1
func consume(group *sarama.ConsumerGroup, wg *sync.WaitGroup, name string) error {
	log.Println(name+" group "+"start ok kafka consume name is", name)
	wg.Done()
	ctx := context.Background()
	for {
		var topics []string
		topics = append(topics, KafkaTopic)
		log.Println("kafka的topics:", topics)
		//name c1
		handler := consumerGroupHandler{name: name}
		//消费
		err := (*group).Consume(ctx, topics, handler)
		if err != nil {
			log.Println("(*group).Consume  error", err)
			return err
		}
		log.Println("kafka 处理 ok")
	}
}

//main 调用 消费kafka
func ConsumerGroup() error {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0
	client, err := sarama.NewClient([]string{KafkaIp}, config)
	defer func() {
		_ = client.Close()
	}()
	if err != nil {
		log.Println("sarama.NewClient 执行出错: ", err)
		return err
	}
	//c1 组
	group1, err := sarama.NewConsumerGroupFromClient("c1", client)
	defer func() {
		_ = group1.Close()
	}()
	if err != nil {
		log.Println("sarama.NewConsumerGroupFromClient 执行出错:", err)
		return err
	}
	wg.Add(1)
	//处理kafka数据
	cerr := consume(&group1, &wg, "c1")
	if cerr != nil {
		return cerr
	}
	//go consume(&group2,&wg,"c2")
	//go consume(&group3,&wg,"c3")
	wg.Wait()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
	return nil
}

func handleErrors(group *sarama.ConsumerGroup, wg *sync.WaitGroup) {
	wg.Done()
	for err := range (*group).Errors() {
		log.Println("ERROR", err)
	}
}
func UseKafka() {
	k := NewKafka()
	f := k.Init()
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	f()
}

//本机测试消费者
func Consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{KafkaIp}, config)
	if err != nil {
		log.Println("sarama.NewConsumer error:", err)
		return
	}
	//conf := &sarama.Config{}
	ConsumerGroup, grouperr := sarama.NewConsumerGroup([]string{KafkaIp}, "", nil)
	if grouperr != nil {
		log.Println("sarama.NewConsumerGroup error:", grouperr)
		return
	}
	log.Println(ConsumerGroup)
	//ConsumerGroup.Consume()
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions(KafkaTopic)
	if err != nil {
		log.Println(err)
		return
	}
	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition(KafkaTopic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Println(err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				log.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	_ = consumer.Close()
}

type Kafka struct {
	brokers []string
	topics  []string
	//OffsetNewest int64 = -1
	//OffsetOldest int64 = -2
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
}

func NewKafka() *Kafka {
	return &Kafka{
		brokers:           brokers,
		topics:            topics,
		group:             group,
		channelBufferSize: 2,
		ready:             make(chan bool),
		version:           "1.1.1",
	}
}

func (p *Kafka) Init() func() {
	log.Println("kafka init...")

	version, err := sarama.ParseKafkaVersion(p.version)
	if err != nil {
		log.Println("Error parsing Kafka version:  ", err)
	}
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = -2                                   // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = p.channelBufferSize                         // channel长度

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(p.brokers, p.group, config)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++++++++++Error creating consumer group client: ", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			//util.HandlePanic("client.Consume panic", log.StandardLogger())
		}()
		for {
			if err := client.Consume(ctx, p.topics, p); err != nil {
				log.Println("Error from consumer: ", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println("ctx.Err():", ctx.Err())
				return
			}
			p.ready = make(chan bool)
		}
	}()
	<-p.ready
	log.Infoln("Sarama consumer up and running!...")
	// 保证在系统退出时，通道里面的消息被消费
	return func() {
		log.Println("+++++++++++++++++++++++++++++kafka close")
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			log.Println("++++++++++++++++++++++++++++++++++Error closing client: ", err)
		}
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (p *Kafka) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(p.ready)
	return nil
}

//清理
// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (p *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

//消费主张
// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (p *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	// 具体消费消息
	for message := range claim.Messages() {
		msg := string(message.Value)
		log.Println("+++++++++++++++++++++++++++++++msg:", msg)
		time.Sleep(time.Second)
		//run.Run(msg)
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}
