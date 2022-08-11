package utils

import "time"

func ParseTimeStampMs(timestamp int64) time.Time {
	return time.Unix(timestamp/1000, timestamp%1000)
}

func ParseTimeStampS(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
