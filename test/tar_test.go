package bruno_test

import (
	"archive/tar"
	"bytes"
	"io"
	"testing"

	"github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"
	brufs "github.com/giuliano-macedo/go-bruno-collection/test/testdata/bru"
	"github.com/giuliano-macedo/go-bruno-collection/test/testdata/collection"
	"github.com/stretchr/testify/require"
)

func TestCreateBruTar(t *testing.T) {
	const (
		brunoJson      = "bruno.json"
		environmentBru = "environments/EnvironmentName.bru"
		getBru         = "FolderName/RequestName2.bru"
		postBru        = "RequestName1.bru"
		collectionBru  = "collection.bru"
	)

	writer := bytes.NewBuffer([]byte{})
	require.NoError(t, bruno.CreateBruTar(collection.AllFields, writer))

	reader := bytes.NewReader(writer.Bytes())
	tarReader := tar.NewReader(reader)

	tarFiles := map[string]string{}
	tarFileNames := []string{}
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		data, err := io.ReadAll(tarReader)
		require.NoError(t, err)
		tarFiles[header.Name] = string(data)
		tarFileNames = append(tarFileNames, header.Name)
	}
	require.Equal(t, tarFileNames, []string{
		brunoJson,
		environmentBru,
		getBru,
		postBru,
		collectionBru,
	})

	require.JSONEq(t, tarFiles[brunoJson], `{
		"name": "CollectionName",
		"type": "collection",
		"version": "1"
	}`)

	require.Equal(t, tarFiles[environmentBru], brufs.ReadExpectedBru(t, "environment_with_variables"))
	require.Equal(t, tarFiles[getBru], brufs.ReadExpectedBru(t, "item_get"))
	require.Equal(t, tarFiles[postBru], brufs.ReadExpectedBru(t, "item_post"))
	require.Equal(t, tarFiles[collectionBru], brufs.ReadExpectedBru(t, "collection"))
}
