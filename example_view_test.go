package gol

import (
	"fmt"
	"time"

	"github.com/kavkaz/gore"
)

func ExampleViewText() {
	t, _ := time.Parse("2006/01/02 15:04:05", "2015/01/11 15:04:05")
	record := &Record{
		Time:   t,
		Pid:    34572,
		Level:  ERROR,
		Caller: NewCaller(0),
		Body:   "something went not so",
	}
	record.Context = []string{"line 1 with some context info", "line 2 with another context info"}
	buf, _ := ViewText(record)

	fmt.Print(buf.String())
	// Output:
	// 34572 2015/01/11 15:04:05 ERROR example_view_test.go:16 gol.ExampleViewText something went not so
	//       line 1 with some context info
	//       line 2 with another context info
}

func ExampleViewTexti_Gore() {
	t, _ := time.Parse("2006/01/02 15:04:05", "2015/01/11 15:04:05")
	record := &Record{
		Time:   t,
		Pid:    34572,
		Level:  ERROR,
		Caller: NewCaller(0),
	}

	err := gore.New("Some error")
	gore.Append(err, "step one")
	gore.Append(err, "step two")

	record.Body, record.Context = errorInfo(err)

	buf, _ := ViewText(record)

	fmt.Print(buf.String())
	// Output:
	// 34572 2015/01/11 15:04:05 ERROR example_view_test.go:35 gol.ExampleViewTexti_Gore [example_view_test.go:38 gol.ExampleViewTexti_Gore] Some error
	//       #0 [example_view_test.go:39 gol.ExampleViewTexti_Gore] step one
	//       #1 [example_view_test.go:40 gol.ExampleViewTexti_Gore] step two
}
