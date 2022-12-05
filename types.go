package hl7

import (
	"strings"
)

type HL7Type interface {
	String() string
	Len() int
	Index(int) HL7Type
}

type Group []HL7Type

func (g Group) String() string {
	var builder strings.Builder
	for i := range g {
		s := g[i].String()
		if s == "" {
			continue
		}
		builder.WriteString(s)
		if i != len(g)-1 {
			builder.WriteByte(CR)
		}
	}
	return builder.String()
}

func (g Group) Len() int { return len(g) }
func (g Group) Index(i int) HL7Type {
	if i < len(g) && g[i] != nil {
		return g[i]
	}
	return Group{}
}

type SegmentList []Segment

func (sl SegmentList) String() string {
	var builder strings.Builder
	for i := range sl {
		s := sl[i].String()
		if s == "" {
			continue
		}
		builder.WriteString(s)
		if i != len(sl)-1 {
			builder.WriteByte(CR)
		}
	}
	return builder.String()
}

func (sl SegmentList) Len() int { return len(sl) }
func (sl SegmentList) Index(i int) HL7Type {
	if i < len(sl) {
		return sl[i]
	}
	return Segment{}
}

type GroupList []Group

func (gl GroupList) String() string {
	var builder strings.Builder
	for i := range gl {
		s := gl[i].String()
		if s == "" {
			continue
		}
		builder.WriteString(s)
		if i != len(gl)-1 {
			builder.WriteByte(CR)
		}
	}
	return builder.String()
}
func (gl GroupList) Len() int { return len(gl) }
func (gl GroupList) Index(i int) HL7Type {
	if i < len(gl) {
		return gl[i]
	}
	return Group{}
}

type Message []Segment

func (m Message) String() string {
	var builder strings.Builder
	for i := range m {
		s := m[i].String()
		if s == "" {
			continue
		}
		builder.WriteString(s)
		if i != len(m)-1 {
			builder.WriteByte(CR)
		}
	}
	return builder.String()
}

func (m Message) Segment(name string) Segment {
	for i := range m {
		if m[i].Name() == name {
			return m[i]
		}
	}
	return nil
}

func (m Message) Segments(name string) (segs []Segment) {
	for i := range m {
		if m[i].Name() == name {
			segs = append(segs, m[i])
		}
	}
	return
}

func (m Message) Index(i int) (typ HL7Type) {
	if i < len(m) {
		typ = m[i]
	}
	return
}

func (m Message) Len() int { return len(m) }

func (m Message) Delimters() Delimiters {
	if len(m) > 0 && len(m[0]) > 2 {
		return Delimiters{
			m[0][1][0][0][0][0],
			m[0][2][0][0][0][0],
			m[0][2][0][0][0][1],
			m[0][2][0][0][0][2],
			m[0][2][0][0][0][3],
		}
	}
	return Delimiters{}
}

type Delimiters struct {
	Field, Component, Repeat, Escape, SubComponent byte
}

var DefaultDelimiters = Delimiters{'|', '^', '~', '\\', '&'}

type Segment []Field

func (s Segment) Len() int            { return len(s) }
func (s Segment) Index(i int) HL7Type { return s.Field(i) }

func (s Segment) Name() string {
	if len(s) > 0 && len(s[0]) > 0 {
		return string(s[0][0][0][0])
	}
	return ""
}

func (s Segment) Field(i int) Field {
	if i < len(s) {
		return s[i]
	}
	return Field{}
}

func (s Segment) String() (res string) {
	var builder strings.Builder
	start := 0
	if s.Name() == "MSH" {
		builder.WriteString("MSH")
		builder.WriteByte(DefaultDelimiters.Field)
		builder.WriteByte(DefaultDelimiters.Component)
		builder.WriteByte(DefaultDelimiters.Repeat)
		builder.WriteByte(DefaultDelimiters.Escape)
		builder.WriteByte(DefaultDelimiters.SubComponent)
		builder.WriteByte(DefaultDelimiters.Field)
		start = 3
	}
	for i := start; i < len(s); i++ {
		builder.WriteString(s[i].String())
		if i != len(s)-1 {
			builder.WriteByte(DefaultDelimiters.Field)
		}
	}
	return builder.String()
}

type Field []FieldItem

func (f Field) Len() int            { return len(f) }
func (f Field) Index(i int) HL7Type { return f.Rep(i) }

func (f Field) Rep(i int) FieldItem {
	if i < len(f) {
		return f[i]
	}
	return FieldItem{}
}

func (f Field) String() string {
	values := make([]string, len(f))
	for i := range f {
		values[i] = f[i].String()
	}
	return strings.Join(values, string(DefaultDelimiters.Repeat))
}

func (f Field) Component(i int) Component {
	return f.Rep(0).Component(i)
}

type FieldItem []Component

func (fi FieldItem) Len() int            { return len(fi) }
func (fi FieldItem) Index(i int) HL7Type { return fi.Component(i) }

func (fi FieldItem) Component(i int) Component {
	if i < len(fi) {
		return fi[i]
	}
	return Component{}
}

func (fi FieldItem) String() string {
	values := make([]string, len(fi))
	for i := range fi {
		values[i] = fi[i].String()
	}
	return strings.Join(values, string(DefaultDelimiters.Component))
}

type Component []SubComponent

func (c Component) Len() int            { return len(c) }
func (c Component) Index(i int) HL7Type { return c.SubComponent(i) }
func (c Component) SubComponent(i int) SubComponent {
	if i < len(c) {
		return c[i]
	}
	return SubComponent{}
}

func (c Component) String() string {
	values := make([]string, len(c))
	for i := range c {
		values[i] = c[i].String()
	}
	return strings.Join(values, string(DefaultDelimiters.SubComponent))
}

type SubComponent []byte

func (sc SubComponent) Len() int            { return 0 }
func (sc SubComponent) Index(i int) HL7Type { return nil }
func (sc SubComponent) String() string      { return string(sc) }
