package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pkg/errors"
)

const (
	DefineServiceLimit = 1
)

var (
	ignoreFiles = []string{"google/"}
)

type Parser struct {
	files []*Proto
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Files() []*Proto {
	return p.files
}

func (p *Parser) Load(req plugin.CodeGeneratorRequest) error {
	for _, file := range req.GetProtoFile() {
		err := p.load(file)
		if err != nil {
			return errors.Wrapf(err, "loadFile %s error", file.GetName())
		}
	}

	return nil
}

func (p *Parser) load(file *descriptor.FileDescriptorProto) error {
	prt := &Proto{
		PackageName: file.GetPackage(),
	}

	prt.Messages = make(map[string]Message, 0)
	for _, message := range file.GetMessageType() {
		prt.LoadMessage(message)
	}

	if len(file.GetService()) == DefineServiceLimit {
		service := file.GetService()[0]
		prt.LoadService(service)
	}

	p.files = append(p.files, prt)

	return nil
}
