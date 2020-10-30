package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func BenchmarkParseBasicGenerated(b *testing.B) {
	assert := assert.New(b)
	for i := 0; i < b.N; i++ {
		_, err := ParseBasicFromString(inBasic)
		assert.NoError(err)
	}
}

func BenchmarkParseBasicYAML(b *testing.B) {
	assert := assert.New(b)
	for i := 0; i < b.N; i++ {
		obj := &Basic{}
		err := yaml.Unmarshal([]byte(inBasic), obj)
		assert.NoError(err)
	}
}

func BenchmarkParseInlineMapGenerated(b *testing.B) {
	assert := assert.New(b)
	for i := 0; i < b.N; i++ {
		_, err := ParseInlineMapFromString(inInlineMap)
		assert.NoError(err)
	}
}

func BenchmarkParseInlineMapYAML(b *testing.B) {
	assert := assert.New(b)
	for i := 0; i < b.N; i++ {
		obj := &InlineMap{}
		err := yaml.Unmarshal([]byte(inInlineMap), obj)
		assert.NoError(err)
	}
}
