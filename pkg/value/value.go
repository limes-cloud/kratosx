package value

import (
	json "github.com/json-iterator/go"
	"regexp"
	"unsafe"
)

var (
	emailReg = regexp.MustCompile(`\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`)
	phoneReg = regexp.MustCompile(`^1[3456789]\d{9}$`)
)

func Transform(in any, dst any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func ObjToString(obj any) string {
	if obj == nil {
		return ""
	}
	b, _ := json.Marshal(obj)
	v := *(*string)(unsafe.Pointer(&b))
	if v == "{}" || v == "null" {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

func IsEmail(email string) bool {
	return emailReg.MatchString(email)
}

func IsPhone(phone string) bool {
	return phoneReg.MatchString(phone)
}
