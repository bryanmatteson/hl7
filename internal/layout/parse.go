package layout

import (
	"fmt"
	"reflect"

	"hl7"

	"github.com/pkg/errors"
)

type Parser interface {
	Name() string
	Type() reflect.Type
	Parse(ParseInput) ParseResult
}

type segmentParser struct{ name string }

func (p segmentParser) Type() reflect.Type { return reflect.ValueOf([]hl7.Segment{}).Type().Elem() }
func (p segmentParser) Name() string       { return fmt.Sprintf("Segment(%s)", p.name) }
func (p segmentParser) Parse(pi ParseInput) ParseResult {
	seg, err := pi.Peek()
	if err != nil {
		return Failure(p.Name(), err)
	}

	if seg.Name() == p.name {
		_, err = pi.Advance()
		return Success(p.Name(), seg, err)
	}

	return Failure(p.Name(), fmt.Errorf("expected segment %s, got %s", p.name, seg.Name()))
}

type GroupBuilder interface {
	Parser
	With(Parser) GroupBuilder
	Segment(string) GroupBuilder
	Maybe(Parser) GroupBuilder
	List(Parser) GroupBuilder
}

type groupParser struct {
	name    string
	parsers []Parser
}

func Group(n string) GroupBuilder { return &groupParser{n, nil} }

func (p groupParser) Type() reflect.Type { return reflect.ValueOf([]hl7.Group{}).Type().Elem() }
func (p groupParser) Name() string       { return fmt.Sprintf("Group(%s)", p.name) }
func (p groupParser) Parse(pi ParseInput) ParseResult {
	name := fmt.Sprintf("Group %s", p.name)
	results := make(hl7.Group, len(p.parsers))
	start := pi.Index()
	for i, f := range p.parsers {
		r := f.Parse(pi)
		if !r.Success {
			rewind(pi, int(pi.Index()-start))
			r.Error = errors.WithMessage(r.Error, name)
			return r
		}
		if r.Item == nil {
			results[i] = nil
		} else {
			results[i] = r.Item.(hl7.HL7Type)
		}
	}
	return Success(p.Name(), results, nil)
}

func (p *groupParser) add(parser Parser) GroupBuilder {
	p.parsers = append(p.parsers, parser)
	return p
}

func (p *groupParser) With(parser Parser) GroupBuilder  { return p.add(parser) }
func (p *groupParser) Segment(name string) GroupBuilder { return p.add(Segment(name)) }
func (p *groupParser) Maybe(parser Parser) GroupBuilder { return p.add(Maybe(parser)) }
func (p *groupParser) List(parser Parser) GroupBuilder  { return p.add(List(parser)) }

type maybeParser struct{ parser Parser }

func (p maybeParser) Type() reflect.Type { return p.parser.Type() }
func (p maybeParser) Name() string       { return "Maybe" + p.parser.Name() }
func (p maybeParser) Parse(pi ParseInput) ParseResult {
	return Success(p.Name(), p.parser.Parse(pi).Item, nil)
}

type listParser struct {
	parser   Parser
	min, max int
}

func (p listParser) Type() reflect.Type {
	switch p.parser.Type().Name() {
	case "Group":
		return reflect.TypeOf(hl7.GroupList{})
	case "Segment":
		return reflect.TypeOf(hl7.SegmentList{})
	default:
		panic("??")
	}
}
func (p listParser) Name() string { return "List" + p.parser.Name() }
func (p listParser) Parse(pi ParseInput) ParseResult {
	count := 0
	slice := reflect.MakeSlice(p.Type(), 0, 0)

	start := pi.Index()
	for {
		r := p.parser.Parse(pi)
		if !r.Success {
			if count < p.min {
				// Roll back, because we didn't get enough.
				rewind(pi, pi.Index()-start)
				r.Error = errors.WithMessagef(r.Error, "while parsing %s", p.Name())
				return r
			}
			break
		}

		count++
		slice = reflect.Append(slice, reflect.ValueOf(r.Item))

		if p.max > 0 && count == p.max {
			break
		}
	}

	values := slice.Interface()
	return Success(p.Name(), values, nil)
}

// ParseResult is the result of a parse operation.
type ParseResult struct {
	Name    string
	Success bool
	Item    interface{}
	Error   error
}

// Success creates a successful result of a parse operation.
func Success(name string, item interface{}, err error) ParseResult {
	return ParseResult{
		Name:    name,
		Success: true,
		Item:    item,
		Error:   err,
	}
}

// Failure creates an unsuccessful result of a parse operation.
func Failure(name string, err error) ParseResult {
	return ParseResult{
		Name:    name,
		Success: false,
		Error:   err,
	}
}

func Segment(n string) Parser { return &segmentParser{n} }
func Maybe(p Parser) Parser   { return &maybeParser{p} }
func List(p Parser) Parser    { return &listParser{p, 1, -1} }

func rewind(pi ParseInput, times int) (err error) {
	for i := 0; i < times; i++ {
		_, err = pi.Retreat()
		if err != nil {
			return
		}
	}
	return
}
