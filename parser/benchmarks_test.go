package parser

import "testing"

func BenchmarkReadLine(b *testing.B) {
	in := `WeekendInfo:
  TrackName: bathurst

  TrackID: 219
  TrackWeather: sunny
SessionInfo: active
`
	p := New(in)

	for i := 0; i < b.N; i++ {
		for p.Depth != -1 {
			p.AdvanceLine()
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
	p := New(in)

	for i := 0; i < b.N; i++ {
		for p.Depth != -1 {
			p.AdvanceLine()
			p.ReadKey()
		}
	}
}
