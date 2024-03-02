package time

import (
	"fmt"
	"time"
)

func RunAllTimeFunc() {
	fmt.Println(GetTimeNow())
}

func GetTimeNow() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
	// return time.Now() // returns 2024-03-02 19:04:05.083112408 +0100 CET m=+0.000053101
}
