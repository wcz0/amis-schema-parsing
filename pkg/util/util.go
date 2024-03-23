package util

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

// IsEndWith 是否以指定内容结尾
func IsEndWith(str, suffix string) bool {
	if len(str) < len(suffix) {
		return false
	}

	return str[len(str)-len(suffix):] == suffix
}

func InitDir(dir string) {
	//检查文件夹是否存在
	if IsExist(dir) {
		// 删除文件夹
		_ = os.RemoveAll(dir)
	}

	//创建文件夹
	_ = os.MkdirAll(dir, os.ModePerm)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CopyFile(src, dst string) {
	srcFile, _ := os.Open(src)
	defer srcFile.Close()

	dstFile, _ := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	defer dstFile.Close()

	_, _ = io.Copy(dstFile, srcFile)
}

func FileToArray(path string) []string {
	content, _ := ioutil.ReadFile(path)

	// 按行分割
	return strings.Split(string(content), "\r")
}

func IsInArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}

	return false
}

// ClearLineBreak 清除传入内容中的换行符
func ClearLineBreak(content string) string {
	return strings.Replace(content, "\n\n", "", -1)
}

// KebabToCamel 中划线转小驼峰
func KebabToCamel(str string) string {
	arr := strings.Split(str, "-")

	for i, v := range arr {
		if i > 0 {
			arr[i] = strings.Title(v)
		}
	}

	return strings.Join(arr, "")
}

// ReadDir 读取目录下的所有文件, 返回文件名列表
func ReadDir(dir string) []string {
	var files []string

	dirList, _ := ioutil.ReadDir(dir)

	for _, v := range dirList {
		if !v.IsDir() {
			files = append(files, v.Name())
		}
	}

	return files
}

// GetClassName 获取类名
func GetClassName(className string) string {
	// 去除 Schema | Object
	if IsEndWith(className, "Schema") || IsEndWith(className, "Object") {
		className = className[:len(className)-6]
	}

	// 关键字冲突
	if className == "List" {
		className = "ListRenderer"
	}

	return className
}

// PascalToCamel 大驼峰转小驼峰
func PascalToCamel(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// PascalToSnake 大驼峰转蛇形
func PascalToSnake(s string) string {
	var result string
	for i, v := range s {
		if unicode.IsUpper(v) {
			if i != 0 {
				result += "_"
			}
			result += string(unicode.ToLower(v))
		} else {
			result += string(v)
		}
	}
	return result
}


func SnakeToPascal(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_'
	})

	for i := 0; i < len(words); i++ {
		runes := []rune(words[i])
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}

	return strings.Join(words, "")
}
