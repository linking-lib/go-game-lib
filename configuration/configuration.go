package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	MysqlName      string
	MysqlUrl       string
	MysqlUser      string
	MysqlPwd       string
	MysqlDb        string
	MysqlMaxIdle   int
	MysqlMaxActive int
	RedisName      string
	RedisUrl       string
	RedisPwd       string
	RedisMaxIdle   int
	RedisMaxActive int
	RedisTimeout   int
	RedisDb        int
	MqEndpoint     string
	MqAccessKey    string
	MqSecretKey    string
	MqGroup        string
	MqInstanceId   string
	MqName         string
	MqTopic        string
	EtcdEndpoints  string
	NatsConnect    string
}

func LoadConfig(path string) (*Configuration, error) {
	file, _ := os.Open(path)
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	conf := &Configuration{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
