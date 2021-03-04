package common

import (
	"encoding/base64"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/linking-lib/go-game-lib/protos"
	"github.com/linking-lib/go-game-lib/utils/serialize"
	lkJson "github.com/linking-lib/go-game-lib/utils/serialize/json"
	"github.com/linking-lib/go-game-lib/utils/strs"
	"strings"
)

type ConvertUtils struct {
	serializer serialize.Serializer
}

var (
	convert = &ConvertUtils{
		serializer: lkJson.NewSerializer(),
	}
)

// SetSerializer customize application serializer, which automatically Marshal
// and UnMarshal handler payload
func SetSerializer(ser serialize.Serializer) {
	convert.serializer = ser
}

// GetSerializer gets the app serializer
func GetSerializer() serialize.Serializer {
	return convert.serializer
}

func ConvertJson(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		panic(`ConvertJson: ` + err.Error())
	}
	return string(jsonStr)
}

func ParseJson(str string, data interface{}) interface{} {
	err := json.Unmarshal([]byte(str), data)
	if err != nil {
		panic(`ParseJson: strs(` + str + `): ` + err.Error())
	}
	return data
}

func ConvertRequest(param string, m proto.Message) {
	serializerName := convert.serializer.GetName()
	var b []byte
	var err error
	switch serializerName {
	case "protos":
		b, err = base64.StdEncoding.DecodeString(param)
	case "json":
		b = []byte(param)
	default:
		panic(`ConvertRequest type error: type(` + serializerName + `): `)
	}
	if err != nil {
		panic(`ConvertRequest Decode: strs(` + param + `): ` + err.Error())
	}
	err = convert.serializer.Unmarshal(b, m)
	if err != nil {
		panic(`ConvertRequest Unmarshal error: strs(` + param + `): ` + err.Error())
	}
}

func ConvertResult(result *protos.LResult) string {
	serializerName := convert.serializer.GetName()
	var b []byte
	var err error
	var data string
	switch serializerName {
	case "protos":
		b, err = convert.serializer.Marshal(result)
		data = base64.StdEncoding.EncodeToString(b)
	case "json":
		b, err = convert.serializer.Marshal(convertJsonResult(result))
		data = string(b)
	default:
		panic(`ConvertRequest type error: type(` + serializerName + `): `)
	}
	if err != nil {
		panic(`ConvertResult Marshal error: ` + err.Error())
	}
	return data
}

func convertJsonResult(result *protos.LResult) LResult {
	var sResult LResult
	sResult.Api = result.Api
	sResult.Code = result.Code
	sResult.ErrCode = result.ErrCode
	sResult.Msg = result.Msg
	var data = result.GetData()
	var obj interface{}
	if !strs.IsEmpty(data) {
		var b = []byte(data)
		if json.Valid(b) {
			if err := json.Unmarshal(b, &obj); err == nil {
				sResult.Data = obj
			} else {
				panic(`convertJsonResult string 2 json error: ` + err.Error())
			}
		}
	}
	return sResult
}

func ConvertApi(api string) string {
	a := strings.Split(api, ".")
	num := len(a)
	if num >= 3 {
		return a[len(a)-3] + "." + a[len(a)-2] + "." + a[len(a)-1]
	} else if num >= 2 {
		return a[len(a)-2] + "." + a[len(a)-1]
	} else {
		panic("ConvertApi api len error" + api)
	}
}
