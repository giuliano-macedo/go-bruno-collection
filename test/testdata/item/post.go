package item

import "github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"

var Post = bruno.Item{
	Type: "http-request",
	Name: "RequestName1",
	Seq:  1,
	Request: &bruno.Request{
		URL:    "{{baseUrl}}/request1",
		Method: "POST",
		Headers: []bruno.Header{
			{Name: "Content-Type", Value: "application/json", Enabled: true},
			{Name: "Disabled", Value: "foo", Enabled: false},
		},
		Body: bruno.Body{
			Mode:           "json",
			Json:           "{\n  \"hello\": \"world\",\n  \"var\": \"{{myVariable}}\"\n}",
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
		Script: bruno.Script{
			Req: "console.log(\"hellow rold\")",
			Res: "console.log(\"hellow rold\")",
		},
		Vars: bruno.Vars{
			Req: []bruno.Variable{
				{
					Name:    "x",
					Value:   "1",
					Enabled: true,
					Local:   false,
				},
			},
			Res: []bruno.Variable{
				{
					Name:    "y",
					Value:   "2",
					Enabled: true,
					Local:   false,
				},
			},
		},
		Assertions: []string{},
		Tests:      "",
		Docs:       "Request docs",
	},
}
