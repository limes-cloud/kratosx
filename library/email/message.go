package email

type Message struct {
	email string //发送者邮箱
	name  string //发送着名称
}

func NewMessage(email string, fs ...func(*Message)) *Message {
	msg := &Message{
		email: email,
	}
	if len(fs) == 0 {
		return msg
	}
	fs[0](msg)
	return msg
}

// Name 设置接受者者名称
func (msg *Message) Name(val string) {
	msg.name = val
}
