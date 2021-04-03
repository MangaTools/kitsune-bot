package router

import (
	"github.com/ShaDream/kitsune-bot/models"
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`(".+")|(\S+)`)

type RouterContext struct {
	Args       []string
	StartText  string
	UserAccess models.RoleAccess
}

func NewRouterContext(input string) *RouterContext {
	trimmedSpaces := strings.TrimSpace(input)
	result := r.FindAllString(input, -1)
	for i, val := range result {
		if strings.HasPrefix(val, "\"") {
			result[i] = strings.Replace(result[i], "\"", "", -1)
		}
	}
	return &RouterContext{
		Args:       result,
		StartText:  trimmedSpaces,
		UserAccess: models.Reader,
	}
}
