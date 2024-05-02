package util

import (
	"archive/zip"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	json "github.com/json-iterator/go"
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

func Sha256(in []byte) string {
	m := sha256.New()
	m.Write(in)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func Date() string {
	return time.Now().Format("2006-01-02")
}

func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func IsExistFolder(folderPath string) bool {
	fileInfo, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return fileInfo.IsDir()
}

func IsExistFile(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ZipDir(dir, output string) error {
	zipfile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 递归遍历目录
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 修正文件路径
		header.Name = filepath.ToSlash(path[len(dir):])
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
}

func ZipFiles(output string, files map[string]string) error {
	// 创建一个 ZIP 文件
	zipFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建一个新的 ZIP 写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 文件名及其在 ZIP 内的重命名
	filesToZip := files

	// 遍历文件列表，逐个添加到 ZIP 文件
	for originalName, newName := range filesToZip {
		// 打开待压缩的文件
		fileToZip, err := os.Open(originalName)
		if err != nil {
			panic(err)
		}
		defer fileToZip.Close()

		// 获取文件的信息，以便复制文件的元数据
		info, err := fileToZip.Stat()
		if err != nil {
			return err
		}

		// 创建 ZIP 文件中的一个条目，并指定新的文件名
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = newName
		header.Method = zip.Deflate // 设置压缩算法

		// 创建条目的写入器
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 将文件内容复制到 ZIP 文件中的条目
		if _, err = io.Copy(writer, fileToZip); err != nil {
			return err
		}
	}
	return nil
}
