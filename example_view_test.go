package gol

import (
	"fmt"
	"time"
)

func ExampleViewText() {
	t, _ := time.Parse("2006/01/02 15:04:05", "2015/01/11 15:04:05")
	record := &Record{
		Time:  t,
		Pid:   34572,
		Level: ERROR,
		Body:  "something went not so"}
	record.Context = []string{"line 1 with some context info", "line 2 with another context info"}
	buf, _ := ViewText(record)

	fmt.Print(buf.String())
	// Output:
	// 34572 2015/01/11 15:04:05 ERROR something went not so
	//       line 1 with some context info
	//       line 2 with another context info
}
