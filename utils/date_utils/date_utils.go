package date_utils

import "time"

const (
	apiDateTime = "2006-01-02T15:4:5Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateTime)
}
