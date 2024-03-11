package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestReadAll(t *testing.T) {
	f, err := os.Open("./go.mod")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(string(b))
}

func TestReaddir(t *testing.T) {
	b, err := os.ReadDir("img")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("err: %v\n", b)
}

func TestReadWriteFile(t *testing.T) {
	b, err := os.ReadFile("./go.mod")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(b)
	os.WriteFile("a.txt", []byte("hello world"), 0664)
}

func TestTempDir(t *testing.T) {
	//获取临时目录
	fmt.Println(os.TempDir())
	//路径拼接
	fmt.Println(filepath.Join(os.TempDir(), "a.txt"))
	//创建临时目录
	b, err := os.MkdirTemp(os.TempDir(), "a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(b)
}
