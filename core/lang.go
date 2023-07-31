package core

import (
	"admin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type lang struct {
	options []string
}

var Lang = &lang{
	options: []string{"zh_CN", "zh_TW", "en"},
}

// GetMassage 获取消息
func (l *lang) GetMassage(c *gin.Context, key string, args ...[]interface{}) string {
	message := key
	locale := l.options[0]
	language := c.GetHeader("Accept-Language")
	if util.ArrUtil.InArray(language, l.options) {
		locale = language
	}
	bytes, err := os.ReadFile("lang/" + locale + ".yaml")
	if err != nil {
		return message
	}
	if string(bytes) == "" {
		return message
	}
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		s := strings.Split(line, ": ")
		if key == s[0] {
			message = s[1]
			break
		}
	}
	if len(args) > 0 {
		message = fmt.Sprintf(message, args[0]...)
	}
	return message
}
