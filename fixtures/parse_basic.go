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

func ParseBasicFromString(in string) (*Basic, error) {
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

		case "SessionInfo":
			p.AdvanceLine()
			o, err := p.parseSessionInfoDetails()
			if err != nil {
				return nil, err
			}
			result.SessionInfo = o

		default:
			p.SkipValue()
		}

		if p.IsNew {
			break
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
			p.SkipValue()
		}

		if p.IsNew {
			break
		}
	}
	return result, nil
}

func (p *parseBasic) parseSessionInfoDetails() (SessionInfoDetails, error) {
	result := SessionInfoDetails{}

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		case "Sessions":
			p.AdvanceLine()
			ad := p.Depth
			for ad == p.Depth {
				o, err := p.parseSessionInfoSession()
				if err != nil {
					return SessionInfoDetails{}, err
				}
				result.Sessions = append(result.Sessions, o)
			}
		default:
			p.SkipValue()
		}

		if p.IsNew {
			break
		}
	}
	return result, nil
}

func (p *parseBasic) parseSessionInfoSession() (SessionInfoSession, error) {
	result := SessionInfoSession{}

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		case "SessionType":
			o, err := p.ReadString()
			if err != nil {
				return SessionInfoSession{}, err
			}
			result.SessionType = o
			p.AdvanceLine()
		case "ResultsPositions":
			p.AdvanceLine()
			ad := p.Depth
			for ad == p.Depth {
				o, err := p.parseSessionResultsPosition()
				if err != nil {
					return SessionInfoSession{}, err
				}
				result.ResultsPositions = append(result.ResultsPositions, o)
			}
		default:
			p.SkipValue()
		}

		if p.IsNew {
			break
		}
	}
	return result, nil
}

func (p *parseBasic) parseSessionResultsPosition() (SessionResultsPosition, error) {
	result := SessionResultsPosition{}

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		case "Position":
			o, err := p.ReadInt()
			if err != nil {
				return SessionResultsPosition{}, err
			}
			result.Position = o
			p.AdvanceLine()
		case "ClassPosition":
			o, err := p.ReadInt()
			if err != nil {
				return SessionResultsPosition{}, err
			}
			result.ClassPosition = o
			p.AdvanceLine()
		case "Time":
			o, err := p.ReadFloat64()
			if err != nil {
				return SessionResultsPosition{}, err
			}
			result.Time = o
			p.AdvanceLine()
		case "FastestTime":
			o, err := p.ReadFloat32()
			if err != nil {
				return SessionResultsPosition{}, err
			}
			result.FastestTime = o
			p.AdvanceLine()
		default:
			p.SkipValue()
		}

		if p.IsNew {
			break
		}
	}
	return result, nil
}
