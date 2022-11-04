package formatters

import (
	"testing"

	"github.com/Hiroya-W/chroma/v2"
	assert "github.com/alecthomas/assert/v2"
)

func TestClosestColour(t *testing.T) {
	actual := findClosest(ttyTables[256], chroma.MustParseColour("#e06c75"))
	assert.Equal(t, chroma.MustParseColour("#d75f87"), actual)
}
