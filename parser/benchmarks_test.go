package parser

import "testing"

func BenchmarkReadLine(b *testing.B) {
	in := `WeekendInfo:
  TrackName: bathurst

  TrackID: 219
  TrackWeather: sunny
SessionInfo: active
`
	p := NewParseInfo(in)

	for i := 0; i < b.N; i++ {
		for p.depth != -1 {
			p.ReadLine()
		}
	}
}

func BenchmarkReadKey(b *testing.B) {
	in := `WeekendInfo:
  TrackName: bathurst

  TrackID: 219
  TrackWeather: sunny
SessionInfo: active
`
	p := NewParseInfo(in)

	for i := 0; i < b.N; i++ {
		for p.depth != -1 {
			p.ReadLine()
			p.ReadKey()
		}
	}
}
