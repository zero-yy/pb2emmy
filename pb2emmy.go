package pb2emmy

import (
	"fmt"
	"github.com/emicklei/proto"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type pb2emmy struct {
	*Config
	proto.NoopVisitor

	fieldDes []string
	output   string

	enumMap map[string]interface{}
}

func NewPb2Emmy(c *Config) *pb2emmy {
	return &pb2emmy{
		Config: c,

		enumMap: make(map[string]interface{}),
	}
}

func (m *pb2emmy) Do() {
	m.output = "---auto gen by pb2emmy\n"

	def, err := m.parseProto()
	if err != nil {
		panic(err)
	}

	proto.Walk(def,
		proto.WithEnum(func(enum *proto.Enum) {
			m.enumMap[enum.Name] = nil
		}))

	proto.Walk(def,
		proto.WithMessage(m.handleMessage))

	err = os.WriteFile(m.Config.OutputName, []byte(m.output), 0644)
	if err != nil {
		panic(err)
	}
}

func (m *pb2emmy) handleMessage(msg *proto.Message) {
	//fmt.Println(msg.Name)

	m.fieldDes = nil
	for _, each := range msg.Elements {
		each.Accept(m)
	}

	fields := strings.Join(m.fieldDes, ",")
	alias := fmt.Sprintf("---@alias %s {%s}", "Pb"+msg.Name, fields)
	fmt.Println(alias)

	m.output += alias
	m.output += "\n"
}

func (m *pb2emmy) parseProto() (definition *proto.Proto, err error) {
	paths, err := m.readDir()
	if err != nil {
		panic(err)
	}

	var readers []io.Reader
	for _, v := range paths {
		reader, _ := os.Open(v)
		defer reader.Close()
		readers = append(readers, reader)
	}
	allReader := io.MultiReader(readers...)
	parser := proto.NewParser(allReader)
	return parser.Parse()
}

func (m *pb2emmy) readDir() (paths []string, err error) {
	err = filepath.Walk(m.Config.InputDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		// 忽略目录
		if info.IsDir() {
			return nil
		}

		if strings.Contains(strings.ToLower(path), ".proto") {
			//fmt.Println(path)
			paths = append(paths, path)
		}
		return nil
	})
	return
}

func (m *pb2emmy) pbType2Lua(t string) (luaType string) {
	if v, ok := __pbType2Lua[t]; ok {
		return v
	}

	if _, ok := m.enumMap[t]; ok {
		return "number"
	}

	return t
}

//////////////////////////////////
func (m *pb2emmy) VisitOption(o *proto.Option) {
	//fmt.Println(o.Name)
}

func (m *pb2emmy) VisitNormalField(i *proto.NormalField) {
	field := fmt.Sprintf("%s:%s", i.Name, m.pbType2Lua(i.Type))
	if i.Repeated {
		field += "[]"
	}
	m.fieldDes = append(m.fieldDes, field)
}
