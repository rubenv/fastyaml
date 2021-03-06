package fastyaml

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rubenv/fastyaml/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	out, err := Generate("fixtures", fixtures.Basic{})
	assert.NoError(err)
	assert.True(strings.Contains(out, "package fixtures"))
}

func TestGenerateInlineMap(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	out, err := Generate("fixtures", fixtures.InlineMap{})
	assert.NoError(err)
	assert.True(strings.Contains(out, "package fixtures"))

	fmt.Println(out)
}
