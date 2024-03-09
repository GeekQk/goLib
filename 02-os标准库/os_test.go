package main

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

// 创建文件
func TestCreateFile(t *testing.T) {
	fileName := "test.txt"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

// 创建文件夹
func TestMkdir(t *testing.T) {
	err := os.Mkdir("ms", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// pe := err.(*os.PathError)
	// fmt.Printf("pe: %v\n", pe)
}

func TestMkdirAll(t *testing.T) {
	err := os.MkdirAll("ms/one/two", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestRemoveFile(t *testing.T) {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func TestRemoveDir(t *testing.T) {
	err := os.Remove("ms/one")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestRemoveAllDir(t *testing.T) {
	err := os.RemoveAll("ms")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestGetWd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

func TestChDir(t *testing.T) {
	err := os.Chdir("/www/code")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}

func TestTempDir(t *testing.T) {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func TestRenameFile(t *testing.T) {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func TestRenameDir(t *testing.T) {
	err := os.Rename("ms", "ms1")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestChmod(t *testing.T) {
	err := os.Chmod("./texg.go", 0666)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestChown(t *testing.T) {
	err := os.Chown("test23.txt", 10, 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestOpenClose(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("file name : ", f.Name())
	defer f.Close()

}

func TestFileStat(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("file info : %#v", fileInfo)
}

func TestFileRead(t *testing.T) {
	f, err := os.OpenFile("./test.txgt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var body []byte
	for {
		buf := make([]byte, 4)
		n, err := f.Read(buf)
		if err == io.EOF {
			//读完了
			break
		}
		fmt.Printf("读到的位置:%d \n", n)
		body = append(body, buf[:n]...)
	}
	fmt.Printf("内容：%s \n", body)
}

func TestFileReadAt(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 5)
	n, err := f.ReadAt(buf, 6)
	fmt.Printf("read content:%s \n", buf[:n])
}

func TestFileReadDir(t *testing.T) {
	f, err := os.Open("/www/")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//-1代表所有
	dirs, err := f.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, v := range dirs {
		fmt.Println("is dir:", v.IsDir())
		fmt.Println("dir name :", v.Name())
	}
}

func TestFileSeek(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("读取内容:%s\n", buf[:n])
}
func TestFileWrite(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("hello golang~\n"))
	f.WriteString("hello golang~\n")
}

func TestFileWriteAt(t *testing.T) {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteAt([]byte(" insert content~\n"), 5)
	if err == nil {
		fmt.Println("insert succ")
	}
}

func TestOsInfo(t *testing.T) {
	// 获得当前正在运行的进程id
	fmt.Println("---------")
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 父id
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())

	// 设置新进程的属性
	attr := &os.ProcAttr{
		// files指定新进程继承的活动文件对象
		// 前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		// 新进程的环境变量
		Env: os.Environ(),
	}

	// 开始一个新进程
	p, err := os.StartProcess("/usr/bin/env", []string{"c:\\windows\\system32\\notepad.exe", "d:\\test.txt"}, attr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println("进程ID:", p.Pid)

	// 通过进程ID查找进程
	p2, _ := os.FindProcess(p.Pid)
	fmt.Println(p2)

	// 等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		// 向p进程发出退出信号
		p.Signal(os.Kill)
	})

	// 等待进程p的退出，返回进程状态
	ps, _ := p.Wait()
	fmt.Println(ps.String())
}
