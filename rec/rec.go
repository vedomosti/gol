package rec

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

var (
	colorReset = ansi.ColorCode("reset")

	colorReqType   = ansi.ColorCode("blue+bh")
	colorReqMethod = ansi.ColorCode("white+h")
	colorReqIp     = ansi.ColorCode("magenta")
	colorReqUrl    = ansi.ColorCode("yellow")

	colorResStatus = ansi.ColorCode("white+b")
	colorDuration  = ansi.ColorCode("green")
)

// Req struct for http request
type Req struct {
	Ip     string
	Method string
	Url    string
}

func (req *Req) Map() map[string]string {
	return map[string]string{}
}

func (req *Req) Slice() []string {
	return []string{
		colorReqType + "HTTP Request" + colorReset,
		colorReqIp + req.Ip + colorReset,
		colorReqMethod + req.Method + colorReset,
		colorReqUrl + req.Url + colorReset,
	}
}

func (req *Req) String() string {
	return strings.Join(req.Slice(), " ")
}

// Req struct for http response
type Res struct {
	Status   int
	Url      string
	Duration time.Duration
}

func (res *Res) Map() {}

func (res *Res) Slice() []string {
	return []string{
		colorReqType + "HTTP Response" + colorReset,
		colorResStatus + strconv.Itoa(res.Status) + colorReset,
		colorReqUrl + res.Url + colorReset,
		colorDuration + fmt.Sprintf("%v", res.Duration) + colorReset,
	}
}

func (res *Res) String() string {
	return strings.Join(res.Slice(), " ")
}

// Sql struct for sql query
type Sql struct {
	Query    string
	Params   interface{}
	Duration time.Duration
}

// Msg type for raw message
type Msg string
