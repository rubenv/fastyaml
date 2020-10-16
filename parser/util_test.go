package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUntil(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	in := `WeekendInfo:
 TrackName: bathurst
 TrackID: 219
`

	name := ReadUntil(in, 0, ':')
	assert.Equal(name, "WeekendInfo")
}
