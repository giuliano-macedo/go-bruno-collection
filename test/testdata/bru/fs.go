package brufs

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed *.bru
var Fs embed.FS

func ReadExpectedBru(t *testing.T, name string) string {
	data, err := Fs.ReadFile(name + ".bru")
	require.NoError(t, err)
	return string(data)
}
