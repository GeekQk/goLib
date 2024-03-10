package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

type User struct {
	Id int64
}

func TestPrint(t *testing.T) {
	fmt.Println("---------打印结构体----------")
	user := &User{Id: 1}
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	fmt.Printf("%%\n")
	fmt.Println("---------打印布尔类型----------")
	fmt.Printf("%t\n", true)
	fmt.Printf("%t\n", false)
	fmt.Println("---------打印数值类型----------")
	n := 180
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n", n)
	a := 96
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", 0x4E2D)
	fmt.Println("---------打印浮点数----------")
	f := 18.54
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
	fmt.Println("---------打印字符串----------")
	s := "我是字符串"
	b := []byte{65, 66, 67}
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
	fmt.Println("---------打印指针----------")
	fmt.Printf("%p\n", &s)
	fmt.Println("---------占位符打印----------")
	s = "我是字符串"
	fmt.Printf("%  d\n", 10)
	fmt.Printf("%s\n", s)
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)
	fmt.Printf("%10.2f\n", 10.14)
	fmt.Printf("%-10.2f\n", 10.14)
	fmt.Printf("%010s\n", s)
}

func TestFprint(t *testing.T) {
	fmt.Println("-------Fprint 把内容输出到io输出流中----------")
	fmt.Fprint(os.Stdout, "标准输出打印\n")
	fmt.Fprintln(os.Stdout, "标准输出打印")
	fmt.Fprintf(os.Stdout, "标准输出打印:%v\n", 10)

	fmt.Println("-------Fprint 把内容输出到文件中----------")
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "文件输出打印:%v\n", 10)
	fmt.Fprintln(file, "文件输出打印")
}


func TestSprint(*testing.T) {
	fmt.Println("-------Sprint/f/ln字符串拼接-------------------")
	fmt.Println(fmt.Sprint("10\n"))
	fmt.Println(fmt.Sprintln("10"))
	fmt.Println(fmt.Sprintf("10+10=%d", 20))
	host := "127.0.0.1"
	port := 80
	addr := fmt.Sprintf("%s:%d", host, port)
	addr2 := host + ":" + fmt.Sprint(port)
	fmt.Println(addr, addr2)
}


func TestError(*testing.T) {
	fmt.Println("-------使用Errorf生成错误----------")
	err1 := errors.New("用户名格式不正确")
	err2 := fmt.Errorf("用户名格式不正确：%s", "@#￥哈哈")
	if err1 != nil {
		// panic(err)
		fmt.Println(err1, err2)
		fmt.Printf("%T %[1]v", err1)
	}

}

func TestScan(t *testing.T) {
	fmt.Println("------------扫描Scan----------------")
	var (
		name    string
		age     int
		married bool
	)
	//从标准输入中读取数据
	// fmt.Scan(&name, &age, &married)
	// fmt.Scanln(&name, &age, &married)  //一样的
	// fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married) // 1:张三 2:18 3:true 按照这个顺序进行接收数据
	//从文件、字符串中读取数据
	reader := strings.NewReader("1:zhangsan 2:18 3:true")
	fmt.Fscanf(reader, "1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("name:%s,age:%d,married:%t", name, age, married)
}
