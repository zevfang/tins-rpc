package common

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"os"
	"strings"

	parse "github.com/emicklei/proto"
)

type ProtoData struct {
	PackageName string            //package
	ServiceName string            //service
	Relation    map[string]string //req:rpc
	RequestList map[string]string //rpc：reqJson
}

// GerProtoData 解析并获取proto
func GerProtoData(path string) (ProtoData, error) {
	protoData := ProtoData{
		PackageName: "",
		ServiceName: "",
		Relation:    make(map[string]string),
		RequestList: make(map[string]string),
	}
	reader, err := os.Open(path)
	if err != nil {
		return protoData, err
	}
	defer reader.Close()
	parser := parse.NewParser(reader)
	definition, _ := parser.Parse()

	parse.Walk(definition,
		parse.WithPackage(protoData.handlePackage),
		parse.WithService(protoData.handleService),
		parse.WithRPC(protoData.handleRPC),
	)
	parse.Walk(definition,
		parse.WithMessage(protoData.handleMessage),
	)
	return protoData, nil
}
func (p *ProtoData) handlePackage(s *parse.Package) {
	p.PackageName = s.Name
}
func (p *ProtoData) handleService(s *parse.Service) {
	p.ServiceName = s.Name
}
func (p *ProtoData) handleRPC(s *parse.RPC) {
	p.Relation[s.RequestType] = s.Name
	p.RequestList[s.Name] = ""
}

func (p *ProtoData) handleMessage(m *parse.Message) {
	list := make(map[string]interface{})
	for _, each := range m.Elements {
		if f, ok := each.(*parse.NormalField); ok {
			var ty interface{}
			switch {
			case f.Type == "int32" || f.Type == "uint32" || f.Type == "int64" || f.Type == "uint64":
				if strings.Contains(f.Name, "id") {
					ty = 100
				} else {
					ty = 10
				}
			case f.Type == "string":
				if strings.Contains(f.Name, "id") {
					uid, _ := uuid.NewUUID()
					ty = uid
				} else {
					ty = "Hello"
				}
			case f.Type == "repeated" || f.Repeated:
				ty = []interface{}{}
			default:
				ty = nil
			}
			list[f.Name] = ty
		}
	}

	rpcName := p.Relation[m.Name]
	if len(rpcName) == 0 {
		return
	}

	byteList, _ := json.Marshal(&list)
	// 转码格式化
	p.RequestList[rpcName] = ParseJSON(byteList)
	return
}

//
//type decoder struct {
//	r map[string]interface{}
//}
//
//func (d *decoder) decodeTag(elements []parse.Visitee, wire uint64) error {
//	for _, each := range elements {
//		if f, ok := each.(*parse.NormalField); ok {
//			d.decodeNormalField(f)
//		}
//		if f, ok := each.(*parse.MapField); ok {
//			d.decodeMapField(f, wire)
//		}
//	}
//	return nil
//}
//
//func (d *decoder) decodeNormalField(f *parse.NormalField) {
//	if "string" == f.Type {
//		d.add(f.Name, "Test", f.Repeated, false)
//	}
//	if "int64" == f.Type || "uint64" == f.Type || "int32" == f.Type || "uint32" == f.Type || "float" == f.Type {
//		d.add(f.Name, 123, f.Repeated, false)
//	}
//	if "bool" == f.Type {
//		d.add(f.Name, true, f.Repeated, false)
//	}
//	if _, ok := d.d.Message(d.p, f.Type); ok {
//		d.decodeNormalFieldMessage(f)
//	}
//}

// ParseJSON 格式化为json
func ParseJSON(b []byte) string {
	// 转码格式化
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "\t")
	return out.String()
}
