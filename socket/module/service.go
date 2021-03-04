package module

import (
	"errors"
	"fmt"
	"github.com/topfreegames/pitaya/conn/message"
	"reflect"
)

var (
	handlers = make(map[string]*Handler, 0)
)

type (
	Service struct {
		Name     string              // name of service
		Type     reflect.Type        // type of the receiver
		Receiver reflect.Value       // receiver of methods for the service
		Handlers map[string]*Handler // registered methods
		Options  options             // options
	}

	Handler struct {
		Receiver    reflect.Value  // receiver of method
		Method      reflect.Method // method stub
		Type        reflect.Type   // low-level type of method
		IsRawArg    bool           // whether the data need to serialize
		MessageType message.Type   // handler allowed message type (either request or notify)
	}
)

func ContainsHandler(name string) bool {
	_, ok := handlers[name]
	return ok
}

func NewService(comp SelfComponent, opts []Option) *Service {
	s := &Service{
		Type:     reflect.TypeOf(comp),
		Receiver: reflect.ValueOf(comp),
	}
	// apply options
	for i := range opts {
		opt := opts[i]
		opt(&s.Options)
	}
	if name := s.Options.name; name != "" {
		s.Name = name
	} else {
		s.Name = reflect.Indirect(s.Receiver).Type().Name()
	}
	return s
}

func Register(comp SelfComponent, options ...Option) {
	s := NewService(comp, options)
	if err := s.ExtractHandler(); err != nil {
		return
	}
	// register all handlers
	for name, handler := range s.Handlers {
		handlers[fmt.Sprintf("%s.%s", s.Name, name)] = handler
	}
}

func (s *Service) ExtractHandler() error {
	s.Handlers = suitableHandlerMethods(s.Type, s.Options.nameFunc)
	if len(s.Handlers) == 0 {
		str := ""
		// To help the user, see if a pointer receiver would work.
		method := suitableHandlerMethods(reflect.PtrTo(s.Type), s.Options.nameFunc)
		if len(method) != 0 {
			str = "type " + s.Name + " has no exported methods of handler type (hint: pass a pointer to value of that type)"
		} else {
			str = "type " + s.Name + " has no exported methods of handler type"
		}
		return errors.New(str)
	}
	for i := range s.Handlers {
		s.Handlers[i].Receiver = s.Receiver
	}
	return nil
}

func suitableHandlerMethods(typ reflect.Type, nameFunc func(string) string) map[string]*Handler {
	methods := make(map[string]*Handler)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mn := method.Name
		// rewrite handler name
		if nameFunc != nil {
			mn = nameFunc(mn)
		}
		if isHandlerMethod(method) {
			handler := &Handler{
				Method: method,
			}
			methods[mn] = handler
		}
	}
	return methods
}

func isHandlerMethod(method reflect.Method) bool {
	mt := method.Type
	// Method must be exported.
	if method.PkgPath != "" {
		return false
	}

	// Method needs two or three ins: receiver, context.Context and optional []byte or pointer.
	if mt.NumIn() != 2 && mt.NumIn() != 3 {
		return false
	}

	return true
}
