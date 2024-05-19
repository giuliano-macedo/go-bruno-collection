package bruno

import (
	"fmt"
	"strings"

	"github.com/giuliano-macedo/go-bruno-collection/internal/bru"
)

func writeVariables(writer *bru.Writer, name string, vars []Variable) {
	if len(vars) == 0 {
		return
	}
	writer.WriteBlock("vars", name, func(bw *bru.Writer) {
		for _, reqVar := range vars {
			writer.WriteEnableableField(reqVar.Name, reqVar.Value, reqVar.Enabled)
		}
	})
}

func writeHeaders(writer *bru.Writer, headers []Header) {
	if len(headers) == 0 {
		return
	}
	writer.WriteBlock("headers", "", func(bw *bru.Writer) {
		for _, header := range headers {
			bw.WriteEnableableField(header.Name, header.Value, header.Enabled)
		}
	})
}

func writeAuth(writer *bru.Writer, auth Auth) {
	switch auth.Mode {
	case "none":
	}
	basic := auth.Basic
	writer.WriteBlock("auth", "basic", func(bw *bru.Writer) {
		bw.WriteField("username", basic.Username)
		bw.WriteField("password", basic.Password)
	})
	writer.WriteBlock("auth", "bearer", func(bw *bru.Writer) {
		bw.WriteField("token", auth.Bearer.Token)
	})
}

func converItemType(mode string) string {
	switch mode {
	case "http-request":
		return "http"
	}
	return mode
}

func writeRequest(writer *bru.Writer, req *Request) {
	if req == nil {
		return
	}

	writer.WriteBlock(strings.ToLower(req.Method), "", func(bw *bru.Writer) {
		bw.WriteField("url", req.URL)
		bw.WriteField("body", req.Body.Mode)
		bw.WriteField("auth", "none")
	})

	writeHeaders(writer, req.Headers)
	writeAuth(writer, req.Auth)

	if req.Body.Mode != "" && req.Body.Mode != "none" {
		writer.WriteLiteralBlock("body", req.Body.Mode, req.Body.Json)
	}
	writeVariables(writer, "pre-request", req.Vars.Req)
	writeVariables(writer, "post-response", req.Vars.Res)
	if req.Script.Req != "" {
		writer.WriteLiteralBlock("script", "pre-request", req.Script.Req)
	}
	if req.Script.Res != "" {
		writer.WriteLiteralBlock("script", "post-response", req.Script.Res)
	}
	if req.Docs != "" {
		writer.WriteLiteralBlock("docs", "", req.Docs)
	}
}

func writeEnvironment(writer *bru.Writer, env Environment) {
	secretValues := make([]string, 0, len(env.Variables))
	writer.WriteBlock("vars", "", func(bw *bru.Writer) {
		for _, fieldVar := range env.Variables {
			if fieldVar.Secret {
				secretValues = append(secretValues, fieldVar.Name)
			} else {
				writer.WriteEnableableField(fieldVar.Name, fieldVar.Value, fieldVar.Enabled)
			}
		}
	})
	writer.WriteArrayBlock("vars", "secret", secretValues)
}

func (item *Item) MarshalBru() (data []byte) {
	if item.IsFolder() {
		return data
	}

	writer := bru.NewWriter()
	writer.WriteBlock("meta", "", func(bw *bru.Writer) {
		bw.WriteField("name", item.Name)
		bw.WriteField("type", converItemType(item.Type))
		bw.WriteField("seq", fmt.Sprint(item.Seq))
	})

	writeRequest(writer, item.Request)

	return writer.Bytes()
}

func (env *Environment) MarshalBru() []byte {
	writer := bru.NewWriter()

	writeEnvironment(writer, *env)

	return writer.Bytes()
}

func (collection *Collection) MarshalBru() []byte {
	writer := bru.NewWriter()

	if collection.Docs != "" {
		writer.WriteLiteralBlock("docs", "", collection.Docs)
	}
	return writer.Bytes()
}
