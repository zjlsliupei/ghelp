package ghelp

import (
	"testing"
)

func TestCopyFile(t *testing.T) {
	_, err := CopyFile("d:\\CfTodo.tar.gz", "e:\\CfTodo111111.tar.gz")
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestIsFile(t *testing.T) {
	if !IsFile(".gitignore") {
		t.Error(".gitignore is file")
	}
}

func TestIsDir(t *testing.T) {
	if !IsDir("test_dir") {
		t.Error("test_dir is dir")
	}
}
