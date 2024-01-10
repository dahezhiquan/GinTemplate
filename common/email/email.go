package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Email struct {
	email      *email.Email
	serverAddr string
	smtp       string
	token      string
	host       string
}

type EmailOtpions struct {
	From       string
	ServerAddr string
	Smtp       string
	Token      string
	Host       string
}

// DefaultEmailOptions 配置邮件信息
var DefaultEmailOptions = &EmailOtpions{
	From:       "demo <111@qq.com>",
	ServerAddr: "111@qq.com",
	Smtp:       "smtp.qq.com:25",
	Token:      "demodemodemodemodemo",
	Host:       "smtp.qq.com",
}

func (e *Email) SetText(text string) *Email {
	e.email.Text = []byte(text)
	return e
}

func (e *Email) SetSubject(subject string) *Email {
	e.email.Subject = subject
	return e
}

func (e *Email) SetHtml(html []byte) *Email {
	e.email.HTML = html
	return e
}

func (e *Email) SetTo(receivers ...string) *Email {
	e.email.To = receivers
	return e
}

func (e *Email) SetFrom(from string) *Email {
	e.email.From = from
	return e
}

func (e *Email) Send() error {
	return e.email.Send(e.smtp, smtp.PlainAuth("", e.serverAddr, e.token, e.host))
}

func NewMail(opt *EmailOtpions) *Email {
	if opt != nil {
		mail := email.NewEmail()
		mail.From = opt.From
		return &Email{
			email:      mail,
			serverAddr: opt.ServerAddr,
			smtp:       opt.Smtp,
			token:      opt.Token,
			host:       opt.Host,
		}
	}
	return &Email{email: email.NewEmail()}
}
