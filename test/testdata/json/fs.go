package jsonfs

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed *.json
var Fs embed.FS

func ReadJson(t *testing.T, name string) string {
	jsonBytes, err := Fs.ReadFile(name + ".json")
	require.NoError(t, err)
	return string(jsonBytes)
}
