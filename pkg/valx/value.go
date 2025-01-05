package valx

import (
	"regexp"
	"strconv"

	json "github.com/json-iterator/go"
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

func ToUint32(in string) uint32 {
	uint32Value, _ := strconv.ParseUint(in, 10, 32)
	return uint32(uint32Value)
}

func ToInt64(in string) int64 {
	val, _ := strconv.ParseInt(in, 10, 64)
	return val
}
