package main

import (
	"log"
	"net/smtp"
	"strconv"

	"gopkg.in/mail.v2"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

func SendMail(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": "cloud_monitor@cloudminds.com.cn",
		"pass": "Cloud2020",
		"host": "mail.cloudminds.com.cn",
		"port": "25",
	}
	port, _ := strconv.Atoi(mailConn["port"])

	m := mail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "达闼云告警"))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	auth := LoginAuth(mailConn["user"], mailConn["pass"])
	d := mail.Dialer{Host: mailConn["host"], Port: port, StartTLSPolicy: mail.NoStartTLS, Auth: auth}

	err := d.DialAndSend(m)
	return err
}

func main() {
	err := SendMail([]string{"xiaobing.song@cloudminds.com"}, "TEST", "Body")
	if err != nil {
		log.Printf("发送错误:%s", err.Error())
	}
}
