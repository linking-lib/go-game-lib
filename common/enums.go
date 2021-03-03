/**
 * @Title  公用包
 * @Description 公用的属性和方法
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:07
 */
package common

type ProtocolType string

const (
	ProtocolProtobuf ProtocolType = "proto"
	ProtocolJson     ProtocolType = "json"
)

func (p ProtocolType) ValueOf(value string) ProtocolType {
	switch value {
	case "proto":
		return ProtocolProtobuf
	case "json":
		return ProtocolJson
	default:
		return "none"
	}
}

func (p ProtocolType) String() string {
	switch p {
	case ProtocolProtobuf:
		return "proto"
	case ProtocolJson:
		return "json"
	default:
		return "none"
	}
}
