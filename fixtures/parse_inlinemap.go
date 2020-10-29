package fixtures

// ---------------------------------------
//    Auto-generated file, do not edit!
// ---------------------------------------

import (
	"github.com/rubenv/fastyaml/parser"
)

type parseInlineMap struct {
	*parser.Parser
}

func ParseInlineMap(in string) (*InlineMap, error) {
	p := &parseInlineMap{
		Parser: parser.New(in),
	}
	return p.parseInlineMap()
}

func (p *parseInlineMap) parseInlineMap() (*InlineMap, error) {
	result := &InlineMap{}
	result.Options = CarSetupOptions{}

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		case "UpdateCount":
			o, err := p.ReadInt()
			if err != nil {
				return nil, err
			}
			result.UpdateCount = o
			p.AdvanceLine()
		default:
			p.AdvanceLine()
			o, err := p.consumeMap2()
			if err != nil {
				return nil, err
			}
			result.Options[key] = o
		}

		if p.IsNew {
			break
		}
	}
	return result, nil
}

func (p *parseInlineMap) consumeMap1() (map[string]string, error) {
	result := make(map[string]string)

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()
		o, err := p.ReadString()
		if err != nil {
			return nil, err
		}
		result[key] = o
		p.AdvanceLine()
	}

	return result, nil
}

func (p *parseInlineMap) consumeMap2() (map[string]map[string]string, error) {
	result := make(map[string]map[string]string)

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()
		p.AdvanceLine()
		o, err := p.consumeMap1()
		if err != nil {
			return nil, err
		}
		result[key] = o
	}

	return result, nil
}
