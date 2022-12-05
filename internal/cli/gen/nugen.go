package gen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

// Kinds
const (
	UnknownGoKind = iota
	GoMessageKind
	GoSegmentKind
	GoDataTypeKind
	GoGroupKind
)

// GoType ...
type GoType struct {
	Name        string
	Title       string
	Description string
	Kind        int
	Fields      []*GoField
}

// Comment ...
func (t GoType) Comment() string {
	var builder strings.Builder
	builder.WriteString(t.Name)
	if t.Title != "" {
		builder.WriteString(" - ")
		builder.WriteString(t.Title)
	}

	if t.Description != "" {
		builder.WriteByte('\n')
		builder.WriteString(t.Description)
	}

	if t.Title == "" && t.Description == "" {
		builder.WriteString(" ...")
	}

	return builder.String()
}

func (t GoType) ConcreteName() string {
	return t.Name
}

// GoField ...
type GoField struct {
	Name     string
	TypeName string
	Type     *GoType
	Pos      string
	Len      string
	Use      string
	Rep      string
	IsArray  bool
	IsGroup  bool
}

func (f GoField) ConcreteTypeName() string {
	return f.TypeName
}

// IsRequired ...
func (f *GoField) IsRequired() bool {
	return f.Use == "R"
}

func isRptArray(rpt string) (result bool) {
	result = rpt == "*"
	if rpt, err := strconv.ParseInt(rpt, 10, 32); err == nil {
		result = rpt > 1
	}
	return
}

func genGoFieldFromFieldInfo(f *APIFieldInfo) (field *GoField) {
	field = &GoField{}
	field.Name = goName(f.Name)
	field.TypeName = goName(f.DataType)
	field.IsArray = isRptArray(f.Rpt)
	field.Len = fmt.Sprint(f.Length)

	pos := f.Position
	parts := strings.SplitN(pos, ".", 2)
	if len(parts) > 1 {
		pos = parts[1]
	}

	field.Pos = pos
	field.Rep = f.Rpt
	field.Use = f.Usage
	return
}

func genGoFieldFromSegInfo(s *APISegmentInfo) (field *GoField) {
	field = &GoField{}
	longName := strcase.ToCamel(s.LongName)
	name := s.Name

	if longName == "" {
		parts := strings.Split(s.Name, "_")
		for i, p := range parts {
			parts[i] = strcase.ToCamel(strings.ToLower(p))
		}
		longName = strings.Join(parts, " ")
		name = strcase.ToCamel(longName)
	}

	field.Name = goName(longName)
	field.TypeName = goName(name)
	field.IsArray = isRptArray(s.Rpt)
	field.Len = ""
	field.Pos = s.Sequence
	field.Rep = s.Rpt
	field.Use = s.Usage
	field.IsGroup = s.IsGroup
	return
}

var fixupOrder []string = []string{"Id", "IDent", "Cpu", "Oid", "Uid"}
var fixups map[string]string = map[string]string{
	"Id":    "ID",
	"IDent": "Ident",
	"Cpu":   "CPU",
	"Oid":   "OID",
	"Uid":   "UID",
}

func goName(nm string) (name string) {
	name = strcase.ToCamel(nm)
	for _, n := range fixupOrder {
		name = strings.ReplaceAll(name, n, fixups[n])
	}
	if strings.HasPrefix(strings.ToLower(name), "setid") {
		name = "SetID"
	}
	return
}
