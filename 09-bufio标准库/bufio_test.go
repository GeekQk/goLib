package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	s := strings.NewReader("ABCDEFG")
	str := strings.NewReader("12345")
	//把文件读入到缓冲区里面
	br := bufio.NewReader(s)
	//从缓冲区中读取文件内容
	b, _ := br.ReadString('\n')
	fmt.Println(b)
	//清空缓冲区域数据 同时把str写入缓冲区中
	br.Reset(str)
	b, _ = br.ReadString('\n')
	fmt.Println(b)
}

func TestReaderStr(t *testing.T) {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	for {
		n, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p[0:n]): %v\n", string(p[0:n]))
		}
	}
}

func TestReaderFromFile(t *testing.T) {
	file, err := os.Open("./go.mod")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// 创建一个带缓冲的读取器
	reader := bufio.NewReader(file)

	// 使用 bufio.Reader 的 ReadString 方法读取文件内容
	line, err := reader.ReadString('\n') // 读取直到遇到换行符
	if err != nil {
		log.Fatalf("Failed to read from file: %v", err)
	}

	fmt.Println("First line of the file:", line)

	// 继续读取剩余内容（如果有的话）
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			if err == bufio.ErrBufferFull {
				fmt.Println("Buffer is full")
				continue
			}
			if err == io.EOF {
				break // 文件结束
			}
			log.Fatalf("Failed to read from file: %v", err)
		}
		fmt.Println(line)
	}
}

func TestReaderPeek(t *testing.T) {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	// 读取一个字节
	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	// 读取下一个字节
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	// 吐出最近一次读取操作读取的最后一个字节
	br.UnreadByte()
	// 读取下一个字节
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}

func TestReaderRune(t *testing.T) {
	s := strings.NewReader("你好，世界！")
	br := bufio.NewReader(s)

	//ReadRune读取一个utf-8编码的unicode码值
	c, size, _ := br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	// ReadRune读取一个utf-8编码的unicode码值
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	// UnreadRune将最近一次读取的unicode码值退回到缓冲区中
	br.UnreadRune()
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	// 读取一行
	c1, status, _ := br.ReadLine()
	fmt.Printf("%c %v %v\n", c1, status, size)
}

func TestReaderLine(t *testing.T) {
	s := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}

func TestReaderReadSlice(t *testing.T) {
	s := strings.NewReader("ABC,DEF,GHI,JKL")
	br := bufio.NewReader(s)

	//ReadSlice读取直到第一次遇到delim字节，返回缓冲里的包含已读取的数据和delim字节的切片
	w, _ := br.ReadSlice(',')
	fmt.Printf("%q %[1]s\n", w, string(w))

	// ReadSlice读取直到第一次遇到delim字节，返回缓冲里的包含已读取的数据和delim字节的切片
	w, _ = br.ReadSlice(',')
	fmt.Printf("%q %[1]s\n", w, string(w))

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q %[1]s\n", w, string(w))
}

func TestReaderEmpty(t *testing.T) {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadString(' ')
	fmt.Printf("%q %[1]v\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q %[1]v\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q %[1]v\n", w)

}

func TestWriterToFile(t *testing.T) {
	// 打开或创建文件
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// 创建一个带缓冲的写入器
	writer := bufio.NewWriter(file)

	// 写入字符串到缓冲写入器
	_, err = writer.WriteString("Hello, World!\n")
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	// 写入其他数据到缓冲写入器
	_, err = writer.Write([]byte("This is some more data."))
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	// 刷新缓冲写入器，将缓冲区中的数据写入到底层的 io.Writer（这里是文件）
	err = writer.Flush()
	if err != nil {
		log.Fatalf("Failed to flush writer: %v", err)
	}

}

func TestWriterToFile2(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123456789")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)
	fmt.Println(c)
}

func TestWriterFun(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	//可用大小
	fmt.Println(bw.Available()) // 4096
	//已写入大小
	fmt.Println(bw.Buffered()) // 0

	bw.WriteString("ABCDEFGHIJKLMN")
	//可用大小
	fmt.Println(bw.Available())
	//已写入大小
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}

func TestReaderWriter(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")
	br := bufio.NewReader(s)

	// 创建一个ReadWriter 可读可写
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')
	fmt.Println(string(p)) // 123
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b) // asdf
}

func TestReaderWriterSplit(t *testing.T) {
	file, err := os.Open("test2.txt")

	// handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
	// read line by line
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}

func TestReaderWriterRead(t *testing.T) {
	//按照空行读取
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
	//按照bytes读取
	s = strings.NewReader("ABC DEF GHI JKL")
	bs = bufio.NewScanner(s)
	bs.Split(bufio.ScanBytes)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
	//按照行读取
	s = strings.NewReader("ABC\nDEF\nGHI\nJKL")
	bs = bufio.NewScanner(s)
	bs.Split(bufio.ScanLines)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}

}
