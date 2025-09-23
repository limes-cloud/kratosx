package value

import (
	json "github.com/json-iterator/go"
	"regexp"
)

var (
	emailReg = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(.[a-zA-Z]{2,})+$`)
	phoneReg = regexp.MustCompile(`^1[3456789]\d{9}$`)
)

func Transform(in any, dst any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func IsEmail(email string) bool {
	return emailReg.MatchString(email)
}

func IsPhone(phone string) bool {
	return phoneReg.MatchString(phone)
}
