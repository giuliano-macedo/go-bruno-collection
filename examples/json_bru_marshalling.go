package main

import (
	"encoding/json"
	"os"

	"github.com/giuliano-macedo/go-bruno-collection/pkg/bruno"
)

func createMyCollection() bruno.Collection {
	itemReq := &bruno.Request{
		URL:    "{{baseUrl}}/request1",
		Method: "POST",
		Headers: []bruno.Header{
			{Name: "Content-Type", Value: "application/json", Enabled: true},
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
	}

	item := bruno.Item{
		Type:    "http-request",
		Name:    "MyAwesomeRequest",
		Seq:     1,
		Request: itemReq,
	}

	return bruno.Collection{
		Name:                 "MyAwesomeCollection",
		Version:              "1",
		Items:                []bruno.Item{item},
		ActiveEnvironmentUid: "0b6ec80b0beb42499f0cf",
		Environments: []bruno.Environment{{
			Variables: []bruno.EnvironmentVariable{
				{
					Name:    "baseUrl",
					Value:   "https://example.com",
					Enabled: true,
					Secret:  false,
					Type:    "text",
				},
			},
			Name: "MyEnvironment",
		}},
		Docs: `HelloWorld`,
	}
}

func saveBrunoJsonCollection(collection bruno.Collection, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	json.NewEncoder(f).Encode(collection)
	return f.Close()
}

func saveBruTarFile(collection bruno.Collection, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	return bruno.CreateBruTar(collection, f)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	collection := createMyCollection()

	err := saveBrunoJsonCollection(collection, "bruno_collection.json")
	handleError(err)
	err = saveBruTarFile(collection, "bruno_collection.tar")
	handleError(err)
}
