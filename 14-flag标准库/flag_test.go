package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestArgs(t *testing.T) {
	// 测试命令行参数
	if len(os.Args) > 0 {
		for i, v := range os.Args {
			fmt.Println(i, v)
		}
	}
}

func TestFlag(t *testing.T) {
	// flag.Type() 的使用
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	flag.Parse()

	fmt.Println(*name, *age, *married, *delay)
}

func TestFlag2(t *testing.T) {
	var name string
	var age uint
	var married bool
	var d time.Duration

	//增加默认值
	flag.StringVar(&name, "name", "王五", "姓名")
	flag.UintVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "m", false, "婚否")
	flag.DurationVar(&d, "duration", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name, age, married, d)
}

func TestFlag3(t *testing.T) {
	var name string
	var age uint
	var married bool
	var d time.Duration

	flag.StringVar(&name, "name", "王五", "姓名")
	flag.UintVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "m", false, "婚否")
	flag.DurationVar(&d, "duration", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name, age, married, d) // lisi 35 true 1h15m36s

	fmt.Println(flag.Args())  // [abc true 123]
	fmt.Println(flag.NArg())  // 3
	fmt.Println(flag.NFlag()) //  4
}
