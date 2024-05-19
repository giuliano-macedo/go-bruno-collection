package item

import "github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"

var Folder = bruno.Item{
	Type:  "folder",
	Name:  "FolderName",
	Items: []bruno.Item{Get},
}
