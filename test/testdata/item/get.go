package item

import "github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"

var Get = bruno.Item{
	Type: "http",
	Name: "RequestName2",
	Seq:  1,
	Request: &bruno.Request{
		URL:    "{{baseUrl}}/request2",
		Method: "GET",
		Headers: []bruno.Header{
			bruno.Header{
				Name:    "Content-Type",
				Value:   "application/json",
				Enabled: true,
			},
		},
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
