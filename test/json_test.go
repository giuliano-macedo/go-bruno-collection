package bruno

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"
	"github.com/giuliano-macedo/go-bruno-collection/test/testdata/collection"
	jsonfs "github.com/giuliano-macedo/go-bruno-collection/test/testdata/json"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name         string
	collection   bruno.Collection
	expectedJson string
}

func loadTestCase(t *testing.T, name string, collection bruno.Collection) testCase {
	return testCase{
		name:         name,
		collection:   collection,
		expectedJson: jsonfs.ReadJson(t, name),
	}
}

func TestCollectionJson(t *testing.T) {
	testCases := []testCase{
		loadTestCase(t, "all_fields", collection.AllFields),
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				encodedJson, err := json.Marshal(testCase.collection)
				require.NoError(t, err)
				fmt.Println(string(encodedJson))
				require.JSONEq(t, testCase.expectedJson, string(encodedJson))
			})

			t.Run("unmarshal", func(t *testing.T) {
				actualCollection := testCase.collection
				actualCollection.Docs = ""

				var expectedCollection bruno.Collection
				err := json.Unmarshal([]byte(testCase.expectedJson), &expectedCollection)
				require.NoError(t, err)
				require.Equal(t, expectedCollection, actualCollection)
			})
		})
	}
}
