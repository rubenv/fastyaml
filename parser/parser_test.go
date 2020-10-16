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

	p := NewParseInfo(in)
	assert.Equal(p.in, in)
	assert.Equal(p.line, "WeekendInfo:")
	assert.Equal(p.offset, 13)
	assert.Equal(p.depth, 0)
	assert.Equal(p.ReadKey(), "WeekendInfo")

	p.ReadLine()
	assert.Equal(p.line, "  TrackName: bathurst")
	assert.Equal(p.offset, 13+22)
	assert.Equal(p.depth, 2)
	assert.Equal(p.ReadKey(), "TrackName")

	p.ReadLine()
	assert.Equal(p.line, "  TrackID: 219")
	assert.Equal(p.offset, 13+22+1+15)
	assert.Equal(p.depth, 2)
	assert.Equal(p.ReadKey(), "TrackID")

	p.ReadLine()
	assert.Equal(p.line, "  TrackWeather: sunny")
	assert.Equal(p.offset, 13+22+1+15+22)
	assert.Equal(p.depth, 2)
	assert.Equal(p.ReadKey(), "TrackWeather")

	p.ReadLine()
	assert.Equal(p.line, "SessionInfo: active")
	assert.Equal(p.offset, 13+22+1+15+22+20)
	assert.Equal(p.depth, 0)
	assert.Equal(p.ReadKey(), "SessionInfo")

	p.ReadLine()
	assert.Equal(p.line, "")
	assert.Equal(p.depth, -1)
}
