package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestLogPrint(t *testing.T) {
	log.Print("this is a log")
	log.Printf("hello World year: %d", time.Now().Year()) // 格式化输出
	log.Println("hello World")
}

func TestLogPanic(t *testing.T) {
	defer fmt.Println("发生了 panic错误！")
	log.Print("this is a log")
	log.Panic("this is a panic log ")
	fmt.Println("运行结束。。。")
}

func TestLogFatal(t *testing.T) {
	defer fmt.Println("defer。。。")
	log.Print("this is a log")
	log.Fatal("this is a fatal log") // 输出日志，然后退出程序
	fmt.Println("运行结束。。。")
}

func TestLogConfig(t *testing.T) {
	//设置日志格式
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	//增加日期 文件  路径 行号
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("this is a log")
	//前缀配置
	s := log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.SetPrefix("[MyLog] ")
	s = log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.Print("this is a log...")
	//设置输出文件
	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetOutput(f)
	log.Print("this is a file log...")

}

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger = log.New(logFile, "[Mylog]", log.Ldate|log.Ltime|log.Lshortfile)
}

// 自定义logger
func TestLogger(t *testing.T) {
	logger.Print("this is a log")
}
