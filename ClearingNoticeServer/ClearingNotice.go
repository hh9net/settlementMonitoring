package ClearingNoticeServer

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type KafKaClearingNoticeMsg struct {
	ClearingNum  string `json:"clearing_num"`  //清分消息包号
	ClearingDate string `json:"clearing_date"` //清分目标日
}

//从kafka 清分通知消息发送
func ClearingNotice(ClearingDate, ClearingNum string) {
	data := new(KafKaClearingNoticeMsg)
	data.ClearingDate = ClearingDate
	data.ClearingNum = ClearingNum
	Producer(data)
}

//清分通知消息消费 测试使用
func ClearingNoticeConsumer() {
KafkaI:
	log.Println("启动kafka")
	//处理kafka数据
	err := ConsumerGroup()
	if err != nil {
		log.Println("执行ConsumerGroup() 处理kafka数据 error :", err)
		time.Sleep(time.Second * 3)
		goto KafkaI
	}
}
