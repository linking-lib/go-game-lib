package linking

import (
	lkCommon "github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/utils/serialize/json"
	"github.com/linking-lib/go-game-lib/utils/serialize/protobuf"
)

// Serializer implements the serialize.Serializer interface
type SelfSerializer struct {
	protobuf *protobuf.Serializer
	json     *json.Serializer
}

var ProtocolType string

// NewSerializer returns a new Serializer.
func NewSerializer(protocolType string) *SelfSerializer {
	ProtocolType = protocolType
	return &SelfSerializer{protobuf: protobuf.NewSerializer(), json: json.NewSerializer()}
}

// Marshal returns the JSON encoding of v.
func (s *SelfSerializer) Marshal(v interface{}) ([]byte, error) {
	switch ProtocolType {
	case lkCommon.ProtocolProtobuf.String():
		return s.protobuf.Marshal(v)
	case lkCommon.ProtocolJson.String():
		return s.json.Marshal(v)
	default:
		return s.json.Marshal(v)
	}
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func (s *SelfSerializer) Unmarshal(data []byte, v interface{}) error {
	switch ProtocolType {
	case lkCommon.ProtocolProtobuf.String():
		return s.protobuf.Unmarshal(data, v)
	case lkCommon.ProtocolJson.String():
		return s.json.Unmarshal(data, v)
	default:
		return s.json.Unmarshal(data, v)
	}
}

// GetName returns the name of the serializer.
func (s *SelfSerializer) GetName() string {
	switch ProtocolType {
	case lkCommon.ProtocolProtobuf.String():
		return "protos"
	case lkCommon.ProtocolJson.String():
		return "json"
	default:
		return "json"
	}
}
