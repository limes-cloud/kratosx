package proto

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
)

type Error struct {
	*gen.Builder
}

type ErrorCode struct {
	sort   []string
	bucket map[string]string
}

func NewErrorBuilder(builder *gen.Builder) *Error {
	return &Error{
		Builder: builder,
	}
}

// ParseErrorByContent 解析proto文本，生成proto结构
func (p *Error) ParseErrorByContent(content string) *ErrorCode {
	reply := &ErrorCode{bucket: make(map[string]string)}
	enumEntryRegex := regexp.MustCompile(`(\w+)\s*=\s*\d+\s*\[(.*?)\];`)
	matches := enumEntryRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 3 {
			reply.sort = append(reply.sort, match[1])
			reply.bucket[match[1]] = fmt.Sprintf("[%s]", match[2])
		}
	}
	return reply
}

// ScanError 扫描现在的已经存在的proto error定义
func (p *Error) ScanError() (*ErrorCode, error) {
	path := p.ProtoErrorPath()
	if !pkg.IsExistFile(path) {
		return &ErrorCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseErrorByContent(string(code)), nil
}

// MakeError 根据模板生成proto
func (p *Error) MakeError() (*ErrorCode, error) {
	code, err := os.ReadFile(p.ProtoErrorTplPath())
	if err != nil {
		return nil, err
	}
	return p.ParseErrorByContent(string(code)), nil
}

func (p *Error) RenderError(proto *ErrorCode) string {
	var lines []string
	code, _ := os.ReadFile(p.ProtoErrorTplPath())
	enumEntryRegex := regexp.MustCompile(`(\w+)\s*=\s*\d+\s*\[(.*?)\];\s`)
	matches := enumEntryRegex.ReplaceAllString(string(code), "")
	matches = strings.TrimSuffix(strings.TrimSpace(matches), "}")
	for ind, name := range proto.sort {
		lines = append(lines, fmt.Sprintf("  %s = %d%s;", name, ind, proto.bucket[name]))
	}
	matches = strings.TrimSpace(matches) + "\n\n"
	matches += strings.Join(lines, "\n") + "\n}"
	return matches
}

// GenError 生成error proto
func (p *Error) GenError() (string, error) {
	// 扫描已经生成的proto
	scanError, err := p.ScanError()
	if err != nil {
		return "", fmt.Errorf("扫描proto代码失败，%s", err.Error())
	}

	// 生成新的proto
	makeError, err := p.MakeError()
	if err != nil {
		return "", fmt.Errorf("生成proto代码失败，%s", err.Error())
	}

	// 合并proto
	code := &ErrorCode{bucket: make(map[string]string)}
	code.sort = append(makeError.sort, scanError.sort...)
	for key, val := range scanError.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeError.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)

	// 生成proto
	return p.RenderError(code), nil
}
