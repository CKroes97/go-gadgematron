package timecheck

import (
	"fmt"
	"time"
)

func Run() {
	now := time.Now()
	fmt.Println("Current time:", now.Format(time.RFC1123))
}
