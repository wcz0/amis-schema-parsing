package tpl

import (
	"amis_schema_parsing/pkg/util"
	"strings"
)

// AmisClassGenerate 生成 amisClass
func AmisClassGenerate() string {
	// 读取 Renderers 目录下的文件名

	files := util.ReadDir("./dist/renderers")

	content := amisClassHeader()

	content += "\n\n"
	content += "/**\n"
	content += " * @author: wcz0\n"
	content += " * version: 6.4.1\n"
	content += " */"
	content += "\npublic class Amis {\n"

	for _, v := range files {
		fileName := strings.ReplaceAll(v, ".java", "")
		className := util.SnakeToPascal(fileName)

		content += `
	public static `+ className + ` ` + className + `(){
		return new ` + className + `();
	}
`
	}

	content += `
}`

	return content
}

// amisClassHeader amisClass 头部
func amisClassHeader() string {
	content := `package com.wcz0;

`
	files := util.ReadDir("./dist/renderers")

	for _, v := range files {
		fileName := strings.ReplaceAll(v, ".java", "")
		content += `import com.wcz0.renderers.` + fileName + `;
`
	}
	return content
}
