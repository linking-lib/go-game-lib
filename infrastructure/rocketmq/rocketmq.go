package rocketmq

import (
	"fmt"
	"github.com/alecthomas/log4go"
	mqHttpSdk "github.com/aliyunmq/mq-http-go-sdk"
	"github.com/gogap/errors"
	lkError "github.com/linking-lib/go-game-lib/lkerrors"
	"strings"
	"time"
)

type MqMsg struct {
	MessageBody string            // 消息内容
	MessageTag  string            // 消息标签
	Properties  map[string]string // 消息属性
	MessageKey  string            // 消息KEY
}

var DEBUG = true

/**
初始rocketmq
*/
func Init(config MqConfig, producers []MqTopicConfig, consumers []MqTopicConfig, debug bool) {
	// 初始客户端
	InitMq(config)
	// 初始生产者
	for i := range producers {
		InitProducer(producers[i])
	}
	// 初始消费者
	for i := range consumers {
		InitConsumer(consumers[i])
	}
	DEBUG = debug
}

/**
发送消息
*/
func PublishMsg(name string, msg MqMsg) (bool, error) {
	mqMsg := mqHttpSdk.PublishMessageRequest{
		MessageBody: msg.MessageBody,
		MessageTag:  msg.MessageTag,
		Properties:  msg.Properties,
		MessageKey:  msg.MessageKey,
	}
	cli := GetProducer(name)
	if cli == nil {
		return false, lkError.ErrWrongRocketmqProducer
	}
	_, err := cli.PublishMessage(mqMsg)
	if err != nil {
		_ = log4go.Error("rocketmq PublishMsg error=========" + err.Error())
		return false, err
	} else {
		return true, nil
	}
}

func ConsumeMsg(name string, numOfMessages int32, waitSeconds int64, consumeFunc func(msg MqMsg) (bool, error)) {
	cli := GetConsumer(name)
	if cli == nil {
		log4go.Debug("rocketmq: consumer not found")
		return
	}
	for {
		endChan := make(chan int)
		respChan := make(chan mqHttpSdk.ConsumeMessageResponse)
		errChan := make(chan error)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("捕获到的错误：%s\n", r)
				}
			}()
			select {
			case resp := <-respChan:
				{
					// 处理业务逻辑
					var handles []string
					if DEBUG {
						log4go.Debug("Consume %d messages---->\n", len(resp.Messages))
					}
					for _, v := range resp.Messages {
						if DEBUG {
							log4go.Debug("\tMessageID: %s, PublishTime: %d, MessageTag: %s\n"+
								"\tConsumedTimes: %d, FirstConsumeTime: %d, NextConsumeTime: %d\n"+
								"\tBody: %s\n"+
								"\tProps: %s\n",
								v.MessageId, v.PublishTime, v.MessageTag, v.ConsumedTimes,
								v.FirstConsumeTime, v.NextConsumeTime, v.MessageBody, v.Properties)
						}
						msg := MqMsg{MessageBody: v.MessageBody, MessageTag: v.MessageTag,
							MessageKey: v.MessageKey, Properties: v.Properties}
						_, err := consumeFunc(msg)
						if err != nil {
							_ = log4go.Error("rocketmq ConsumeMsg error=========" + err.Error())
						} else {
							handles = append(handles, v.ReceiptHandle)
						}
					}
					// NextConsumeTime前若不确认消息消费成功，则消息会重复消费
					// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样
					ackErr := cli.AckMessage(handles)
					if ackErr != nil {
						// 某些消息的句柄可能超时了会导致确认不成功
						_ = log4go.Error("rocketmq AckMessage error=========" + ackErr.Error())
						for _, errAckItem := range ackErr.(errors.ErrCode).Context()["Detail"].([]mqHttpSdk.ErrAckItem) {
							log4go.Info("\tErrorHandle:%s, ErrorCode:%s, ErrorMsg:%s\n",
								errAckItem.ErrorHandle, errAckItem.ErrorCode, errAckItem.ErrorMsg)
						}
						time.Sleep(time.Duration(3) * time.Second)
					} else {
						if DEBUG {
							log4go.Debug("Ack ---->\n\t%s\n", handles)
						}
					}
					endChan <- 1
				}
			case err := <-errChan:
				{
					// 没有消息
					if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
						if DEBUG {
							log4go.Debug("\nNo new message, continue!")
						}
					} else {
						_ = log4go.Error("rocketmq errChan error=========" + err.Error())
						time.Sleep(time.Duration(3) * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					log4go.Info("Timeout of consumer message ??")
					endChan <- 1
				}
			}
		}()
		// 长轮询消费消息
		// 长轮询表示如果topic没有消息则请求会在服务端挂住3s，3s内如果有消息可以消费则立即返回
		cli.ConsumeMessage(respChan, errChan,
			numOfMessages, // 一次最多消费多少条(最多可设置为16条)
			waitSeconds,   // 长轮询时间多少秒（最多可设置为30秒）
		)
		<-endChan
	}
}
