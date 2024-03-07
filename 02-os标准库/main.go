package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 创建文件
func createFile(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

// 创建文件夹
func mkdir() {
	err := os.Mkdir("ms", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// pe := err.(*os.PathError)
	// fmt.Printf("pe: %v\n", pe)
}

func mkdirAll() {
	err := os.MkdirAll("ms/one/two", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func removeFile() {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func removeDir() {
	err := os.Remove("ms/one")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func removeAllDir() {
	err := os.RemoveAll("ms")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

func chDir() {
	err := os.Chdir("/www/code")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}

func tempDir() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func renameFile() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func renameDir() {
	err := os.Rename("ms", "ms1")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func chmod() {
	err := os.Chmod("./texg.go", 0666)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func chown() {
	err := os.Chown("test23.txt", 10, 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func openClose() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("file name : ", f.Name())
	defer f.Close()

}

func fileStat() {
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

func fileRead() {
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

func fileReadAt() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 5)
	n, err := f.ReadAt(buf, 6)
	fmt.Printf("read content:%s \n", buf[:n])
}

func fileReadDir() {
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

func fileSeek() {
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
func fileWrite() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("hello golang~\n"))
	f.WriteString("hello golang~\n")
}

func fileWriteAt() {
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

func osInfo() {
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

func main() {
	/*
		// 创建文件 默认使用666权限
		createFile("test.txt")
		//创建单级文件目录
		mkdir()
		//多级文件目录
		mkdirAll()
		//删除文件
		removeFile()
		//删除文件夹
		removeDir()
		//删除多级文件夹
		removeAllDir()
		//获取当前工作目录
		getWd()
		//修改工作目录
		// chDir()
		//获取临时目录地址
		tempDir()
		//文件改名
		renameFile()
		//文件夹改名
		renameDir()
		//文件修改权限
		chmod()
		//修改文件所有者权限
		chown()
	*/

	/*
		//文件打开和读取
		// O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
		// O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
		// O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
		// O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
		// O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
		// O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		// O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
		// O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件

		// 打开文件
		openClose()
		//文件状态
		fileStat()
		//文件读取
		fileRead()
		//文件从哪里开始读
		// fileReadAt()
		//打开目录
		fileReadDir()
		//文件内容偏移
		fileSeek()
		//文件写
		fileWrite()
		//文件偏移位置写,但是会把重合位置字符替换掉
		fileWriteAt()
	*/
	osInfo()

}
