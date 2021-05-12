package dates_utils

import "time"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(DEFAULT_DATE_FORMAT)
}
