package date

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println("Current Data: ", dt)

	fmt.Println("format - 01-02-2006", dt.Format("01-02-2006"))
	fmt.Println("format - 01/02/2006", dt.Format("01/02/2006"))
	fmt.Println("format - long month", dt.Format("2006-January-02"))
	fmt.Println("format - 01-02-2006", dt.Format("01-02-2006"))
}
