package main

import (
	"fmt"
	"time"
)

const (
	dateF = "2006-01-02"
	timeF = "15:04:05"
	datetimeF = "2006-01-02 15:04:05"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now)		// 2020-07-23 10:56:00.100285 +0800 CST m=+0.000081686
	// 格式化输出
	fmt.Println(time.Now().Format(datetimeF))	// 2020-07-23 10:56:00

	// 日期操作
	thisDateS := "2020-07-22"
	// 解析特定时间字符串，必须填入格式
	thisDate,_ := time.Parse(dateF, thisDateS)
	fmt.Println(thisDate.Format(dateF))			// 2020-07-22
	// 日期转化为时间，代表当前的 0 点
	fmt.Println(thisDate.Format(datetimeF))		// 2020-07-22 00:00:00

	// 日期加减 , + 为日期向后推，- 为日期向前移
	// 7天前
	yesDate := thisDate.AddDate(0,0,-7)				// 2020-07-15
	fmt.Println(yesDate.Format(dateF))

	// 上个月的今天
	thisDateOfLastMouth := thisDate.AddDate(0,-1,0)	// 2020-06-22
	fmt.Println(thisDateOfLastMouth.Format(dateF))
	// 也可以通过 -30 天得到（不过每个月的天数都一样，需要根据实际的月份进行减）
	thisDateOfLastMouth = thisDate.AddDate(0,0,-30)	// 2020-06-22
	fmt.Println(thisDateOfLastMouth.Format(dateF))

	// 明年的今天
	thisDateOfNextYear := thisDate.AddDate(1,0,0)	// 2021-07-22
	fmt.Println(thisDateOfNextYear.Format(dateF))
	thisDateOfNextYear = thisDate.AddDate(0,0,365)  // 2021-07-22
	fmt.Println(thisDateOfNextYear.Format(dateF))

	// 比较时间前后
	fmt.Println(yesDate.Before(thisDate))			// 是否在之前
	fmt.Println(yesDate.Equal(thisDate))			// 是否相等
	fmt.Println(yesDate.After(thisDate))			// 是否在之后

	// 时间操作
	thisTimeS := "2020-07-22 11:00:00"
	thisTime, _ := time.Parse(datetimeF, thisTimeS)
	// 只打印时间
	fmt.Println(thisTime.Format(timeF))

	// 10秒钟前
	t1:= thisTime.Add(time.Second * -10)
	fmt.Println(t1.Format(timeF))
	// 20分钟后
	t2 := thisTime.Add(time.Minute * 20)
	fmt.Println(t2.Format(timeF))
	// 28小时前
	t3 := thisTime.Add(time.Hour * -28)
	fmt.Println(t3.Format(timeF))

	// 计算时间差并显示
	begin,_ := time.Parse(datetimeF, "2020-07-22 11:00:00")
	end,_ := time.Parse(datetimeF, "2020-07-13 11:00:00")
	subT := begin.Sub(end)
	fmt.Println(subT.Hours(), "小时")
	fmt.Println(subT.Hours()/24, "天")

	// 10位时间戳（精确到s）
	fmt.Println(time.Now().Unix())
	// 13位时间戳（精确到ms）, 后3位为ms
	fmt.Println(time.Now().UnixNano()/1000000)
	// 19位时间戳（精确到ns）
	fmt.Println(time.Now().UnixNano())
	// UTC 时间
	fmt.Println(time.Now().UTC())

}
