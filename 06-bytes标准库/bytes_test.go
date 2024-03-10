package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestBytesUpperAndLower(t *testing.T) {
	var b = []byte("seafood") //强制类型转换

	a := bytes.ToUpper(b)
	fmt.Println("全部大写:", string(a), "bytes:", a)
	b = bytes.ToLower(a)
	fmt.Println("全部小写:", string(b), "bytes:", b)
}

func TestBytesEquip(t *testing.T) {
	a := "Abc"
	b := "ABc"
	//常规比较
	ret := bytes.Compare([]byte(a), []byte(b))
	fmt.Printf("compare result: %d\n", ret)
	//不忽略大小写
	bools := bytes.Equal([]byte(a), []byte(b))
	fmt.Printf("compare result: %v\n", bools)
	//忽略大小写
	bools = bytes.EqualFold([]byte(a), []byte(b))
	fmt.Printf("compare result: %v\n", bools)

}

func TestBytesTrim(t *testing.T) {
	a := []byte("  Abc")
	// b := []byte("ABc  ")
	fmt.Printf("trim: %v,value:%[1]c\n", bytes.Trim(a, " "))
	fmt.Printf("ltrim: %v,value:%[1]c\n", bytes.TrimLeft(a, " "))
	fmt.Printf("ltrim: %v,value:%[1]c\n", bytes.TrimRight(a, " "))
	fmt.Printf("ltrim: %v,value:%[1]c\n", bytes.TrimPrefix(a, []byte("  A")))

	b := []byte("!!Hello World")
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!!"), r) //判断r字符是否包含在"!
	}
	fmt.Printf("去掉两边: %q\n", bytes.TrimFunc(b, f)) //去掉两边满足函数的字符
}

func TestBytesSplit(t *testing.T) {
	b := []byte("Hello,World,Go")
	//按照指定字符分割
	fmt.Printf("%T %[1]c\n", bytes.Split(b, []byte(",")))
	//以连续空白为分隔符将 s 切分成多个子串
	a := []byte(" Hello World Go")
	fmt.Printf("%T %[1]c\n", bytes.Fields(a))
	//子串拼接
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%T %[1]c\n", bytes.Join(s, []byte(", ")))
	//字符重复
	fmt.Printf("%T %[1]c\n", bytes.Repeat([]byte("a"), 3))
}

func TestBytesExicute(t *testing.T) {
	b := []byte("Hello,World,Go")
	//判断是否包含指定字符
	fmt.Printf("%v\n", bytes.Contains(b, []byte(",")))
	//判断是否以指定字符开头
	fmt.Printf("%v\n", bytes.HasPrefix(b, []byte("H")))
	//判断是否以指定字符结尾
	fmt.Printf("%v\n", bytes.HasSuffix(b, []byte("o")))
	//打印指定字符出现的次数
	fmt.Println(bytes.Count(b, []byte("o"))) //1
	//返回指定字符第一次出现的位置
	fmt.Println(bytes.Index(b, []byte("o"))) //7
	//返回指定字符最后一次出现的位置
	fmt.Println(bytes.LastIndex(b, []byte("o"))) //10
	//返回指定字符第一次出现的位置和最后一次出现的位置
	fmt.Println(bytes.IndexByte(b, 'o'))     //7
	fmt.Println(bytes.LastIndexByte(b, 'o')) //10
}

func TestBytesReplace(t *testing.T) {
	s := []byte("Hello,World,Go")

	fmt.Println(string(bytes.Replace(s, []byte("o"), []byte("ee"), 0)))  //hello,world
	fmt.Println(string(bytes.Replace(s, []byte("o"), []byte("ee"), 1)))  //hello,world
	fmt.Println(string(bytes.Replace(s, []byte("o"), []byte("ee"), 2)))  //hello,world
	fmt.Println(string(bytes.Replace(s, []byte("o"), []byte("ee"), -1))) //hello,world

	//将 s 转换为 []rune 类型返回
	s1 := []byte("你好世界")
	r := bytes.Runes(s1)
	fmt.Println("转换前字符串的长度：", len(s1), s1) //12
	fmt.Println("转换后字符串的长度：", len(r), r)   //4

}

func TestBytesBuffer(t *testing.T) {
	rd := bytes.NewBufferString("Hello World!")
	buf := make([]byte, 6)
	// 获取数据切片
	b := rd.Bytes()
	// 读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd.String())
	fmt.Printf("%s\n\n", b)

	// 写入一部分数据，看看切片有没有变化
	rd.Write([]byte("abcdefg"))
	fmt.Printf("%s\n", rd.String())
	fmt.Printf("%s\n\n", b)

	// 再读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd.String())
	fmt.Printf("%s\n", b)
}

// 从buff中读取数据
func TestBytesBuffer2(t *testing.T) {
	data := "123456789"
	//通过[]byte创建Reader
	re := bytes.NewReader([]byte(data))
	//返回未读取部分的长度
	fmt.Println("re len : ", re.Len())
	//返回底层数据总长度
	fmt.Println("re size : ", re.Size())

	fmt.Println("---------------")

	buf := make([]byte, 2)
	for {
		//读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

/*
我们首先打开了一个名为 source.txt 的源文件
并读取了它的全部内容到字节切片 sourceBytes 中
我们创建了一个 bytes.Reader 实例 sourceReader，它将以 sourceBytes 作为数据源。
接着，我们创建了一个名为 destination.txt 的目标文件，并使用 io.Copy 函数将 sourceReader 中的内容复制到目标文件中。最后，我们调用了 destFile.Sync() 来确保数据被写入到磁盘
*/
func TestBytesBuffer3(t *testing.T) {
	// 打开源文件
	sourceFile, err := os.Open("source.txt")
	if err != nil {
		log.Fatalf("Failed to open source file: %v", err)
	}
	defer sourceFile.Close()

	// 读取源文件内容到字节切片
	sourceBytes, err := io.ReadAll(sourceFile)
	if err != nil {
		log.Fatalf("Failed to read source file: %v", err)
	}

	// 创建 bytes.Reader
	sourceReader := bytes.NewReader(sourceBytes)

	// 打开或创建目标文件
	destFile, err := os.Create("destination.txt")
	if err != nil {
		log.Fatalf("Failed to create destination file: %v", err)
	}
	defer destFile.Close()

	// 使用 io.Copy 将 bytes.Reader 的内容写入到目标文件中
	if _, err := io.Copy(destFile, sourceReader); err != nil {
		log.Fatalf("Failed to write to destination file: %v", err)
	}

	// 确保数据被写入到磁盘
	if err := destFile.Sync(); err != nil {
		log.Fatalf("Failed to sync destination file: %v", err)
	}

	// 打印完成信息
	log.Println("File copied successfully!")
}
