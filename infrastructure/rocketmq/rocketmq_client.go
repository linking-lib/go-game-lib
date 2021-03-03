package rocketmq

import (
	mqHttpSdk "github.com/aliyunmq/mq-http-go-sdk"
)

type MqConfig struct {
	Endpoint   string // 设置HTTP接入域名（此处以公共云生产环境为例）
	AccessKey  string // AccessKey 阿里云身份验证，在阿里云服务器管理控制台创建
	SecretKey  string // SecretKey 阿里云身份验证，在阿里云服务器管理控制台创建
	Group      string // 您在控制台创建的 Consumer ID(Group ID)
	InstanceId string // Topic所属实例ID，默认实例为空
}

type MqTopicConfig struct {
	Name       string // 客户端名字
	Topic      string // 所属的 Topic
	Group      string // 您在控制台创建的 Consumer ID(Group ID)
	InstanceId string // Topic所属实例ID，默认实例为空
}

var (
	producers map[string]mqHttpSdk.MQProducer
	consumers map[string]mqHttpSdk.MQConsumer
	client    mqHttpSdk.MQClient
)

func InitMq(mqConf MqConfig) {
	client = mqHttpSdk.NewAliyunMQClient(mqConf.Endpoint, mqConf.AccessKey, mqConf.SecretKey, "")
	producers = make(map[string]mqHttpSdk.MQProducer)
	consumers = make(map[string]mqHttpSdk.MQConsumer)
}

func InitProducer(config MqTopicConfig) {
	if _, ok := producers[config.Name]; !ok {
		producers[config.Name] = client.GetProducer(config.InstanceId, config.Topic)
	}
}

func InitConsumer(config MqTopicConfig) {
	if _, ok := consumers[config.Name]; !ok {
		consumers[config.Name] = client.GetConsumer(config.InstanceId, config.Topic, config.Group, "")
	}
}

func GetProducer(name string) mqHttpSdk.MQProducer {
	if cli, ok := producers[name]; ok {
		return cli
	} else {
		return nil
	}
}

func GetConsumer(name string) mqHttpSdk.MQConsumer {
	if cli, ok := consumers[name]; ok {
		return cli
	} else {
		return nil
	}
}
