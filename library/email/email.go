package email

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"text/template"
	"unsafe"

	"gopkg.in/gomail.v2"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
)

type Email interface {
	Template(name string) Sender
}

type email struct {
	mu   sync.RWMutex
	set  map[string][]byte
	conf *config.Email
}

var (
	ins *email

	once sync.Once
)

// Instance 获取email对象实例
func Instance() Email {
	return ins
}

func Init(conf *config.Email, watcher config.Watcher) {
	// 没有模板则跳过初始化
	if conf == nil || len(conf.Template) == 0 {
		return
	}

	once.Do(func() {
		ins = &email{
			mu:   sync.RWMutex{},
			set:  make(map[string][]byte),
			conf: conf,
		}
		// 遍历初始化模板
		for key, tpc := range conf.Template {
			if err := ins.initFactory(key, tpc); err != nil {
				panic("Email 初始化失败 :" + err.Error())
			}

			watcher("email.template."+key, func(value config.Value) {
				if err := value.Scan(&tpc); err != nil {
					logger.Instance().Error("email watch config error", logger.F("err", err))
					return
				}
				if err := ins.initFactory(key, tpc); err != nil {
					logger.Instance().Error("email init error", logger.F("err", err))
				}
			})
		}
	})
}

func (c *email) Template(name string) Sender {
	return &sender{
		stp:  name,
		set:  c.set,
		conf: c.conf,
	}
}

func (c *email) initFactory(name string, et config.EmailTemplate) error {
	if et.Enable != nil && !*et.Enable {
		c.delete(name)
		return nil
	}

	// 获取文件内容
	file, err := os.Open(et.Path)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	defer file.Close()

	c.mu.Lock()
	c.set[name] = all
	c.mu.Unlock()
	return nil
}

func (c *email) delete(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.set, name)
}

type Sender interface {
	// Send 发送邮件信息
	Send(email string, opts ...SendOption) error
}

type sendOpt struct {
	name     string
	variable any
}

type SendOption func(*sendOpt)

// WithRecvName 接受者姓名
func WithRecvName(name string) SendOption {
	return func(s *sendOpt) {
		s.name = name
	}
}

// WithTemplateVariable 模板变量
func WithTemplateVariable(variable any) SendOption {
	return func(s *sendOpt) {
		s.variable = variable
	}
}

type sender struct {
	set  map[string][]byte
	conf *config.Email
	stp  string
}

func (s *sender) Send(email string, opts ...SendOption) error {
	o := &sendOpt{}
	for _, opt := range opts {
		opt(o)
	}

	if s.stp == "" {
		return errors.New("please choose send template")
	}

	if s.set[s.stp] == nil {
		return fmt.Errorf("template file %v not exist", s.stp)
	}

	tpc := s.conf.Template[s.stp]
	tpv := s.set[s.stp]
	conf := s.conf

	n := template.New("content")

	parser, err := n.Parse(*(*string)(unsafe.Pointer(&tpv)))
	if err != nil {
		return err
	}
	html := bytes.NewBuffer([]byte(""))
	if err = parser.Execute(html, o.variable); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(conf.User, conf.Name))
	m.SetHeader("To", m.FormatAddress(email, o.name))
	m.SetHeader("Subject", tpc.Subject)
	m.SetBody(fmt.Sprintf("%v; charset=UTF-8", tpc.Type), html.String())
	d := gomail.NewDialer(
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
	)
	return d.DialAndSend(m)
}
