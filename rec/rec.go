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

	colorSqlType   = ansi.ColorCode("blue+bh")
	colorSqlQuery  = ansi.ColorCode("white+h")
	colorSqlParams = ansi.ColorCode("yellow")

	colorMsg = ansi.ColorCode("white+h")
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

func (sql *Sql) Slice() []string {
	return []string{
		colorSqlType + "SQL" + colorReset,
		colorSqlQuery + sql.Query + colorReset,
		colorSqlParams + fmt.Sprintf("%+v", sql.Params) + colorReset,
		colorDuration + fmt.Sprintf("%v", sql.Duration) + colorReset,
	}

}

func (sql *Sql) String() string {
	return strings.Join(sql.Slice(), " ")
}

// Msg type for raw message
type Msg string

func (msg *Msg) String() string {
	return colorMsg + string(*msg) + colorReset
}
