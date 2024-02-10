package unixTimeConverter

import (
	"time"
)

// Takes in Unix time code and spis out readable string.
func ConvertUnixTimeToReadable(epochTime int64) string {
	t := time.Unix(epochTime, 0)
	return t.Format("2006-01-02 15:04:05")
}
