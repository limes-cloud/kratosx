package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	json "github.com/json-iterator/go"
	"github.com/mbndr/figlet4go"
	"golang.org/x/crypto/bcrypt"
)

func Transform(in any, dst any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func MD5(in []byte) string {
	sum := md5.Sum(in)
	return hex.EncodeToString(sum[:])
}

func MD5ToUpper(in []byte) string {
	return strings.ToUpper(MD5(in))
}

func ParsePwd(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

type ListType interface {
	~string | ~int | ~uint32 | ~[]byte | ~rune | ~float64
}

func InList[ListType comparable](list []ListType, val ListType) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func IsEmail(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(.[a-zA-Z]{2,})+$`)
	return reg.MatchString(email)
}

func IsPhone(phone string) bool {
	reg := regexp.MustCompile(`^1[3456789]\d{9}$`)
	return reg.MatchString(phone)
}

func ToUint32(in string) uint32 {
	uint32Value, _ := strconv.ParseUint(in, 10, 32)
	return uint32(uint32Value)
}

func ToInt64(in string) int64 {
	val, _ := strconv.ParseInt(in, 10, 64)
	return val
}

func HexToByte(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)
	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}

func PrintArtFont(str string) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	hexColor, _ := figlet4go.NewTrueColorFromHexString("885DBA")
	options.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
		hexColor,
	}

	renderStr, _ := ascii.RenderOpts(str, options)
	fmt.Println(renderStr)
}

func Sha256(in []byte) string {
	m := sha256.New()
	m.Write(in)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
