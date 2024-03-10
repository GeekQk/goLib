package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestTimeInfo(t *testing.T) {
	now := time.Now()                                //获取当前时间
	fmt.Printf("type:%T ,current time:%[1]v\n", now) //时间对象为time.Time类型

	year := now.Year()         //年
	month := now.Month()       //月
	day := now.Day()           //日
	hour := now.Hour()         //小时
	minute := now.Minute()     //分钟
	second := now.Second()     //秒
	week := int(now.Weekday()) //星期
	//02d输出的整数不足两位 用0补足
	fmt.Printf("foramt time:%d-%02d-%02d %02d:%02d:%02d week:%v\n", year, month, day, hour, minute, second, week)
}

func TestUnixTime(t *testing.T) {
	now := time.Now()
	log.Println("时间戳（秒）：", now.Unix())       // 输出：时间戳（秒） ： 1665807442
	log.Println("时间戳（毫秒）：", now.UnixMilli()) // 输出：时间戳（毫秒）： 1665807442207
	log.Println("时间戳（微秒）：", now.UnixMicro()) // 输出：时间戳（微秒）： 1665807442207974
	log.Println("时间戳（纳秒）：", now.UnixNano())  // 输出：时间戳（纳秒）： 1665807442207974500

}

func TestTimeParse(t *testing.T) {
	/*
		goalng的诞生时间->2006-1-02 15:04:05
		Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
		ANSIC       = "Mon Jan _2 15:04:05 2006"
		UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		RFC822      = "02 Jan 06 15:04 MST"
		RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
		RFC3339     = "2006-01-02T15:04:05Z07:00"
		RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		Kitchen     = "3:04PM"
		// Handy time stamps.
		Stamp      = "Jan _2 15:04:05"
		StampMilli = "Jan _2 15:04:05.000"
		StampMicro = "Jan _2 15:04:05.000000"
		StampNano  = "Jan _2 15:04:05.000000000"
		DateTime   = "2006-01-02 15:04:05"
		DateOnly   = "2006-01-02"
		TimeOnly   = "15:04:05"
	*/
	// 字符串时间 解析为时间对象
	t1, _ := time.Parse(time.DateTime, "2022-07-28 18:06:00")
	t2, _ := time.Parse(time.TimeOnly, "2022-07-28 18:06:00")
	t3, _ := time.Parse(time.DateOnly, "2022-07-28 18:06:00")
	fmt.Printf("%T %v\n", t1, t1)
	fmt.Printf("%T %v\n", t2, t2)
	fmt.Printf("%T %v\n", t3, t3)

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2008/01/02 15:14:05", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("format:%T %v\n", timeObj, timeObj)

	// 按照指定时区和指定格式解析字符串时间
	timeObj2, err := time.ParseInLocation("2006/01/02", "2008/01/02", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("format:%T %v\n", timeObj2, timeObj2)

}

// 当前时间格式化为具体时间字符串
func TestTimeFormat(t *testing.T) {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分05秒
	// 24小时制
	fmt.Printf("%T %[1]v\n", now.Format("2006-01-02 15:04:05.000"))
	// 12小时制
	fmt.Printf("%T %[1]v\n", now.Format("2006-01-02 03:04:05"))
	fmt.Printf("%T %[1]v\n", now.Format("2006/01/02 03:04"))
	fmt.Printf("%T %[1]v\n", now.Format("2006/01/02"))

}

func TestUnixToTime(t *testing.T) {
	timeStamp := time.Now().Unix()
	//方法一
	timeObj := time.Unix(timeStamp, 0) //将时间戳转为时间格式
	year := timeObj.Year()             //年
	month := timeObj.Month()           //月
	day := timeObj.Day()               //日
	hour := timeObj.Hour()             //小时
	minute := timeObj.Minute()         //分钟
	second := timeObj.Second()         //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//方法二
	timeObj2 := time.Unix(int64(timeStamp), 0)
	// 格式化time.Time为可读的字符串
	formattedTime := timeObj2.Format("2006-01-02 15:04:05")
	fmt.Printf("%T %[1]v\n", formattedTime)

}

func TestTimeToUnix(t *testing.T) {
	// 方式一
	fmt.Println(time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local).Unix())
	//方式二
	timeObj, err := time.Parse(time.DateTime, "2014-01-07 05:50:04")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeObj.Unix())
}

func TestTimeCalc(t *testing.T) {
	timeObj := time.Now()
	fmt.Printf("start:%T %[1]v\n", timeObj)
	//增加时间
	oneDay := 24 * time.Hour
	newTime := timeObj.Add(oneDay)
	fmt.Printf("add:%T %[1]v\n", newTime.Format(time.DateTime))
	//减少时间
	newTime = timeObj.Add(-oneDay)
	fmt.Printf("sub:%T %[1]v\n", newTime.Format(time.DateTime))
	//使用sub减少时间 返回的是两个时间差
	subTime := timeObj.Sub(newTime)
	fmt.Printf("sub:%T %[1]v %[2]v %[3]v\n", subTime, subTime.Hours(), subTime.Minutes())
	//时间比较 返回bool值
	fmt.Println(timeObj.Before(newTime))
	fmt.Println(timeObj.After(newTime))
	fmt.Println(timeObj.Equal(newTime))
}

func TestTimeer(t *testing.T) {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	time.AfterFunc(time.Second*10, func() {
		fmt.Println("10秒后执行")
	})
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func TestTimeDemo(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(2 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			//Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
			//还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer1)

	wg.Wait()

}
