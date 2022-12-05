package hl7

import (
	"bytes"
	"errors"
	"fmt"
)

type (
	// ErrTooShort is returned if a message isn't long enough to contain a valid header
	ErrTooShort error
	// ErrInvalidHeader is returned if a message doesn't start with "MSH", or
	// the header isn't exactly the correct length, or any of the control
	// characters aren't unique
	ErrInvalidHeader error
)

// Constants describing possible message boundaries.
const (
	CR = '\r'
	LF = '\n'
	FF = '\f'
	NB = '\x00'
)

func getDelimiters(buf []byte) (Delimiters, error) {
	if len(buf) < 8 {
		return Delimiters{}, ErrTooShort(fmt.Errorf("message must be at least eight bytes long; instead was %d", len(buf)))
	}

	if !bytes.HasPrefix(buf, []byte("MSH")) {
		return Delimiters{}, ErrInvalidHeader(fmt.Errorf("expected message to begin with MSH; instead found %q", buf[0:3]))
	}

	fs := buf[3]
	cs := buf[4]
	rs := buf[5]
	ec := buf[6]
	ss := buf[7]

	if fs == cs || fs == rs || fs == ec || fs == ss || cs == rs || cs == ec || cs == ss || rs == ec || rs == ss || ec == ss {
		return Delimiters{}, ErrInvalidHeader(errors.New("all control characters must be unique"))
	}

	return Delimiters{Field: fs, Component: cs, Repeat: rs, Escape: ec, SubComponent: ss}, nil
}

// ParseMessage takes input as a `[]byte`, and returns the whole message, and maybe an error.
func ParseMessage(buf []byte) (message Message, err error) {
	d, err := getDelimiters(buf)
	if err != nil {
		return
	}

	var (
		segment      Segment
		field        Field
		fieldItem    FieldItem
		component    Component
		lastPos, pos int
	)

	segment = Segment{
		Field{FieldItem{Component{SubComponent("MSH")}}},
		Field{FieldItem{Component{SubComponent(buf[3:4])}}},
		Field{FieldItem{Component{SubComponent(buf[4:8])}}},
	}

	if len(buf) == 8 {
		message = append(message, segment)
		return message, nil
	}

	if len(buf) > 8 && buf[8] != d.Field {
		return nil, ErrInvalidHeader(fmt.Errorf("invalid character found after header content; expected \\x%02x but got \\x%02x", d.Field, buf[8]))
	}

	commitBuffer := func(force bool) {
		if lastPos != pos || force {
			component = append(component, SubComponent(buf[lastPos:pos]))
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
			segment = append(segment, field)
			field = nil
		}
	}

	commitSegment := func(force bool) {
		commitField(false)

		if segment != nil || force {
			message = append(message, segment)
			segment = nil
		}
	}

	sawNewline := false
	buflen := len(buf)
	for lastPos, pos = 9, 9; pos < buflen; pos++ {
		switch buf[pos] {
		case '\r', '\n':
			if !sawNewline {
				commitSegment(true)
			}
			sawNewline = true
			lastPos++
		case d.Field:
			sawNewline = false
			commitField(true)
			lastPos++
		case d.Repeat:
			sawNewline = false
			commitFieldItem(true)
			lastPos++
		case d.Component:
			sawNewline = false
			commitComponent(true)
			lastPos++
		case d.SubComponent:
			sawNewline = false
			commitBuffer(true)
			lastPos++
		default:
			sawNewline = false
		}
	}

	commitSegment(false)
	return message, nil
}
