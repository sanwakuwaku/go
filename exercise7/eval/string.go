package eval

import (
	"strconv"
	"strings"
)

//!+exercise7.13

func (v Var) String() string {
	return "" + string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', 3, 64)
}

func (u unary) String() string {
	return "(" + string(u.op) + u.x.String() + ")"
}

func (b binary) String() string {
	return "(" + b.x.String() + string(b.op) + b.y.String() + ")"
}

func (c call) String() string {
	str := c.fn + "("
	for _, expr := range c.args {
		str = strings.Join([]string{str, expr.String(), ","}, "")
	}
	str = str[:len(str)-1]
	str += ")"

	return str
}

//!-exercise7.13
