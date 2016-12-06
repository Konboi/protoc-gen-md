package generator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Konboi/protoc-gen-md/parser"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pkg/errors"
)

type Generator struct {
	tmpl   *template.Template
	option Option
}

type Option struct {
	Template string
	FileName string
}

func New(opt Option) *Generator {
	tmpl := template.Must(template.ParseFiles(opt.Template))

	return &Generator{
		tmpl:   tmpl,
		option: opt,
	}
}

func (g *Generator) Generate(prts []*parser.Proto) (*plugin.CodeGeneratorResponse_File, error) {
	var doc bytes.Buffer

	err := g.tmpl.Execute(&doc, prts)
	if err != nil {
		return nil, errors.Wrap(err, "[error] template execute error")
	}
	fileName := fmt.Sprintf("%s.md", g.option.FileName)
	content := doc.String()

	return &plugin.CodeGeneratorResponse_File{
		Name:    &fileName,
		Content: &content,
	}, nil
}
