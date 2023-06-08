package utils

import "time"

func currentTime() time.Time {
	return time.Now()
}

// GetNowStr 获取时间戳
func GetNowStr() string {
	formattedTime := currentTime().Format("2006-01-02 15:04:05")
	//fmt.Println(formattedTime)
	return formattedTime
}

// 获取11位unix时间戳
func GetTime() int64 {

	// 获取11位的Unix时间戳（秒级）
	unixTimestamp := currentTime().Unix()

	//fmt.Println(unixTimestamp)
	return unixTimestamp
	//return strconv.FormatInt(unixTimestamp, 10)
}

// 获取上一个小时
func GetPrevHourId() string {
	previousHour := currentTime().Add(-time.Hour)

	//fmt.Println("当前时间：", currentTime)
	//fmt.Println("上一个小时的时间：", previousHour)
	return previousHour.Format("2006010215")
}

// 获取前x小时
func GetAgoHourId(n int) string {

	// 计算 n 小时前的时间
	hoursAgo := currentTime().Add(-time.Duration(n) * time.Hour)

	//fmt.Println(n, "小时前的时间是:", hoursAgo)
	return hoursAgo.Format("2006010215")
}

// 获取当前小时证书
func GetNowHourId() string {
	//fmt.Println("当前时间：", currentTime)
	//fmt.Println("上一个小时的时间：", previousHour)
	return currentTime().Format("2006010215")
}

// 获取当前毫秒
func GetMilliSecond() int64 {
	milliseconds := currentTime().UnixNano() / int64(time.Millisecond)

	//fmt.Println("当前毫秒时间：", milliseconds)
	return milliseconds
}

// 获取当前微秒
func GetMicroseconds() int64 {
	microseconds := currentTime().UnixNano() / int64(time.Microsecond)

	//fmt.Println("当前微秒时间：", microseconds)
	return microseconds
}
