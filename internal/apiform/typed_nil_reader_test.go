package apiform

import (
	"bytes"
	"io"
	"mime/multipart"
	"testing"
)

type panicReader struct{}

func (*panicReader) Read([]byte) (int, error) {
	panic("Read called on nil receiver")
}

func TestMarshalTypedNilReaderAsEmptyField(t *testing.T) {
	var reader *panicReader
	var output bytes.Buffer
	writer := multipart.NewWriter(&output)

	if err := Marshal(map[string]any{"file": reader}, writer); err != nil {
		t.Fatalf("marshal typed nil reader: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	part, err := multipart.NewReader(&output, writer.Boundary()).NextPart()
	if err != nil {
		t.Fatalf("read multipart part: %v", err)
	}
	contents, err := io.ReadAll(part)
	if err != nil {
		t.Fatalf("read multipart contents: %v", err)
	}
	if string(contents) != "" {
		t.Fatalf("typed nil reader contents = %q, want empty field", contents)
	}
}
