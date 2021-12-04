package main

import (
	"fmt"
	"time"
)

func main() {

	//分别使用 time.Now 的 Unix 和 UnixNano， 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	mills := now.UnixMilli()
	println(secs)
	println(nanos)
	println(mills)

	//你也可以将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
	fmt.Println(time.Unix(0, mills*1000000))
}
