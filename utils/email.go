package utils

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/robfig/config"
)

type EmailEntity struct {
	Subject  string   `json:"subject"`
	Content  string   `json:"content"`
	From     string   `json:"from"`
	To       []string `json:"to"`
	Nickname string
}

func SendEmail(e *EmailEntity) {
	cfgPath := CacuCurrentConfigFile()
	cfg, _ := config.ReadDefault(cfgPath)
	authCode, _ := cfg.String("email", "authCode")
	auth := smtp.PlainAuth("", e.From, authCode, "smtp.qq.com")
	to := e.To
	nickname := "Ke Robot"
	if e.Nickname != "" {
		nickname = e.Nickname
	}
	user := e.From
	body := e.Content
	contentType := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + e.Subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("Send mail error: %v", err)
		panic(err)
	}
	fmt.Printf("Send mail success")
}
