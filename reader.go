package hl7

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type Reader struct {
	reader     *bufio.Reader
	delimiters *Delimiters
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		reader: bufio.NewReader(r),
	}
}

func (r *Reader) ReadSegment() (seg Segment, err error) {
	data, err := r.reader.ReadBytes(CR)
	if len(data) == 0 {
		return nil, io.EOF
	}

	var offset, lastPos, pos int
	var field Field
	var fieldItem FieldItem
	var component Component

	for unicode.IsSpace(rune(data[offset])) {
		offset++
	}

	if offset == len(data) {
		return nil, io.EOF
	}

	if r.delimiters == nil {
		d, err := getDelimiters(data)
		if err != nil {
			return nil, err
		}
		r.delimiters = &d
		seg = Segment{
			Field{FieldItem{Component{SubComponent("MSH")}}},
			Field{FieldItem{Component{SubComponent(data[offset+3 : offset+4])}}},
			Field{FieldItem{Component{SubComponent(data[offset+4 : offset+8])}}},
		}
		if len(data) == 8 {
			return seg, nil
		}
		offset += 9
		if len(data) > 8 && data[8] != r.delimiters.Field {
			return nil, ErrInvalidHeader(fmt.Errorf("invalid character found after header content; expected \\x%02x but got \\x%02x", d.Field, data[8]))
		}
	}

	commitBuffer := func(force bool) {
		if lastPos != pos || force {
			component = append(component, SubComponent(data[lastPos:pos]))
			lastPos = pos
		}
	}

	commitComponent := func(force bool) {
		commitBuffer(false)
		if component != nil || force {
			fieldItem = append(fieldItem, component)
			component = nil
		}
	}

	commitFieldItem := func(force bool) {
		commitComponent(false)
		if fieldItem != nil || force {
			field = append(field, fieldItem)
			fieldItem = nil
		}
	}

	commitField := func(force bool) {
		commitFieldItem(false)
		if field != nil || force {
			seg = append(seg, field)
			field = nil
		}
	}

	for lastPos, pos = offset, offset; pos < len(data); pos++ {
		switch data[pos] {
		case r.delimiters.Field:
			commitField(true)
			lastPos++
		case r.delimiters.Repeat:
			commitFieldItem(true)
			lastPos++
		case r.delimiters.Component:
			commitComponent(true)
			lastPos++
		case r.delimiters.SubComponent:
			commitBuffer(true)
			lastPos++
		default:
		}
	}
	commitField(false)

	if err != nil {
		return seg, err
	}

	return
}
