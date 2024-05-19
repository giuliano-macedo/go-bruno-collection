package item

import "github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"

var WithoutHeaders = bruno.Item{
	Type: "http",
	Name: "WithoutHeaders",
	Seq:  1,
	Request: &bruno.Request{
		URL:    "{{baseUrl}}/request2",
		Method: "GET",
		Body: bruno.Body{
			Mode:           "none",
			FormUrlEncoded: []string{},
			MultipartForm:  []string{},
		},
		Auth: bruno.Auth{
			Mode: "none",
			Basic: bruno.Basic{
				Username: "",
				Password: "",
			},
			Bearer: bruno.Bearer{
				Token: "",
			},
		},
		Script:     bruno.Script{},
		Vars:       bruno.Vars{},
		Assertions: []string{},
		Tests:      "",
	},
}
