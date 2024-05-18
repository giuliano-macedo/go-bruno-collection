package bruno_test

import (
	"encoding/json"
	"fmt"
	"testing"

	bruno "github.com/giuliano-macedo/go-bruno-collection/pkg"
	"github.com/giuliano-macedo/go-bruno-collection/pkg/test/testdata"
	"github.com/stretchr/testify/require"
)

func TestCollectionJson(t *testing.T) {
	testCases := []testdata.TestCase{
		testdata.AllFields,
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				encodedJson, err := json.Marshal(testCase.Collection)
				require.NoError(t, err)
				fmt.Println(string(encodedJson))
				require.JSONEq(t, testCase.ExpectedJson, string(encodedJson))
			})

			t.Run("unmarshal", func(t *testing.T) {
				var expectedCollection bruno.Collection
				err := json.Unmarshal([]byte(testCase.ExpectedJson), &expectedCollection)
				require.NoError(t, err)
				require.Equal(t, expectedCollection, testCase.Collection)
			})
		})
	}
}
