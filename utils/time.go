package utils

import "time"

// TimestampToTime将时间戳转化成特定格式的时间字符串
func TimestampToTime(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// TimeToTimestamp将时间字符串转化为时间戳
func TimeToTimestamp(layout string, timestr string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return -1, err
	}
	times, err := time.ParseInLocation(layout, timestr, loc)
	if err != nil {
		return -1, err
	}
	// Time => int64
	trantimestamp := times.Unix()
	return trantimestamp, err
}

// GetCurrentSeds获取当前的时间戳
func GetCurrentSeds() int64 {
	return time.Now().Unix()
}
