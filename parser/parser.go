package parser

type ParseInfo struct {
	in     string
	line   string
	offset int

	depth int
	isNew bool
}

func NewParseInfo(in string) *ParseInfo {
	p := &ParseInfo{
		in: in,
	}
	p.ReadLine()
	return p
}

func (p *ParseInfo) ReadLine() {
	for {
		if p.offset == len(p.in) {
			p.line = ""
			p.depth = -1
			return
		}
		p.line = ReadUntil(p.in, p.offset, '\n')
		p.offset += len(p.line) + 1

		allWS := true
		depth := 0
		for _, s := range p.line {
			if s == ' ' || s == '\t' {
				depth += 1
			} else {
				allWS = false
				break
			}
		}
		if allWS {
			continue
		}
		p.depth = depth

		break
	}
}
