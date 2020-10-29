package parser

import (
	"strconv"
	"strings"
)

type Parser struct {
	in     string
	offset int

	Line  string
	Depth int
	IsNew bool

	valueOffset int
}

func New(in string) *Parser {
	p := &Parser{
		in: in,
	}
	p.AdvanceLine()
	return p
}

func (p *Parser) AdvanceLine() {
	for {
		if p.offset == len(p.in) {
			p.Line = ""
			p.Depth = -1
			return
		}
		p.Line = ReadUntil(p.in, p.offset, '\n')
		p.IsNew = false
		p.offset += len(p.Line) + 1

		allWS := true
		depth := 0
		for _, s := range p.Line {
			if s == ' ' || s == '\t' {
				depth += 1
			} else if s == '-' {
				depth += 1
				p.IsNew = true
			} else {
				allWS = false
				break
			}
		}
		if allWS {
			continue
		}
		p.Depth = depth

		break
	}
}

// Advances a line and skips any deeper children
func (p *Parser) SkipLine() {
	p.AdvanceLine()
}

func (p *Parser) ReadKey() string {
	if p.Depth < 0 {
		return ""
	}
	key := ReadUntil(p.Line, p.Depth, ':')
	p.valueOffset = p.Depth + len(key) + 1
	return key
}

func (p *Parser) ReadString() (string, error) {
	s := strings.TrimSpace(p.Line[p.valueOffset:])
	return s, nil
}

func (p *Parser) ReadInt64() (int64, error) {
	s, err := p.ReadString()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(s, 10, 64)
}

func (p *Parser) ReadInt() (int, error) {
	i, err := p.ReadInt64()
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func (p *Parser) ReadFloat32() (float32, error) {
	f, err := p.ReadFloat64()
	return float32(f), err
}

func (p *Parser) ReadFloat64() (float64, error) {
	s, err := p.ReadString()
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return 0, err
	}
	return f, nil
}
