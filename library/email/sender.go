package email

import (
	"bytes"
	"errors"
	"fmt"
	ec "github.com/limes-cloud/kratosx/config"
	"html/template"
	"net/smtp"
	"unsafe"
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

	auth := smtp.PlainAuth("", conf.User, conf.Password, conf.Host)
	ct := fmt.Sprintf("Content-Type: %v; charset=UTF-8", tpc.Type)
	to := fmt.Sprintf("To: %s<%s>", name, email)
	from := fmt.Sprintf("From: %s<%s>", conf.Name, conf.User)
	subject := fmt.Sprintf("Subject: %s", tpc.Subject)
	msg := []byte(to + "\r\n" + from + "\r\n" + subject + "\r\n" + ct + "\r\n\r\n" + html.String())
	host := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	return smtp.SendMail(host, auth, conf.User, []string{email}, msg)
}
