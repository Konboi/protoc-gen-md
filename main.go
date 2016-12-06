package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Konboi/protoc-gen-md/generator"
	"github.com/Konboi/protoc-gen-md/parser"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
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
	//options := make(map[string]string, 0)
	// todo parse parameter
	// support template path
	//log.Println(*req.Parameter)

	p := parser.New()
	if err := p.Load(req); err != nil {
		log.Fatalf("[error] parser load config: %#v", err)
	}

	var outs []*plugin.CodeGeneratorResponse_File

	g := generator.New()

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
