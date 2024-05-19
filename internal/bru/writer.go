package bru

import (
	"bytes"
	"fmt"
	"strings"
)

type Writer struct {
	buffer *bytes.Buffer
}

func NewWriter() *Writer {
	return &Writer{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func (writer *Writer) Bytes() []byte {
	return writer.buffer.Bytes()
}

func (writer *Writer) WriteBlock(name, attr string, blockContentCallback func(*Writer)) {
	if attr != "" {
		fmt.Fprintf(writer.buffer, "%v:%v {\n", name, attr)
	} else {
		fmt.Fprintf(writer.buffer, "%v {\n", name)
	}
	blockContentCallback(writer)
	fmt.Fprintf(writer.buffer, "}\n\n")
}

func (writer *Writer) WriteArrayBlock(name, attr string, arrayValues []string) {
	if len(arrayValues) == 0 {
		return
	}
	if attr != "" {
		fmt.Fprintf(writer.buffer, "%v:%v [\n", name, attr)
	} else {
		fmt.Fprintf(writer.buffer, "%v [\n", name)
	}
	for i, value := range arrayValues {
		sep := ","
		if i == len(arrayValues)-1 {
			sep = ""
		}
		fmt.Fprintf(writer.buffer, "  %v%v\n", value, sep)
	}
	fmt.Fprintf(writer.buffer, "]\n\n")
}

func (writer *Writer) WriteField(key, value string) {
	fmt.Fprintf(writer.buffer, "  %v: %v\n", key, value)
}

func (writer *Writer) WriteEnableableField(key, value string, enabled bool) {
	if !enabled {
		key = "~" + key
	}
	writer.WriteField(key, value)
}

func (writer *Writer) WriteLiteralBlock(name, attr, value string) {
	writer.WriteBlock(name, attr, func(bw *Writer) {
		value = strings.ReplaceAll(strings.TrimSpace(value), "\n", "\n  ")
		fmt.Fprintf(writer.buffer, "  %v\n", value)
	})
}
