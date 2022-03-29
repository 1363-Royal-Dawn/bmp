package bmp

import (
	"bytes"
	"testing"
)

func TestIconHeader(t *testing.T) {
	var buf bytes.Buffer

	ih := NewIconDir(1)
	if err := ih.Save(&buf); err != nil {
		t.Error(err)
	}

	expected := []byte{0, 0, 1, 0, 1, 0}
	got := buf.Bytes()
	if bytes.Compare(got, expected) != 0 {
		t.Errorf("expected %v. got %v", expected, got)
	}
	t.Log(buf.Bytes())

}

func TestIconEntry_Save(t *testing.T) {
	var buf bytes.Buffer

	ie := NewIconDirEntry(32, 32, 16, 0, 1, 74, 100)
	if err := ie.Save(&buf); err != nil {
		t.Error(err)
	}

	t.Log(buf.Bytes())
}
