package fixtures

// ---------------------------------------
//    Auto-generated file, do not edit!
// ---------------------------------------

import (
	"github.com/rubenv/fastyaml/parser"
)

type parseBasic struct {
	*parser.Parser
}

func ParseBasic(in string) (*Basic, error) {
	p := &parseBasic{
		Parser: parser.New(in),
	}
	return p.parseBasic()
}

func (p *parseBasic) parseBasic() (*Basic, error) {
	result := &Basic{}

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		case "WeekendInfo":
			p.AdvanceLine()
			o, err := p.parseWeekendInfo()
			if err != nil {
				return nil, err
			}
			result.WeekendInfo = o

		default:
			p.SkipLine()
		}
	}
	return result, nil
}

func (p *parseBasic) parseWeekendInfo() (WeekendInfo, error) {
	result := WeekendInfo{}

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		case "TrackName":

			o, err := p.ReadString()
			if err != nil {
				return WeekendInfo{}, err
			}
			result.TrackName = o
			p.AdvanceLine()
		case "TrackID":

			o, err := p.ReadInt()
			if err != nil {
				return WeekendInfo{}, err
			}
			result.TrackID = o
			p.AdvanceLine()

		default:
			p.SkipLine()
		}
	}
	return result, nil
}
