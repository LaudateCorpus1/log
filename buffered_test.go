package log

import (
	"bytes"
	"testing"
)

func TestBufferedWriter(t *testing.T) {

	buffer := &bytes.Buffer{}
	buffered := NewBufferedWriter(buffer)
	buffered.Start()

	_, err := buffered.Write([]byte("hello"))
	if err != nil {
		t.Error("expected no err got", err)
	}
	buffered.Close()
	if str := buffer.String(); str != "hello" {
		t.Error("expected hello got:", str)
	}
}

func TestBufferedOverflow(t *testing.T) {
	buffer := &bytes.Buffer{}
	buffered := NewBufferedWriter(buffer)

	for i := 0; i < BufferSize; i++ {
		_, err := buffered.Write([]byte("hello"))
		if err != nil {
			t.Error("expected no err got", err)
		}
	}

	_, err := buffered.Write([]byte("hello"))
	if err != BufferFull {
		t.Error("expected BufferFull")
	}
}
