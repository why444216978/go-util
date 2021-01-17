package mail

import (
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

const (
	HOST = "xxx"
	PORT = 465
	PASS = ""
)

type Options struct {
	MailUser   string // 发件人
	MailTo     string // 收件人 多个用,分割
	Subject    string // 邮件主题
	Body       string // 邮件内容
	MailAttach []string
}

// NewMail 返回新的邮件配置
func NewMail(send, to, subject, body string, attach []string) *Options {
	return &Options{
		MailUser:   send,
		MailTo:     to,
		Subject:    subject,
		Body:       body,
		MailAttach: attach,
	}
}

// SeSend 根据配置发送邮件
func Send(o *Options) error {

	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", o.MailUser)

	//设置发送给多个用户
	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", o.Subject)

	//设置邮件正文
	m.SetBody("text/html", o.Body)

	if o.MailAttach != nil {
		for _, v := range o.MailAttach {
			m.Attach(v)
		}
	}

	d := gomail.NewDialer(HOST, PORT, o.MailUser, PASS)

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
