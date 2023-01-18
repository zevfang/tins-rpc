package common

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/jhump/protoreflect/desc"

	pp "github.com/emicklei/proto"
)

type Definitions struct {
	fd                *desc.FileDescriptor
	fileNamesRead     []string
	fileBody          string
	fileNameToPackage map[string]string
	fileName          string                 // ct.proto
	pkgName           string                 // ct
	svcs              map[string]*pp.Service // Ct
	rpcs              map[string]*pp.RPC     // GetTable
	msgs              map[string]*pp.Message // GetTableReq
}

func NewDefinitions() *Definitions {
	return &Definitions{
		fileNamesRead:     make([]string, 0),
		fileBody:          "",
		fileNameToPackage: make(map[string]string),
		fileName:          "",
		pkgName:           "",
		svcs:              make(map[string]*pp.Service),
		rpcs:              make(map[string]*pp.RPC),
		msgs:              make(map[string]*pp.Message),
	}
}

func (d *Definitions) ReadFile(filepath string) error {
	fileReader, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fileReader.Close()
	return d.ReadFrom(filepath, fileReader)
}

func (d *Definitions) ReadFrom(filepath string, reader io.Reader) error {
	for _, each := range d.fileNamesRead {
		if each == filepath {
			return nil
		}
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	d.fd = GetProtoFileDescriptor(filepath)
	d.fileNamesRead = append(d.fileNamesRead, filepath)
	parser := pp.NewParser(bytes.NewReader(data))
	def, err := parser.Parse()
	if err != nil {
		return err
	}
	pkg := packageOf(def)
	d.fileNameToPackage[filepath] = pkg
	d.fileBody = string(data)
	d.fileName = path.Base(filepath)
	d.pkgName = pkg
	// parse service
	pp.Walk(def, pp.WithService(func(each *pp.Service) {
		d.AddService(each.Name, each)
	}))
	pp.Walk(def, pp.WithRPC(func(each *pp.RPC) {
		d.AddRpc(each.Name, each)
	}))
	pp.Walk(def, pp.WithMessage(func(each *pp.Message) {
		d.AddMessage(each.Name, each)
	}))
	return nil
}

func (d *Definitions) GetFd() *desc.FileDescriptor {
	return d.fd
}

func (d *Definitions) GetFileName() string {
	return d.fileName
}

func (d *Definitions) GetFileBody() string {
	return d.fileBody
}

func (d *Definitions) GetPkgName() string {
	return d.pkgName
}

func (d *Definitions) MessagesInPackage(pkg string) (list []*pp.Message) {
	for k, v := range d.msgs {
		if strings.HasPrefix(k, pkg+".") {
			list = append(list, v)
		}
	}
	return
}

func (d *Definitions) AddService(name string, svc *pp.Service) {
	key := fmt.Sprintf("%s", name)
	d.svcs[key] = svc
}

func (d *Definitions) GetServices() map[string]*pp.Service {
	return d.svcs
}

func (d *Definitions) AddRpc(name string, rpc *pp.RPC) {
	key := fmt.Sprintf("%s", name)
	d.rpcs[key] = rpc
}

func (d *Definitions) Rpc(name string) (m *pp.RPC, ok bool) {
	key := fmt.Sprintf("%s", name)
	m, ok = d.rpcs[key]
	return
}

func (d *Definitions) AddMessage(name string, message *pp.Message) {
	key := fmt.Sprintf("%s", name)
	d.msgs[key] = message
}

func (d *Definitions) Message(name string) (m *pp.Message, ok bool) {
	key := fmt.Sprintf("%s", name)
	m, ok = d.msgs[key]
	return
}

func packageOf(def *pp.Proto) string {
	for _, each := range def.Elements {
		if p, ok := each.(*pp.Package); ok {
			return p.Name
		}
	}
	return ""
}
