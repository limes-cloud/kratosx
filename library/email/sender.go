package email

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"unsafe"

	"gopkg.in/gomail.v2"

	ec "github.com/limes-cloud/kratosx/config"
)

type Sender interface {
	Send(u string, name string, variable ...any) error
}

type sender struct {
	set  map[string][]byte
	conf *ec.Email
	stp  string
}

func (s *sender) Send(email string, name string, variable ...any) error {
	if s.stp == "" {
		return errors.New("please choose send template")
	}

	if s.set[s.stp] == nil {
		return fmt.Errorf("template file %v not exist", s.stp)
	}

	tpc := s.conf.Template[s.stp]
	tpv := s.set[s.stp]
	conf := s.conf

	n := template.New("")
	params := any(nil)
	if len(variable) != 0 {
		params = variable[0]
	}
	parser, err := n.Parse(*(*string)(unsafe.Pointer(&tpv)))
	if err != nil {
		return err
	}
	html := bytes.NewBuffer([]byte(""))
	if err = parser.Execute(html, params); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(conf.User, conf.Name))
	m.SetHeader("To", m.FormatAddress(email, name))
	m.SetHeader("Subject", tpc.Subject)
	m.SetBody(fmt.Sprintf("%v; charset=UTF-8", tpc.Type), html.String())
	d := gomail.NewDialer(
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
	)
	return d.DialAndSend(m)

	// auth := smtp.PlainAuth("", conf.User, conf.Password, conf.Host)
	// ct := fmt.Sprintf("Content-Type: %v; charset=UTF-8", tpc.Type)
	// to := fmt.Sprintf("To: %s<%s>", name, email)
	// from := fmt.Sprintf("From: %s<%s>", conf.Name, conf.User)
	// subject := fmt.Sprintf("Subject: %s", tpc.Subject)
	// msg := []byte(to + "\r\n" + from + "\r\n" + subject + "\r\n" + ct + "\r\n\r\n" + html.String())
	// host := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	// err = smtp.SendMail(host, auth, conf.User, []string{email}, msg)
	// return err
}
