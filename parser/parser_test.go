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

func TestParseIndex(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	in := `Array:
  - key: 1
    val: 2
  - key: 3
    val: 4
  - key: 5
`

	p := New(in)
	assert.Equal(p.in, in)
	assert.Equal(p.Line, "Array:")
	assert.Equal(p.offset, 7)
	assert.Equal(p.Depth, 0)
	assert.Equal(p.ReadKey(), "Array")
	assert.False(p.IsNew)

	p.AdvanceLine()
	assert.Equal(p.Line, "  - key: 1")
	assert.Equal(p.offset, 7+11)
	assert.Equal(p.Depth, 4)
	assert.Equal(p.ReadKey(), "key")
	assert.True(p.IsNew)

	p.AdvanceLine()
	assert.Equal(p.Line, "    val: 2")
	assert.Equal(p.offset, 7+11+11)
	assert.Equal(p.Depth, 4)
	assert.Equal(p.ReadKey(), "val")
	assert.False(p.IsNew)

	p.AdvanceLine()
	assert.Equal(p.Line, "  - key: 3")
	assert.Equal(p.offset, 7+11+11+11)
	assert.Equal(p.Depth, 4)
	assert.Equal(p.ReadKey(), "key")
	assert.True(p.IsNew)

	p.AdvanceLine()
	assert.Equal(p.Line, "    val: 4")
	assert.Equal(p.offset, 7+11+11+11+11)
	assert.Equal(p.Depth, 4)
	assert.Equal(p.ReadKey(), "val")
	assert.False(p.IsNew)

	p.AdvanceLine()
	assert.Equal(p.Line, "  - key: 5")
	assert.Equal(p.offset, 7+11+11+11+11+11)
	assert.Equal(p.Depth, 4)
	assert.Equal(p.ReadKey(), "key")
	assert.True(p.IsNew)
}
