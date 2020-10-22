package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func BenchmarkParseGenerated(b *testing.B) {
	assert := assert.New(b)
	for i := 0; i < b.N; i++ {
		_, err := ParseBasic(inBasic)
		assert.NoError(err)
	}
}

func BenchmarkParseYAML(b *testing.B) {
	assert := assert.New(b)
	for i := 0; i < b.N; i++ {
		obj := &Basic{}
		err := yaml.Unmarshal([]byte(inBasic), obj)
		assert.NoError(err)
	}
}
