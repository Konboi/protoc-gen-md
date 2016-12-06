package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Konboi/protoc-gen-md/generator"
	"github.com/Konboi/protoc-gen-md/parser"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	DefaultGeneratorType = "service"
)

func main() {
	log.SetFlags(log.Lshortfile)

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("[error] read os.Stding: %v", err)
	}

	req := plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(data, &req)
	if err != nil {
		log.Fatalf("[error] load data proto.Unmarshal(): %v", err)
	}

	opt := parserParam(req.GetParameter())

	p := parser.New()
	if err := p.Load(req); err != nil {
		log.Fatalf("[error] parser load config: %#v", err)
	}

	var outs []*plugin.CodeGeneratorResponse_File

	g := generator.New(opt)
	out, err := g.Generate(p.Files())
	if err != nil {
		log.Fatalf("[error] generator write: %#v", err)

	}
	outs = append(outs, out)

	resp := plugin.CodeGeneratorResponse{
		File: outs,
	}

	buf, err := proto.Marshal(&resp)
	if err != nil {
		log.Fatalf("[error] marsahl response: %#v", err)
	}

	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatalf("[error] write generate code: %#v", err)
	}
}

func parserParam(param string) generator.Option {
	val := strings.Split(param, ",")

	params := make(map[string]string, 0)
	for _, p := range val {
		if !strings.Contains(p, "=") {
			continue
		}

		keyVal := strings.Split(p, "=")
		params[keyVal[0]] = keyVal[1]
	}

	if _, ok := params["template"]; !ok {
		log.Fatalf("Please set template option")
	}

	if _, ok := params["name"]; !ok {
		log.Fatalf("Please set generate file name")
	}

	opt := generator.Option{
		Template: params["template"],
		FileName: params["name"],
	}

	return opt
}
