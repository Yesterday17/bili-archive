// From annie/utils/utils.go
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 文件名最长值
const MAXLENGTH = 80

// 限制文件名的长度
func LimitLength(s string, length int) string {
	const ELLIPSES = "..."
	str := []rune(s)
	if len(str) > length {
		return string(str[:length-len(ELLIPSES)]) + ELLIPSES
	}
	return s
}

// 使文件名合法化
func FileName(name string) string {
	rep := strings.NewReplacer("\n", " ", "/", " ", "|", "-", ": ", "：", ":", "：", "'", "’")
	name = rep.Replace(name)
	if runtime.GOOS == "windows" {
		rep = strings.NewReplacer("\"", " ", "?", " ", "*", " ", "\\", " ", "<", " ", ">", " ")
		name = rep.Replace(name)
	}
	return LimitLength(name, MAXLENGTH)
}

// 使文件路径合法化
func FilePath(outputPath, name, ext string, escape bool) (string, error) {
	if outputPath != "" {
		if _, err := os.Stat(outputPath); err != nil {
			return "", err
		}
	}
	fileName := fmt.Sprintf("%s.%s", name, ext)
	if escape {
		fileName = FileName(fileName)
	}
	outputPath = filepath.Join(outputPath, fileName)
	return outputPath, nil
}

// 检测文件存在
func FileSize(filePath string) (int64, bool, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return file.Size(), true, nil
}
