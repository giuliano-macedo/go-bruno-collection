package testdata

import (
	bruno "github.com/giuliano-macedo/go-bruno-collection/pkg"
)

type TestCase struct {
	Name         string
	Collection   bruno.Collection
	ExpectedJson string
}
