package generator

import (
	"bytes"
	"text/template"

	"github.com/Konboi/protoc-gen-md/parser"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pkg/errors"
	"os"
)

type Generator struct {
	tmpl *template.Template
}

func New() *Generator {
	// todo
	tmpl := template.Must(template.ParseFiles("./generator/markdown.template"))

	return &Generator{
		tmpl: tmpl,
	}
}

func (g Generator) Generate(prts []*parser.Proto) (*plugin.CodeGeneratorResponse_File, error) {
	var doc bytes.Buffer
	g.tmpl.Execute(os.Stderr, prts)

	err := g.tmpl.Execute(&doc, prts)
	if err != nil {
		return nil, errors.Wrap(err, "[error] template execute error")
	}
	fileName := "doc.md"
	content := doc.String()

	return &plugin.CodeGeneratorResponse_File{
		Name:    &fileName,
		Content: &content,
	}, nil
}
