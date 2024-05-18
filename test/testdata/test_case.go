package testdata

import (
	bruno "github.com/giuliano-macedo/go-bruno-collection"
)

type TestCase struct {
	Name         string
	Collection   bruno.Collection
	ExpectedJson string
}
