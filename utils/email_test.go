package utils

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	SendEmail(&EmailEntity{
		Subject:  "test entity",
		Content:  "this is content",
		From:     "xxx@qq.com",
		To:       []string{""},
		Nickname: "hapiman",
	})
}
