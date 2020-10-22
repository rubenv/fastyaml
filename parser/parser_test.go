package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	in := `WeekendInfo:
  TrackName: bathurst

  TrackID: 219
  TrackWeather: sunny
SessionInfo: active
`

	p := New(in)
	assert.Equal(p.in, in)
	assert.Equal(p.Line, "WeekendInfo:")
	assert.Equal(p.offset, 13)
	assert.Equal(p.Depth, 0)
	assert.Equal(p.ReadKey(), "WeekendInfo")

	p.AdvanceLine()
	assert.Equal(p.Line, "  TrackName: bathurst")
	assert.Equal(p.offset, 13+22)
	assert.Equal(p.Depth, 2)
	assert.Equal(p.ReadKey(), "TrackName")

	p.AdvanceLine()
	assert.Equal(p.Line, "  TrackID: 219")
	assert.Equal(p.offset, 13+22+1+15)
	assert.Equal(p.Depth, 2)
	assert.Equal(p.ReadKey(), "TrackID")

	p.AdvanceLine()
	assert.Equal(p.Line, "  TrackWeather: sunny")
	assert.Equal(p.offset, 13+22+1+15+22)
	assert.Equal(p.Depth, 2)
	assert.Equal(p.ReadKey(), "TrackWeather")

	p.AdvanceLine()
	assert.Equal(p.Line, "SessionInfo: active")
	assert.Equal(p.offset, 13+22+1+15+22+20)
	assert.Equal(p.Depth, 0)
	assert.Equal(p.ReadKey(), "SessionInfo")

	p.AdvanceLine()
	assert.Equal(p.Line, "")
	assert.Equal(p.Depth, -1)
}
