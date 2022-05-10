package utils

import (
	"github.com/kvdatabase/kvdb/storage"
	"os"
	"testing"
)

func TestExist1(t *testing.T) {
	t.Log(os.TempDir() + "ssds")

	exist := Exist(os.TempDir() + "ssds")
	t.Log(exist)

	if err := os.MkdirAll(os.TempDir()+"abcd", storage.FilePerm); err != nil {
		t.Error(err)
	}
}

func TestExist2(t *testing.T) {
	//目录是否存在
	path := "/tmp/kvdb"

	t.Log(Exist(path))

	//文件是否存在
	t.Log(Exist(path + "/000w000000.data"))

	t.Log(os.TempDir())
}

func TestCopyFile(t *testing.T) {
	//src := "/Users/kvdatabase/resources/books/skiplist cookbook.pdf"
	//dst := "/Users/kvdatabase/resources/books/skiplist cookbook-bak.pdf"
	//
	//err := CopyFile(src, dst)
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestCopyDir(t *testing.T) {
	//src := "/Users/kvdatabase/resources/books-new"
	//dst := "/Users/kvdatabase/resources/books-new2"
	//
	//err := CopyDir(src, dst)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
