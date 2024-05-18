package bruno

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"path"
	"strings"
)

type bruWriter struct {
	writer *bytes.Buffer
}

func newBruWriter() *bruWriter {
	return &bruWriter{
		writer: bytes.NewBuffer([]byte{}),
	}
}

func (writer *bruWriter) Bytes() []byte {
	return writer.writer.Bytes()
}

func (writer *bruWriter) writeBlock(name, attr string, blockContentCallback func(*bruWriter)) {
	if attr != "" {
		fmt.Fprintf(writer.writer, "%v:%v {\n", name, attr)
	} else {
		fmt.Fprintf(writer.writer, "%v {\n", name)
	}
	blockContentCallback(writer)
	fmt.Fprintf(writer.writer, "}\n\n")
}

func (writer *bruWriter) writeField(key, value string) {
	fmt.Fprintf(writer.writer, "  %v: %v\n", key, value)
}

func (writer *bruWriter) writeLiteralBlock(name, attr, value string) {
	writer.writeBlock(name, attr, func(bw *bruWriter) {
		value = strings.ReplaceAll(strings.TrimSpace(value), "\n", "\n  ")
		fmt.Fprintf(writer.writer, "  %v\n", value)
	})
}

func (writer *bruWriter) writeVariables(name string, vars []Variable) {
	if len(vars) > 0 {
		writer.writeBlock("vars", name, func(bw *bruWriter) {
			for _, reqVar := range vars {
				name := reqVar.Name
				if !reqVar.Enabled {
					name = "~" + name
				}
				writer.writeField(name, reqVar.Value)
			}
		})
	}
}

func (item *Item) MarshalBru() (data []byte) {
	if item.IsFolder() {
		return data
	}
	writer := newBruWriter()

	writer.writeBlock("meta", "", func(bw *bruWriter) {
		bw.writeField("name", item.Name)
		bw.writeField("type", item.Type)
		bw.writeField("seq", fmt.Sprint(item.Seq))
	})

	if req := item.Request; req != nil {
		writer.writeBlock(req.Method, "", func(bw *bruWriter) {
			bw.writeField("url", req.URL)
			bw.writeField("body", req.Body.Mode)
			bw.writeField("auth", "none")
		})

		writer.writeLiteralBlock("body", req.Body.Mode, req.Body.Json)
		writer.writeVariables("pre-request", req.Vars.Req)
		writer.writeVariables("pre-response", req.Vars.Res)
		if req.Script.Req != "" {
			writer.writeLiteralBlock("script", "pre-request", req.Script.Req)
		}
		if req.Script.Res != "" {
			writer.writeLiteralBlock("script", "pre-response", req.Script.Res)
		}
	}

	if item.Docs != "" {
		writer.writeLiteralBlock("docs", "", item.Docs)
	}

	return writer.Bytes()
}

func (env *Environment) MarshalBru() []byte {
	writer := newBruWriter()

	writer.writeBlock("vars", "", func(bw *bruWriter) {
		for _, fieldVar := range env.Variables {
			if !fieldVar.Secret {
				writer.writeField(fieldVar.Name, fieldVar.Value)
			}
		}
	})

	return writer.Bytes()
}

func (collection *Collection) MarshalBru() []byte {
	writer := newBruWriter()

	if collection.Docs != "" {
		writer.writeLiteralBlock("docs", "", collection.Docs)
	}
	return writer.Bytes()
}

func writeTarFile(tw *tar.Writer, fileName string, data []byte) error {
	hdr := &tar.Header{
		Name: fileName,
		Mode: 0600,
		Size: int64(len(data)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	_, err := tw.Write(data)
	return err
}

func writeItemsBruTar(tw *tar.Writer, items []Item, dir string) error {
	for _, item := range items {
		fname := item.Name
		if dir == "" {
			fname = path.Join(dir, fname)
		}
		if item.IsFolder() {
			if err := writeItemsBruTar(tw, item.Items, fname); err != nil {
				return err
			}
			continue
		}
		if err := writeTarFile(tw, fname+".bru", item.MarshalBru()); err != nil {
			return err
		}
	}
	return nil
}

func CreateBruTar(collection Collection, writer io.Writer) error {
	tw := tar.NewWriter(writer)

	collectionData, err := json.Marshal(map[string]string{
		"version": collection.Version,
		"name":    collection.Name,
		"type":    "collection",
	})
	if err != nil {
		return err
	}
	if err := writeTarFile(tw, "bruno.json", collectionData); err != nil {
		return err
	}

	for _, env := range collection.Environments {
		if err := writeTarFile(tw, path.Join("environments", env.Name+".bru"), env.MarshalBru()); err != nil {
			return err
		}
	}

	if err := writeItemsBruTar(tw, collection.Items, ""); err != nil {
		return err
	}

	if collection.Docs != "" {
		writeTarFile(tw, "collection.bru", collection.MarshalBru())
	}

	return tw.Close()
}
