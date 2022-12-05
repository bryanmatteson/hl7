package layout

import (
	"io"

	"hl7"
)

// ParseInput represents the input to a parser.
type ParseInput interface {
	Collect() []hl7.Segment
	Advance() (hl7.Segment, error)
	Retreat() (hl7.Segment, error)
	Peek() (hl7.Segment, error)
	Index() int
}

// MessageParseInput ...
type MessageParseInput struct {
	Message        hl7.Message
	Start          int
	Current        int
	CurrentSegment hl7.Segment
}

// NewInput ...
func NewInput(msg hl7.Message) ParseInput {
	return &MessageParseInput{Message: msg}
}

// Collect collects all of the segments parsed so far and returns it, then starts a new collection from the current position in the input.
func (input *MessageParseInput) Collect() (segments []hl7.Segment) {
	segments = input.Message[input.Start:input.Current]
	input.Start = input.Current
	return
}

// Advance advances the input by a single segment and consumes it.
func (input *MessageParseInput) Advance() (seg hl7.Segment, err error) {
	if input.Current >= len(input.Message) {
		return nil, io.EOF
	}
	input.Current++
	input.CurrentSegment = input.Message[input.Current-1]
	seg = input.CurrentSegment
	return
}

// Retreat retreats the input position by a single segment and unconsumes it.
func (input *MessageParseInput) Retreat() (seg hl7.Segment, err error) {
	input.Current--
	input.CurrentSegment = input.Message[input.Current]
	seg = input.CurrentSegment
	return
}

// Peek returns the next segment from the input without consuming it.
func (input *MessageParseInput) Peek() (seg hl7.Segment, err error) {
	seg, err = input.Advance()
	if err != nil {
		return
	}
	_, err = input.Retreat()
	// if err != ErrBeginningOfMessage {
	// 	return
	// }

	return seg, nil
}

// PeekN reads a number of segments from the Input, then sets the current position back.
func (input *MessageParseInput) PeekN(n int) (segs []hl7.Segment, err error) {
	segs = make([]hl7.Segment, n)
	var advanced int

	for i := 0; i < n; i++ {
		segs[i], err = input.Advance()
		if err != nil {
			break
		}
		advanced++
	}

	for i := 0; i < advanced; i++ {
		input.Retreat()
	}

	return
}

// Index returns the current index of the parser input.
func (input *MessageParseInput) Index() int {
	return input.Current
}
