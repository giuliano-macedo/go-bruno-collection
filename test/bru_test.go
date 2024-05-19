package bruno_test

import (
	"testing"

	"github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"
	brufs "github.com/giuliano-macedo/go-bruno-collection/test/testdata/bru"
	"github.com/giuliano-macedo/go-bruno-collection/test/testdata/collection"
	"github.com/giuliano-macedo/go-bruno-collection/test/testdata/environment"
	"github.com/giuliano-macedo/go-bruno-collection/test/testdata/item"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name        string
	item        *bruno.Item
	environment *bruno.Environment
	collection  *bruno.Collection
}

func TestBruCases(t *testing.T) {
	cases := []testCase{
		{name: "item_post", item: &item.Post},
		{name: "item_get", item: &item.Get},
		{name: "item_without_req", item: &item.WithoutReq},
		{name: "item_without_headers", item: &item.WithoutHeaders},
		{name: "environment_with_variables", environment: &environment.WithVariables},
		{name: "environment_without_secrets", environment: &environment.WithoutSecrets},
		{name: "collection", collection: &collection.AllFields},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedBru := brufs.ReadExpectedBru(t, testCase.name)
			var bru []byte

			switch {
			case testCase.item != nil:
				bru = testCase.item.MarshalBru()
			case testCase.environment != nil:
				bru = testCase.environment.MarshalBru()
			case testCase.collection != nil:
				bru = testCase.collection.MarshalBru()
			}

			require.Equal(t, expectedBru, string(bru))
		})
	}
}

func TestBruMarshalFolder(t *testing.T) {
	require.Empty(t, item.Folder.MarshalBru())
}
