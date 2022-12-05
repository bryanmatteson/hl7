package hl7

import "io"

type Scanner struct {
	reader  *Reader
	segment Segment
	done    bool
	err     error
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		reader: NewReader(r),
	}
}

func (s *Scanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

func (s *Scanner) Segment() Segment {
	return s.segment
}

func (s *Scanner) Scan() bool {
	if s.done {
		return false
	}
	s.segment, s.err = s.reader.ReadSegment()
	if s.err == io.EOF {
		s.done = true
		return true
	}
	if s.err != nil {
		return false
	}
	return true
}
