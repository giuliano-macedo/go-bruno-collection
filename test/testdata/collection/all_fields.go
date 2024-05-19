package collection

import (
	"github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"
	"github.com/giuliano-macedo/go-bruno-collection/test/testdata/environment"
	"github.com/giuliano-macedo/go-bruno-collection/test/testdata/item"
)

var AllFields = bruno.Collection{
	Name:                 "CollectionName",
	Version:              "1",
	Items:                []bruno.Item{item.Folder, item.Post},
	ActiveEnvironmentUid: "CmLVov8jg1StssHjThi2G",
	Environments:         []bruno.Environment{environment.WithVariables},
	Docs: `Collection docs [test](google.com)

# h1

## H2

* Yes
* Yes`,
}
