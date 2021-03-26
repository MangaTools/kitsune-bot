package router

import (
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`(".+")|(\S+)`)

type RouterContext struct {
	Args []string
}

func NewRouterContext(input string) *RouterContext {
	result := r.FindAllString(input, -1)
	for i, val := range result {
		if strings.HasPrefix(val, "\"") {
			result[i] = strings.Replace(result[i], "\"", "", -1)
		}
	}
	return &RouterContext{
		Args: result,
	}
}
