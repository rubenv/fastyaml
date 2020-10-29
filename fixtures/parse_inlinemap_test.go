package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var inInlineMap = `
UpdateCount: 2
Tires:
 LeftFront:
  StartingPressure: 145 kPa
  LastHotPressure: 145 kPa
  LastTempsOMI: 77C, 77C, 77C
  TreadRemaining: 100%, 100%, 100%
 LeftRear:
  StartingPressure: 145 kPa
  LastHotPressure: 145 kPa
  LastTempsOMI: 77C, 77C, 77C
  TreadRemaining: 100%, 100%, 100%
 RightFront:
  StartingPressure: 145 kPa
  LastHotPressure: 145 kPa
  LastTempsIMO: 77C, 77C, 77C
  TreadRemaining: 100%, 100%, 100%
`

func TestParseInlineMap(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	result, err := ParseInlineMap(inInlineMap)
	assert.NoError(err)
	assert.NotNil(result)

	assert.Equal(2, result.UpdateCount)
	assert.Len(result.Options, 1)
	assert.NotEmpty(result.Options["Tires"])

	obj := &InlineMap{}
	err = yaml.Unmarshal([]byte(inInlineMap), obj)
	assert.NoError(err)

	assert.Equal(result, obj)
}
