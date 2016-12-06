package parser

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/serenize/snaker"
)

type Proto struct {
	PackageName string
	Service     Service
	Messages    map[string]Message
}

type Service struct {
	Name    string
	Methods []Method
}

type Method struct {
	Name     string
	Request  Message
	Response Message
}

type Message struct {
	Name   string
	Fields []Field
}

type Field struct {
	Type  string
	Name  string
	Label int
}

func (p Proto) Message(key string) Message {
	return p.Messages[key]
}

func (prt *Proto) LoadService(service *descriptor.ServiceDescriptorProto) {

	s := Service{
		Name: service.GetName(),
	}

	s.Methods = make([]Method, 0, len(service.GetMethod()))
	for _, method := range service.GetMethod() {
		m := Method{
			Name:     method.GetName(),
			Request:  prt.Messages[method.GetInputType()],
			Response: prt.Messages[method.GetOutputType()],
		}
		s.Methods = append(s.Methods, m)
	}

	prt.Service = s
}

func (prt *Proto) LoadMessage(message *descriptor.DescriptorProto) {
	msg := Message{}
	msg.Name = message.GetName()
	msg.Fields = make([]Field, 0, len(message.GetField()))
	for _, field := range message.GetField() {
		typeName := typeToString(field.GetType())
		if typeName == "message" {
			typeName = strings.Replace(field.GetTypeName(), ".", "", 1)
		}

		msg.Fields = append(msg.Fields, Field{
			Type:  typeName,
			Name:  field.GetName(),
			Label: int(field.GetLabel()),
		})

	}
	prt.Messages[prt.messageTypeName(msg)] = msg
}

func (prt *Proto) messageTypeName(msg Message) string {
	return fmt.Sprintf(".%s.%s", prt.PackageName, msg.Name)
}

func (s Service) PathStr() string {
	return fmt.Sprintf("%s", snaker.CamelToSnake(s.Name))
}

func (m Method) PathStr() string {
	return fmt.Sprintf("%s", snaker.CamelToSnake(m.Name))
}

func (m Method) RequestMethod() string {
	if len(m.Request.Fields) == 0 {
		return "GET"
	}

	return "POST"
}

func typeToString(t descriptor.FieldDescriptorProto_Type) string {
	switch t {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "double"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return "float"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "bytes"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		return "enum"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		return "message"
	}

	return ""
}
