package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMathFunc(t *testing.T) {
	fmt.Printf("Float64的最大值: %.f\n", math.MaxFloat64)
	fmt.Printf("Float64最小值: %.f\n", math.SmallestNonzeroFloat64)
	fmt.Printf("Float32最大值: %.f\n", math.MaxFloat32)
	fmt.Printf("Float32最小值: %.f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("Int8最大值: %d\n", math.MaxInt8)
	fmt.Printf("Int8最小值: %d\n", math.MinInt8)
	fmt.Printf("Uint8最大值: %d\n", math.MaxUint8)
	fmt.Printf("Int16最大值: %d\n", math.MaxInt16)
	fmt.Printf("Int16最小值: %d\n", math.MinInt16)
	fmt.Printf("Uint16最大值: %d\n", math.MaxUint16)
	fmt.Printf("Int32最大值: %d\n", math.MaxInt32)
	fmt.Printf("Int32最小值: %d\n", math.MinInt32)
	fmt.Printf("Uint32最大值: %d\n", math.MaxUint32)
	fmt.Printf("Int64最大值: %d\n", math.MaxInt64)
	fmt.Printf("Int64最小值: %d\n", math.MinInt64)
	fmt.Printf("圆周率默认值: %v\n", math.Pi)
}

func TestMathCheck(t *testing.T) {
	//检测是否是数值 不是数值返回false
	fmt.Println(math.IsNaN(12321.321321)) //false
	fmt.Println(math.IsNaN(12321))        //false
	fmt.Println(math.IsNaN(-111))         //false
	//向上取整
	fmt.Println(math.Ceil(1.13456))     //2
	fmt.Println(math.Ceil(-1.98765432)) //-1
	//向下取整
	fmt.Println(math.Floor(2.9999)) //2
	fmt.Println(math.Floor(-1.98765432))
	//返回整数部分
	fmt.Println(math.Trunc(2.9999)) //2
	//绝对值
	fmt.Println(math.Abs(2.999312323132141665374)) //2.9993123231321417
	fmt.Println(math.Abs(2.999312323132141465374)) //2.9993123231321412
	//最大值、最小值
	fmt.Println(math.Max(1000, 200)) //1000
	fmt.Println(math.Min(1000, 200)) //200
	//函数返回x-y和0中的最大值
	fmt.Println(math.Dim(1000, 2000)) //0
	fmt.Println(math.Dim(1000, 200))  //800

	//取余运算
	fmt.Println(math.Mod(123, 0))  //NaN
	fmt.Println(math.Mod(123, 10)) //3

	//平方根
	fmt.Println(math.Sqrt(144)) //12
	//立方根
	fmt.Println(math.Cbrt(1728)) //12
	// 求幂，x的y次方
	fmt.Println(math.Pow(2, 3)) //8

}

func TestMathRandom1(t *testing.T) {
	//随机数
	// 直接调用rand的方法生成伪随机int值
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Intn(5))
}

func TestMathRandom2(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano()) // 使用当前的纳秒生成一个随机源，也就是随机种子。NewSource()方法就等价于前面的rand.Seed()方法,都是用来设置随机种子。，这两种方法本质上没有区别。
	ran := rand.New(source)                         // 生成一个rand
	fmt.Println(ran.Int())                          // 获取随机数
}
