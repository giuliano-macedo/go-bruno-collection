package environment

import "github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"

var WithVariables = bruno.Environment{
	Variables: []bruno.EnvironmentVariable{
		{
			Name:    "baseUrl",
			Value:   "https://example.com",
			Enabled: true,
			Secret:  false,
			Type:    "text",
		},
		{
			Name:    "disabled",
			Value:   "1",
			Enabled: false,
			Secret:  false,
			Type:    "text",
		},
		{
			Name:    "myVariable",
			Value:   "",
			Enabled: true,
			Secret:  true,
			Type:    "text",
		},
	},
	Name: "EnvironmentName",
}
