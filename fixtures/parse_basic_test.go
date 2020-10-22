package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var inBasic = `
WeekendInfo:
 TrackName: bathurst
 TrackID: 219
`

func TestParseBasic(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	result, err := ParseBasic(inBasic)
	assert.NoError(err)
	assert.NotNil(result)

	assert.Equal("bathurst", result.WeekendInfo.TrackName)
	assert.Equal(219, result.WeekendInfo.TrackID)

	obj := &Basic{}
	err = yaml.Unmarshal([]byte(inBasic), obj)
	assert.NoError(err)

	assert.Equal(result, obj)
}
