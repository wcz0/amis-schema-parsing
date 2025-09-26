package util

import (
	"os"
	"strings"
)

// PascalToSnake 将 PascalCase 转换为 snake_case
func PascalToSnake(s string) string {
	if s == "" {
		return s
	}
	
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	
	return strings.ToLower(result.String())
}

// InitDir 初始化目录，如果不存在则创建
func InitDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}