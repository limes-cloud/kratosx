package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func MD5(in []byte) string {
	sum := md5.Sum(in)
	return hex.EncodeToString(sum[:])
}

func MD5ToUpper(in []byte) string {
	return strings.ToUpper(MD5(in))
}

func EncodePwd(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

func HexToByte(h string) ([]byte, error) {
	return hex.DecodeString(h)
}

func Sha256(in []byte) string {
	m := sha256.New()
	m.Write(in)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
