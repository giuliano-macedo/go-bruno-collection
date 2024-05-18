package testdata

import bruno "github.com/giuliano-macedo/go-bruno-collection/pkg"

var AllFields = TestCase{
	"all fields",
	bruno.Collection{
		Name:    "CollectionName",
		Version: "1",
		Items: []bruno.Item{
			{
				Type: "folder",
				Name: "FolderName",
				Items: []bruno.Item{
					{
						Type: "http",
						Name: "RequestName2",
						Seq:  1,
						Request: &bruno.Request{
							URL:     "{{baseUrl}}/request2",
							Method:  "GET",
							Headers: []string{},
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
							Query:      []string{},
						},
					},
				},
			},
			{
				Type: "http",
				Name: "RequestName1",
				Seq:  1,
				Request: &bruno.Request{
					URL:     "{{baseUrl}}/request1",
					Method:  "POST",
					Headers: []string{},
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
					Query:      []string{},
				},
			},
		},
		ActiveEnvironmentUid: "CmLVov8jg1StssHjThi2G",
		Environments: []bruno.Environment{
			{
				Variables: []bruno.EnvironmentVariable{
					{
						Name:    "baseUrl",
						Value:   "https://example.com",
						Enabled: true,
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
			},
		},
	},
	`
	{
		"name": "CollectionName",
		"version": "1",
		"items": [
		  {
			"type": "folder",
			"name": "FolderName",
			"items": [
			  {
				"type": "http",
				"name": "RequestName2",
				"seq": 1,
				"request": {
				  "url": "{{baseUrl}}/request2",
				  "method": "GET",
				  "headers": [],
				  "body": {
					"mode": "none",
					"formUrlEncoded": [],
					"multipartForm": []
				  },
				  "auth": {
					"mode": "none",
					"basic": {
					  "username": "",
					  "password": ""
					},
					"bearer": {
					  "token": ""
					}
				  },
				  "script": {},
				  "vars": {},
				  "assertions": [],
				  "tests": "",
				  "query": []
				}
			  }
			]
		  },
		  {
			"type": "http",
			"name": "RequestName1",
			"seq": 1,
			"request": {
			  "url": "{{baseUrl}}/request1",
			  "method": "POST",
			  "headers": [],
			  "body": {
				"mode": "json",
				"json": "{\n  \"hello\": \"world\",\n  \"var\": \"{{myVariable}}\"\n}",
				"formUrlEncoded": [],
				"multipartForm": []
			  },
			  "auth": {
				"mode": "none",
				"basic": {
				  "username": "",
				  "password": ""
				},
				"bearer": {
				  "token": ""
				}
			  },
			  "script": {
				"req": "console.log(\"hellow rold\")",
				"res": "console.log(\"hellow rold\")"
			  },
			  "vars": {
				"req": [
				  {
					"name": "x",
					"value": "1",
					"enabled": true,
					"local": false
				  }
				],
				"res": [
				  {
					"name": "y",
					"value": "2",
					"enabled": true,
					"local": false
				  }
				]
			  },
			  "assertions": [],
			  "tests": "",
			  "query": []
			}
		  }
		],
		"activeEnvironmentUid": "CmLVov8jg1StssHjThi2G",
		"environments": [
		  {
			"variables": [
			  {
				"name": "baseUrl",
				"value": "https://example.com",
				"enabled": true,
				"secret": false,
				"type": "text"
			  },
			  {
				"name": "myVariable",
				"value": "",
				"enabled": true,
				"secret": true,
				"type": "text"
			  }
			],
			"name": "EnvironmentName"
		  }
		]
	  }
	`,
}
