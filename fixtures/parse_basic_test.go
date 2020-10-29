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

SessionInfo:
 Sessions:
 - SessionNum: 0
   SessionType: Practice
   SessionTrackRubberState: slight usage
   ResultsPositions:
   - Position: 1
     ClassPosition: 0
     CarIdx: 6
     Lap: 13
     Time: 122.9443
     FastestTime: 122.9443
   - Position: 2
     ClassPosition: 1
     Time: 122.9713
     FastestLap: 4
     FastestTime: 122.9713
     LastTime: 123.4722
   - Position: 3
     ClassPosition: 2
     Time: 123.1166
     FastestTime: 123.1166
   - Position: 4
     ClassPosition: 3
     CarIdx: 17
     Lap: 2
`

func TestParseBasic(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	result, err := ParseBasic(inBasic)
	assert.NoError(err)
	assert.NotNil(result)

	assert.Equal("bathurst", result.WeekendInfo.TrackName)
	assert.Equal(219, result.WeekendInfo.TrackID)

	assert.Len(result.SessionInfo.Sessions, 1)
	assert.Len(result.SessionInfo.Sessions[0].ResultsPositions, 4)

	obj := &Basic{}
	err = yaml.Unmarshal([]byte(inBasic), obj)
	assert.NoError(err)

	assert.Equal(result, obj)
}
