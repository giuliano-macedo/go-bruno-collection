package environment

import "github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"

var WithoutSecrets = bruno.Environment{
	Name: "WithoutVariables",
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
			Value:   "yes",
			Enabled: true,
			Secret:  false,
			Type:    "text",
		},
	},
}
