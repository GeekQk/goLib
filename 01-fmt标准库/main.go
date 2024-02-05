package main

import (
	"fmt"
)

type User struct {
	Id int64
}

func main() {
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
	fmt.Println("---------占位符打印----------")
	s = "我是字符串"
	fmt.Printf("%  d\n", 10)
	fmt.Printf("%s\n", s)
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)
	fmt.Printf("%10.2f\n", 10.14)
	fmt.Printf("%-10.2f\n", 10.14)
	fmt.Printf("%010s\n", s)
	fmt.Println("-------使用Errorf生成panic错误----------")
	err := fmt.Errorf("用户名格式不正确：%s", "@#￥哈哈")
	if err != nil {
		panic(err)
	}

}
